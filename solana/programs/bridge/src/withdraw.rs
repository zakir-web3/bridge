use anchor_lang::prelude::*;
use anchor_spl::associated_token::AssociatedToken;
use anchor_spl::token::{self, Mint, Token, TokenAccount, TransferChecked};

use crate::state::*;

// ---------------------------------------------------------------------------
// EIP-712 constants (must byte-align with Go / EVM)
// ---------------------------------------------------------------------------

/// `keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)")`
const EIP712_DOMAIN_TYPEHASH: [u8; 32] = {
    const HASH: [u8; 32] = keccak256_const(
        b"EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)",
    );
    HASH
};

/// `keccak256("Withdraw(address user,bytes32 destination,bytes32 token,uint256 amount,uint256 chainId,uint64 nonce)")`
const WITHDRAW_TYPEHASH: [u8; 32] = {
    const HASH: [u8; 32] = keccak256_const(
        b"Withdraw(address user,bytes32 destination,bytes32 token,uint256 amount,uint256 chainId,uint64 nonce)",
    );
    HASH
};

/// `keccak256("Bridge")`
const NAME_HASH: [u8; 32] = keccak256_const(b"Bridge");

/// `keccak256("1")`
const VERSION_HASH: [u8; 32] = keccak256_const(b"1");

/// secp256k1 half-order `n/2` for low-S enforcement.
const SECP256K1_HALF_ORDER: [u8; 32] = [
    0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
    0xFF, 0x5D, 0x57, 0x6E, 0x73, 0x57, 0xA4, 0x50, 0x1D, 0xDF, 0xE9, 0x2F, 0x46, 0x68, 0x1B,
    0x20, 0xA0,
];

// ---------------------------------------------------------------------------
// Compile-time keccak256 (tiny Keccak-f[1600] for const evaluation)
// ---------------------------------------------------------------------------

/// Compile-time keccak256 for constant inputs (typehashes, name/version).
/// NOT used on-chain — the syscall `solana_keccak_hasher::hashv` handles runtime hashing.
const fn keccak256_const(input: &[u8]) -> [u8; 32] {
    const RATE: usize = 136;
    let mut state = [0u64; 25];
    let mut buf = [0u8; RATE];
    let mut buf_len = 0usize;

    let mut i = 0;
    while i < input.len() {
        buf[buf_len] = input[i];
        buf_len += 1;
        if buf_len == RATE {
            state = absorb(state, &buf);
            buf = [0u8; RATE];
            buf_len = 0;
        }
        i += 1;
    }
    buf[buf_len] = 0x01;
    buf[RATE - 1] |= 0x80;
    state = absorb(state, &buf);
    squeeze(state)
}

const fn absorb(mut state: [u64; 25], block: &[u8; 136]) -> [u64; 25] {
    let mut i = 0;
    while i < 136 / 8 {
        let off = i * 8;
        let word = (block[off] as u64)
            | (block[off + 1] as u64) << 8
            | (block[off + 2] as u64) << 16
            | (block[off + 3] as u64) << 24
            | (block[off + 4] as u64) << 32
            | (block[off + 5] as u64) << 40
            | (block[off + 6] as u64) << 48
            | (block[off + 7] as u64) << 56;
        state[i] ^= word;
        i += 1;
    }
    keccak_f1600(state)
}

const fn squeeze(state: [u64; 25]) -> [u8; 32] {
    let mut out = [0u8; 32];
    let mut i = 0;
    while i < 4 {
        let w = state[i];
        let off = i * 8;
        out[off] = w as u8;
        out[off + 1] = (w >> 8) as u8;
        out[off + 2] = (w >> 16) as u8;
        out[off + 3] = (w >> 24) as u8;
        out[off + 4] = (w >> 32) as u8;
        out[off + 5] = (w >> 40) as u8;
        out[off + 6] = (w >> 48) as u8;
        out[off + 7] = (w >> 56) as u8;
        i += 1;
    }
    out
}

#[rustfmt::skip]
const RC: [u64; 24] = [
    0x0000000000000001, 0x0000000000008082, 0x800000000000808A, 0x8000000080008000,
    0x000000000000808B, 0x0000000080000001, 0x8000000080008081, 0x8000000000008009,
    0x000000000000008A, 0x0000000000000088, 0x0000000080008009, 0x000000008000000A,
    0x000000008000808B, 0x800000000000008B, 0x8000000000008089, 0x8000000000008003,
    0x8000000000008002, 0x8000000000000080, 0x000000000000800A, 0x800000008000000A,
    0x8000000080008081, 0x8000000000008080, 0x0000000080000001, 0x8000000080008008,
];

