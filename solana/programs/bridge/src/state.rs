use anchor_lang::prelude::*;

// ---------------------------------------------------------------------------
// Accounts
// ---------------------------------------------------------------------------

/// Global bridge configuration (seeds: [b"config"]).
/// Extended from deposit-only with withdraw fields (chain_id, domain_separator,
/// verifying_contract, withdraw_paused).
#[account]
#[derive(InitSpace)]
pub struct BridgeConfig {
    pub admin: Pubkey,
    pub paused: bool,
    pub bump: u8,
    pub chain_id: u64,
    pub domain_separator: [u8; 32],
    pub verifying_contract: [u8; 20],
    pub withdraw_paused: bool,
}

/// Per-mint vault state (seeds: [b"vault_state", mint]).
#[account]
#[derive(InitSpace)]
pub struct VaultState {
    pub mint: Pubkey,
    pub bump: u8,
}

/// Single validator entry.
#[derive(AnchorSerialize, AnchorDeserialize, Clone, InitSpace)]
pub struct Validator {
    pub eth_address: [u8; 20],
    pub power: u64,
}

/// Validator set PDA (seeds: [b"validator_set"]).
/// `epoch` must strictly increase on every update.
#[account]
#[derive(InitSpace)]
pub struct ValidatorSet {
    pub epoch: u64,
    pub total_power: u64,
    pub bump: u8,
    #[max_len(64)]
    pub validators: Vec<Validator>,
}

/// Nonce replay-protection bitmap (seeds: [b"nonce_page", page.to_le_bytes()]).
/// Each page covers 8192 nonces (1024 bytes × 8 bits).
#[account]
#[derive(InitSpace)]
pub struct NoncePage {
    pub page: u64,
    pub bump: u8,
    pub bits: [u8; 1024],
}

// ---------------------------------------------------------------------------
// Events
// ---------------------------------------------------------------------------

#[event]
pub struct DepositEvent {
    pub user: Pubkey,
    pub destination: [u8; 20],
    pub mint: Pubkey,
    pub amount: u64,
}

#[event]
pub struct WithdrawFinalized {
    pub message: [u8; 32],
    pub user: [u8; 20],
    pub destination: [u8; 32],
    pub mint: Pubkey,
    pub amount: u64,
    pub nonce: u64,
}

// ---------------------------------------------------------------------------
// Errors
// ---------------------------------------------------------------------------

#[error_code]
pub enum BridgeError {
    #[msg("Bridge is paused")]
    Paused,
    #[msg("Amount must be greater than zero")]
    ZeroAmount,
    #[msg("Unauthorized")]
    Unauthorized,
    #[msg("Withdrawals are paused")]
    WithdrawPaused,
    #[msg("Mint does not match vault")]
    MintMismatch,
    #[msg("Destination does not match destination_owner pubkey")]
    DestinationMismatch,
    #[msg("Nonce already used")]
    NonceAlreadyUsed,
    #[msg("Insufficient quorum: 3 * power must exceed 2 * total_power")]
    InsufficientQuorum,
    #[msg("Signature recovery failed")]
    SignatureRecoveryFailed,
    #[msg("Recovered address is not in the validator set")]
    UnknownValidator,
    #[msg("Duplicate validator signature")]
    DuplicateSignature,
    #[msg("Signature has high-S value (must be <= secp256k1 half-order)")]
    HighSValue,
    #[msg("Epoch must be strictly greater than the current epoch")]
    EpochNotIncreasing,
    #[msg("Too many validators (max 64)")]
    TooManyValidators,
    #[msg("Must provide at least one validator")]
    NoValidators,
    #[msg("Validator set not initialized")]
    ValidatorSetNotInitialized,
    #[msg("chain_id must equal the canonical chain ID for this program")]
    ChainIdMismatch,
    #[msg("Stored domain separator does not match program constant")]
    DomainSeparatorMismatch,
    #[msg("Total validator power overflows u64")]
    TotalPowerOverflow,
    #[msg("Validator power must be greater than zero")]
    ZeroPower,
    #[msg("Duplicate eth_address in validator set")]
    DuplicateValidator,
}
