use anchor_lang::prelude::*;
use anchor_spl::associated_token::AssociatedToken;
use anchor_spl::token::{self, Mint, Token, TokenAccount, TransferChecked};

use crate::state::*;

// ---------------------------------------------------------------------------
// Compile-time hex literal helper (must precede all uses)
// ---------------------------------------------------------------------------

/// Decode a fixed-length hex string at compile time.
macro_rules! hex_literal {
    ($hex:literal) => {{
        const N: usize = $hex.len() / 2;
        const fn hex_val(c: u8) -> u8 {
            match c {
                b'0'..=b'9' => c - b'0',
                b'a'..=b'f' => c - b'a' + 10,
                b'A'..=b'F' => c - b'A' + 10,
                _ => panic!("invalid hex character"),
            }
        }
        const fn decode<const M: usize>(hex: &[u8]) -> [u8; M] {
            let mut out = [0u8; M];
            let mut i = 0;
            while i < M {
                out[i] = hex_val(hex[2 * i]) << 4 | hex_val(hex[2 * i + 1]);
                i += 1;
            }
            out
        }
        decode::<N>($hex.as_bytes())
    }};
}

// ---------------------------------------------------------------------------
// Program-specific constants
// ---------------------------------------------------------------------------

/// Canonical chain ID for this program's EIP-712 domain.
/// Must match `SOLANA_CHAIN_ID` in `scripts/solana-e2e.sh` and relayer config.
pub const CANONICAL_CHAIN_ID: u64 = 900_001;

// ---------------------------------------------------------------------------
// EIP-712 typehash constants (precomputed, must byte-align with Go / EVM)
// ---------------------------------------------------------------------------

/// `keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)")`
const EIP712_DOMAIN_TYPEHASH: [u8; 32] =
    hex_literal!("8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f");

/// `keccak256("Withdraw(address user,bytes32 destination,bytes32 token,uint256 amount,uint256 chainId,uint64 nonce)")`
const WITHDRAW_TYPEHASH: [u8; 32] =
    hex_literal!("ae7dee9fe1cf9016b724a74908236e5f57a1789d05c09f2625fa9baee0cde49d");

/// `keccak256("Bridge")`
const NAME_HASH: [u8; 32] =
    hex_literal!("7aa5ae620294318af92bf4e2b2a729646c932a80312a5fa630da993a2ef5cc10");

/// `keccak256("1")`
const VERSION_HASH: [u8; 32] =
    hex_literal!("c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6");

/// secp256k1 half-order `n/2` for low-S enforcement.
const SECP256K1_HALF_ORDER: [u8; 32] = [
    0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
    0xFF, 0x5D, 0x57, 0x6E, 0x73, 0x57, 0xA4, 0x50, 0x1D, 0xDF, 0xE9, 0x2F, 0x46, 0x68, 0x1B,
    0x20, 0xA0,
];

// ---------------------------------------------------------------------------
// Runtime helpers (use Solana syscalls on-chain)
// ---------------------------------------------------------------------------

fn keccak256(data: &[u8]) -> [u8; 32] {
    solana_keccak_hasher::hash(data).to_bytes()
}

fn keccak256v(slices: &[&[u8]]) -> [u8; 32] {
    solana_keccak_hasher::hashv(slices).to_bytes()
}

/// EIP-712 `verifyingContract` pseudo-address: `keccak256(program_id)[12..32]`.
/// Must match Go `VerifyingContractFromProgramID`.
pub fn verifying_contract_from_program_id(program_id: &Pubkey) -> [u8; 20] {
    let hash = keccak256(program_id.as_ref());
    let mut out = [0u8; 20];
    out.copy_from_slice(&hash[12..32]);
    out
}

/// Bridge-domain EIP-712 separator for the deployed program ID and chain ID.
pub fn domain_separator_for_program(program_id: &Pubkey, chain_id: u64) -> [u8; 32] {
    let verifying_contract = verifying_contract_from_program_id(program_id);
    compute_domain_separator(chain_id, &verifying_contract)
}