const fn keccak_f1600(mut a: [u64; 25]) -> [u64; 25] {
    let mut round = 0;
    while round < 24 {
        // θ
        let c = [
            a[0] ^ a[5] ^ a[10] ^ a[15] ^ a[20],
            a[1] ^ a[6] ^ a[11] ^ a[16] ^ a[21],
            a[2] ^ a[7] ^ a[12] ^ a[17] ^ a[22],
            a[3] ^ a[8] ^ a[13] ^ a[18] ^ a[23],
            a[4] ^ a[9] ^ a[14] ^ a[19] ^ a[24],
        ];
        let d = [
            c[4] ^ c[1].rotate_left(1),
            c[0] ^ c[2].rotate_left(1),
            c[1] ^ c[3].rotate_left(1),
            c[2] ^ c[4].rotate_left(1),
            c[3] ^ c[0].rotate_left(1),
        ];
        let mut i = 0;
        while i < 25 {
            a[i] ^= d[i % 5];
            i += 1;
        }
        // ρ + π
        let mut b = [0u64; 25];
        #[rustfmt::skip]
        const PILN: [(usize, u32); 25] = [
            (0,0),(6,44),(12,43),(18,21),(24,14),
            (3,28),(9,20),(10,3),(16,45),(22,61),
            (1,1),(7,6),(13,25),(19,8),(20,18),
            (4,27),(5,36),(11,10),(17,15),(23,56),
            (2,62),(8,55),(14,39),(15,41),(21,2),
        ];
        let mut j = 0;
        while j < 25 {
            let (src, rot) = PILN[j];
            b[j] = a[src].rotate_left(rot);
            j += 1;
        }
        // χ
        i = 0;
        while i < 5 {
            let mut j2 = 0;
            while j2 < 5 {
                a[5 * i + j2] = b[5 * i + j2] ^ (!b[5 * i + (j2 + 1) % 5] & b[5 * i + (j2 + 2) % 5]);
                j2 += 1;
            }
            i += 1;
        }
        // ι
        a[0] ^= RC[round];
        round += 1;
    }
    a
}

// ---------------------------------------------------------------------------
// Runtime helpers (use Solana syscalls on-chain)
// ---------------------------------------------------------------------------

fn keccak256(data: &[u8]) -> [u8; 32] {
    solana_keccak_hasher::hash(data).to_bytes()
}

fn keccak256v(slices: &[&[u8]]) -> [u8; 32] {
    solana_keccak_hasher::hashv(slices).to_bytes()
}

/// Compute the EIP-712 domain separator. Called once during `initialize`.
pub fn compute_domain_separator(chain_id: u64, verifying_contract: &[u8; 20]) -> [u8; 32] {
    let mut chain_id_word = [0u8; 32];
    chain_id_word[24..].copy_from_slice(&chain_id.to_be_bytes());

    let mut vc_word = [0u8; 32];
    vc_word[12..].copy_from_slice(verifying_contract);

    keccak256v(&[
        &EIP712_DOMAIN_TYPEHASH,
        &NAME_HASH,
        &VERSION_HASH,
        &chain_id_word,
        &vc_word,
    ])
}

/// Derive a pseudo-EVM address from a Solana program ID: `keccak256(program_id)[12..32]`.
pub fn derive_verifying_contract(program_id: &Pubkey) -> [u8; 20] {
    let hash = keccak256(program_id.as_ref());
    let mut out = [0u8; 20];
    out.copy_from_slice(&hash[12..32]);
    out
}

/// Compute the EIP-712 struct hash for a Withdraw message.
fn compute_struct_hash(
    user: &[u8; 20],
    destination: &[u8; 32],
    token: &[u8; 32],
    amount: u64,
    chain_id: u64,
    nonce: u64,
) -> [u8; 32] {
    let mut user_word = [0u8; 32];
    user_word[12..].copy_from_slice(user);

    let mut amount_word = [0u8; 32];
    amount_word[24..].copy_from_slice(&amount.to_be_bytes());

    let mut chain_id_word = [0u8; 32];
    chain_id_word[24..].copy_from_slice(&chain_id.to_be_bytes());

    let mut nonce_word = [0u8; 32];
    nonce_word[24..].copy_from_slice(&nonce.to_be_bytes());

    keccak256v(&[
        &WITHDRAW_TYPEHASH,
        &user_word,
        destination,
        token,
        &amount_word,
        &chain_id_word,
        &nonce_word,
    ])
}

/// Compute the final EIP-712 digest: `keccak256(0x19 || 0x01 || domainSeparator || structHash)`.
fn compute_digest(domain_separator: &[u8; 32], struct_hash: &[u8; 32]) -> [u8; 32] {
    keccak256v(&[&[0x19, 0x01], domain_separator, struct_hash])
}

/// Returns `true` when `s <= SECP256K1_HALF_ORDER`.
fn is_low_s(sig: &[u8; 64]) -> bool {
    let s = &sig[32..64];
    for i in 0..32 {
        if s[i] < SECP256K1_HALF_ORDER[i] {
            return true;
        }
        if s[i] > SECP256K1_HALF_ORDER[i] {
            return false;
        }
    }
    true
}

/// Recover an Ethereum address from an ECDSA signature over a keccak256 digest.
fn recover_eth_address(digest: &[u8; 32], sig: &[u8; 64], recovery_id: u8) -> Result<[u8; 20]> {
    let pubkey = solana_secp256k1_recover::secp256k1_recover(digest, recovery_id, sig)
        .map_err(|_| error!(BridgeError::SignatureRecoveryFailed))?;
    let hash = keccak256(&pubkey.0);
    let mut addr = [0u8; 20];
    addr.copy_from_slice(&hash[12..32]);
    Ok(addr)
}

// ---------------------------------------------------------------------------
// Instruction data
// ---------------------------------------------------------------------------

/// ECDSA signature from a Bridge validator.
#[derive(AnchorSerialize, AnchorDeserialize, Clone)]
pub struct EcdsaSignature {
    pub sig: [u8; 64],
    pub recovery_id: u8,
}

// ---------------------------------------------------------------------------
// Accounts
// ---------------------------------------------------------------------------

#[derive(Accounts)]
#[instruction(user: [u8; 20], destination: [u8; 32], amount: u64, nonce: u64)]
pub struct Withdraw<'info> {
    #[account(mut)]
    pub payer: Signer<'info>,

    #[account(
        seeds = [b"config"],
        bump = config.bump,
    )]
    pub config: Account<'info, BridgeConfig>,

    #[account(
        seeds = [b"validator_set"],
        bump = validator_set.bump,
    )]
    pub validator_set: Account<'info, ValidatorSet>,

    pub mint: Account<'info, Mint>,

    #[account(
        seeds = [b"vault_state", mint.key().as_ref()],
        bump = vault_state.bump,
        constraint = vault_state.mint == mint.key() @ BridgeError::MintMismatch,
    )]
    pub vault_state: Account<'info, VaultState>,

    /// CHECK: PDA used as SPL token account authority for the vault.
    #[account(
        seeds = [b"vault", mint.key().as_ref()],
        bump,
    )]
    pub vault_authority: UncheckedAccount<'info>,

    #[account(
        mut,
        associated_token::mint = mint,
        associated_token::authority = vault_authority,
    )]
    pub vault_token_account: Account<'info, TokenAccount>,

    /// CHECK: Must match `destination` bytes (verified in handler).
    pub destination_owner: UncheckedAccount<'info>,

    #[account(
        init_if_needed,
        payer = payer,
        associated_token::mint = mint,
        associated_token::authority = destination_owner,
    )]
    pub destination_token_account: Account<'info, TokenAccount>,

    #[account(
        init_if_needed,
        payer = payer,
        space = 8 + NoncePage::INIT_SPACE,
        seeds = [b"nonce_page" as &[u8], &(nonce / 8192u64).to_le_bytes()],
        bump,
    )]
    pub nonce_page: Account<'info, NoncePage>,

    pub token_program: Program<'info, Token>,
    pub associated_token_program: Program<'info, AssociatedToken>,
    pub system_program: Program<'info, System>,
}

// ---------------------------------------------------------------------------
// Handler
// ---------------------------------------------------------------------------