fn compute_domain_separator(chain_id: u64, verifying_contract: &[u8; 20]) -> [u8; 32] {
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
    pub config: Box<Account<'info, BridgeConfig>>,

    #[account(
        seeds = [b"validator_set"],
        bump = validator_set.bump,
    )]
    pub validator_set: Box<Account<'info, ValidatorSet>>,

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
    pub vault_token_account: Box<Account<'info, TokenAccount>>,

    /// CHECK: Must match `destination` bytes (verified in handler).
    pub destination_owner: UncheckedAccount<'info>,

    #[account(
        init_if_needed,
        payer = payer,
        associated_token::mint = mint,
        associated_token::authority = destination_owner,
    )]
    pub destination_token_account: Box<Account<'info, TokenAccount>>,

    #[account(
        init_if_needed,
        payer = payer,
        space = 8 + NoncePage::INIT_SPACE,
        seeds = [b"nonce_page" as &[u8], &(nonce / 8192u64).to_le_bytes()],
        bump,
    )]
    pub nonce_page: Box<Account<'info, NoncePage>>,

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

    // 5. Defense-in-depth: stored config must match program constants
    require!(
        config.chain_id == CANONICAL_CHAIN_ID,
        BridgeError::ChainIdMismatch
    );
    let expected_verifying_contract = verifying_contract_from_program_id(&crate::ID);
    require!(
        config.verifying_contract == expected_verifying_contract,
        BridgeError::VerifyingContractMismatch
    );
    let expected_domain_separator =
        domain_separator_for_program(&crate::ID, CANONICAL_CHAIN_ID);
    require!(
        config.domain_separator == expected_domain_separator,
        BridgeError::DomainSeparatorMismatch
    );

    // 6. Nonce not already used
    let page_num = nonce / 8192;
    let bit_index = (nonce % 8192) as usize;
    let byte_index = bit_index / 8;
    let bit_mask = 1u8 << (bit_index % 8);

    let nonce_page = &ctx.accounts.nonce_page;
    require!(
        nonce_page.bits[byte_index] & bit_mask == 0,
        BridgeError::NonceAlreadyUsed
    );

    // 7. Compute EIP-712 digest using stored config (set at initialize from program ID)
    let token_bytes: [u8; 32] = ctx.accounts.mint.key().to_bytes();
    let struct_hash = compute_struct_hash(
        &user,
        &destination,
        &token_bytes,
        amount,
        CANONICAL_CHAIN_ID,
        nonce,
    );
    let digest = compute_digest(&config.domain_separator, &struct_hash);

    // 8–9. Verify signatures, accumulate power, check quorum
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

    // 10. Mark nonce as used
    let nonce_page = &mut ctx.accounts.nonce_page;
    nonce_page.page = page_num;
    nonce_page.bump = ctx.bumps.nonce_page;
    nonce_page.bits[byte_index] |= bit_mask;

    // 11. Transfer tokens from vault to destination
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

    // 12. Emit event for future relayer scanning
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
// Unit tests
// ---------------------------------------------------------------------------

#[cfg(test)]
mod tests {
    use super::*;

    // -- Precomputed literal assertions (verify hex literals match runtime keccak) --

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

    #[test]
    fn test_verifying_contract_from_declared_program_id() {
        let vc = verifying_contract_from_program_id(&crate::ID);
        let hash = keccak256(crate::ID.as_ref());
        assert_eq!(&vc, &hash[12..32]);
    }

    #[test]
    fn test_domain_separator_for_declared_program_id() {
        let ds = domain_separator_for_program(&crate::ID, CANONICAL_CHAIN_ID);
        let vc = verifying_contract_from_program_id(&crate::ID);
        let computed = compute_domain_separator(CANONICAL_CHAIN_ID, &vc);
        assert_eq!(ds, computed);
    }

    // -- Go conformance vectors (chain_id=1337, fake verifyingContract) --

    /// Conformance: Go test `TestBridgeHubWithdraw_ToTypedData` in
    /// `internal/contract/bridge_hub_test.go`.
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

        let ds = super::compute_domain_separator(chain_id, &verifying_contract);

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

        let ds = super::compute_domain_separator(chain_id, &verifying_contract);
        let sh = compute_struct_hash(&user, &destination, &token, amount, chain_id, nonce);
        let digest = compute_digest(&ds, &sh);

        assert_eq!(
            hex::encode(digest),
            "04f7234d8a5f15bd6cd22589a81bbff694d8b707f81df216635d52a3934a84e8"
        );
    }

    // -- Security --

    #[test]
    fn test_low_s_boundary() {
        let mut sig_at_half = [0u8; 64];
        sig_at_half[32..].copy_from_slice(&SECP256K1_HALF_ORDER);
        assert!(is_low_s(&sig_at_half));

        let mut sig_above = [0u8; 64];
        sig_above[32..].copy_from_slice(&SECP256K1_HALF_ORDER);
        sig_above[63] += 1;
        assert!(!is_low_s(&sig_above));

        let sig_zero = [0u8; 64];
        assert!(is_low_s(&sig_zero));
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