pub fn handle_withdraw(
    ctx: Context<Withdraw>,
    user: [u8; 20],
    destination: [u8; 32],
    amount: u64,
    nonce: u64,
    signatures: Vec<EcdsaSignature>,
) -> Result<()> {
    let config = &ctx.accounts.config;

    // 1. Not paused
    require!(!config.withdraw_paused, BridgeError::WithdrawPaused);

    // 2. Amount > 0
    require!(amount > 0, BridgeError::ZeroAmount);

    // 3. Mint matches vault (enforced by account constraint, but belt-and-suspenders)
    require!(
        ctx.accounts.vault_state.mint == ctx.accounts.mint.key(),
        BridgeError::MintMismatch
    );

    // 4. destination == destination_owner.key()
    require!(
        destination == ctx.accounts.destination_owner.key().to_bytes(),
        BridgeError::DestinationMismatch
    );

    // 5. Nonce not already used
    let page_num = nonce / 8192;
    let bit_index = (nonce % 8192) as usize;
    let byte_index = bit_index / 8;
    let bit_mask = 1u8 << (bit_index % 8);

    let nonce_page = &ctx.accounts.nonce_page;
    require!(
        nonce_page.bits[byte_index] & bit_mask == 0,
        BridgeError::NonceAlreadyUsed
    );

    // 6. Compute EIP-712 digest
    let token_bytes: [u8; 32] = ctx.accounts.mint.key().to_bytes();
    let struct_hash = compute_struct_hash(
        &user,
        &destination,
        &token_bytes,
        amount,
        config.chain_id,
        nonce,
    );
    let digest = compute_digest(&config.domain_separator, &struct_hash);

    // 7–8. Verify signatures, accumulate power, check quorum
    let validator_set = &ctx.accounts.validator_set;
    require!(
        !validator_set.validators.is_empty(),
        BridgeError::ValidatorSetNotInitialized
    );

    let mut accumulated_power: u128 = 0;
    let mut seen_mask: u64 = 0; // bitmask over validator indices (max 64)

    for sig in &signatures {
        require!(is_low_s(&sig.sig), BridgeError::HighSValue);

        let recovered = recover_eth_address(&digest, &sig.sig, sig.recovery_id)?;

        let (validator_index, power) = validator_set
            .validators
            .iter()
            .enumerate()
            .find(|(_, v)| v.eth_address == recovered)
            .map(|(i, v)| (i, v.power))
            .ok_or(error!(BridgeError::UnknownValidator))?;

        let mask_bit = 1u64 << validator_index;
        require!(seen_mask & mask_bit == 0, BridgeError::DuplicateSignature);
        seen_mask |= mask_bit;

        accumulated_power += power as u128;
    }

    require!(
        3u128 * accumulated_power > 2u128 * (validator_set.total_power as u128),
        BridgeError::InsufficientQuorum
    );

    // 9. Mark nonce as used
    let nonce_page = &mut ctx.accounts.nonce_page;
    nonce_page.page = page_num;
    nonce_page.bump = ctx.bumps.nonce_page;
    nonce_page.bits[byte_index] |= bit_mask;

    // 10. Transfer tokens from vault to destination
    let mint_key = ctx.accounts.mint.key();
    let vault_bump = ctx.bumps.vault_authority;
    let signer_seeds: &[&[u8]] = &[b"vault", mint_key.as_ref(), &[vault_bump]];

    token::transfer_checked(
        CpiContext::new_with_signer(
            Token::id(),
            TransferChecked {
                from: ctx.accounts.vault_token_account.to_account_info(),
                mint: ctx.accounts.mint.to_account_info(),
                to: ctx.accounts.destination_token_account.to_account_info(),
                authority: ctx.accounts.vault_authority.to_account_info(),
            },
            &[signer_seeds],
        ),
        amount,
        ctx.accounts.mint.decimals,
    )?;

    // 11. Emit event for future relayer scanning
    emit!(WithdrawFinalized {
        message: struct_hash,
        user,
        destination,
        mint: ctx.accounts.mint.key(),
        amount,
        nonce,
    });

    Ok(())
}

// ---------------------------------------------------------------------------
// Unit tests (EIP-712 conformance against Go test vectors)
// ---------------------------------------------------------------------------

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_withdraw_typehash() {
        let expected = keccak256(
            b"Withdraw(address user,bytes32 destination,bytes32 token,uint256 amount,uint256 chainId,uint64 nonce)",
        );
        assert_eq!(WITHDRAW_TYPEHASH, expected);
    }

    #[test]
    fn test_domain_typehash() {
        let expected = keccak256(
            b"EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)",
        );
        assert_eq!(EIP712_DOMAIN_TYPEHASH, expected);
    }

    #[test]
    fn test_name_hash() {
        assert_eq!(NAME_HASH, keccak256(b"Bridge"));
    }

    #[test]
    fn test_version_hash() {
        assert_eq!(VERSION_HASH, keccak256(b"1"));
    }

    /// Conformance: Go test `TestBridgeHubWithdraw_ToTypedData` in
    /// `internal/contract/bridge_hub_test.go`.
    ///
    /// Parameters:
    /// - user: 0x1000000000000000000000000000000000000001
    /// - destination: AddressToBytes32(0x1000000000000000000000000000000000000001)
    /// - token: AddressToBytes32(0x2000000000000000000000000000000000000002)
    /// - amount: 1e18
    /// - chainId (domain + struct): 1337
    /// - nonce: 12345
    /// - verifyingContract: 0x3000000000000000000000000000000000000003
    #[test]
    fn test_struct_hash_go_vector() {
        let user: [u8; 20] = hex_decode_20("1000000000000000000000000000000000000001");
        let destination = address_to_bytes32(&hex_decode_20(
            "1000000000000000000000000000000000000001",
        ));
        let token = address_to_bytes32(&hex_decode_20(
            "2000000000000000000000000000000000000002",
        ));
        let amount: u64 = 1_000_000_000_000_000_000; // 1e18
        let chain_id: u64 = 1337;
        let nonce: u64 = 12345;

        let struct_hash =
            compute_struct_hash(&user, &destination, &token, amount, chain_id, nonce);

        assert_eq!(
            hex::encode(struct_hash),
            "2486f3baf331176f06d4eda1e971483466ca196931705258bc5ef971b81d4ad2"
        );
    }

    #[test]
    fn test_domain_separator_go_vector() {
        let verifying_contract: [u8; 20] =
            hex_decode_20("3000000000000000000000000000000000000003");
        let chain_id: u64 = 1337;

        let ds = compute_domain_separator(chain_id, &verifying_contract);

        assert_eq!(
            hex::encode(ds),
            "c41231ad1b7e9b5eed07b994b8348bb7fc45680241556e46317c680e8c588550"
        );
    }

    #[test]
    fn test_full_digest_go_vector() {
        let user: [u8; 20] = hex_decode_20("1000000000000000000000000000000000000001");
        let destination = address_to_bytes32(&hex_decode_20(
            "1000000000000000000000000000000000000001",
        ));
        let token = address_to_bytes32(&hex_decode_20(
            "2000000000000000000000000000000000000002",
        ));
        let amount: u64 = 1_000_000_000_000_000_000;
        let chain_id: u64 = 1337;
        let nonce: u64 = 12345;
        let verifying_contract: [u8; 20] =
            hex_decode_20("3000000000000000000000000000000000000003");

        let ds = compute_domain_separator(chain_id, &verifying_contract);
        let sh = compute_struct_hash(&user, &destination, &token, amount, chain_id, nonce);
        let digest = compute_digest(&ds, &sh);

        assert_eq!(
            hex::encode(digest),
            "04f7234d8a5f15bd6cd22589a81bbff694d8b707f81df216635d52a3934a84e8"
        );
    }

    #[test]
    fn test_low_s_boundary() {
        // Exactly at half-order → valid
        let mut sig_at_half = [0u8; 64];
        sig_at_half[32..].copy_from_slice(&SECP256K1_HALF_ORDER);
        assert!(is_low_s(&sig_at_half));

        // One above half-order → invalid
        let mut sig_above = [0u8; 64];
        sig_above[32..].copy_from_slice(&SECP256K1_HALF_ORDER);
        sig_above[63] += 1; // increment last byte
        assert!(!is_low_s(&sig_above));

        // Zero S → valid
        let sig_zero = [0u8; 64];
        assert!(is_low_s(&sig_zero));
    }

    #[test]
    fn test_const_vs_runtime_keccak() {
        // Ensure compile-time keccak matches runtime syscall-based keccak
        let data = b"Withdraw(address user,bytes32 destination,bytes32 token,uint256 amount,uint256 chainId,uint64 nonce)";
        let const_result = keccak256_const(data);
        let runtime_result = keccak256(data);
        assert_eq!(const_result, runtime_result);
    }

    // Test helpers

    fn hex_decode_20(s: &str) -> [u8; 20] {
        let bytes = hex::decode(s).expect("invalid hex");
        let mut out = [0u8; 20];
        out.copy_from_slice(&bytes);
        out
    }

    fn address_to_bytes32(addr: &[u8; 20]) -> [u8; 32] {
        let mut out = [0u8; 32];
        out[12..].copy_from_slice(addr);
        out
    }
}
