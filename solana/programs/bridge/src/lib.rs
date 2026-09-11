use anchor_lang::prelude::*;
use anchor_spl::associated_token::AssociatedToken;
use anchor_spl::token::{self, Mint, Token, TokenAccount, TransferChecked};

pub mod state;
pub mod withdraw;

use state::*;
pub use withdraw::*;

declare_id!("C4YxxrnCKnE4hVdTPcmTZN6yuHp5U9xVXRs3VanEeYfq");

#[program]
pub mod bridge {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>, chain_id: u64) -> Result<()> {
        require!(
            chain_id == withdraw::CANONICAL_CHAIN_ID,
            BridgeError::ChainIdMismatch
        );

        let config = &mut ctx.accounts.config;
        config.admin = ctx.accounts.admin.key();
        config.paused = false;
        config.bump = ctx.bumps.config;
        config.chain_id = chain_id;
        config.withdraw_paused = false;
        config.verifying_contract = withdraw::VERIFYING_CONTRACT;
        config.domain_separator = withdraw::DOMAIN_SEPARATOR;

        Ok(())
    }

    pub fn initialize_vault(ctx: Context<InitializeVault>) -> Result<()> {
        let vault_state = &mut ctx.accounts.vault_state;
        vault_state.mint = ctx.accounts.mint.key();
        vault_state.bump = ctx.bumps.vault_state;
        Ok(())
    }

    pub fn deposit(ctx: Context<Deposit>, destination: [u8; 20], amount: u64) -> Result<()> {
        let config = &ctx.accounts.config;
        require!(!config.paused, BridgeError::Paused);
        require!(amount > 0, BridgeError::ZeroAmount);

        let mint = &ctx.accounts.mint;
        token::transfer_checked(
            CpiContext::new(
                Token::id(),
                TransferChecked {
                    from: ctx.accounts.user_token_account.to_account_info(),
                    mint: mint.to_account_info(),
                    to: ctx.accounts.vault_token_account.to_account_info(),
                    authority: ctx.accounts.user.to_account_info(),
                },
            ),
            amount,
            mint.decimals,
        )?;

        emit!(DepositEvent {
            user: ctx.accounts.user.key(),
            destination,
            mint: mint.key(),
            amount,
        });

        Ok(())
    }

    pub fn pause(ctx: Context<AdminAction>) -> Result<()> {
        ctx.accounts.config.paused = true;
        Ok(())
    }

    pub fn unpause(ctx: Context<AdminAction>) -> Result<()> {
        ctx.accounts.config.paused = false;
        Ok(())
    }

    pub fn set_validator_set(
        ctx: Context<SetValidatorSet>,
        epoch: u64,
        validators: Vec<Validator>,
    ) -> Result<()> {
        require!(!validators.is_empty(), BridgeError::NoValidators);
        require!(validators.len() <= 64, BridgeError::TooManyValidators);

        let vs = &mut ctx.accounts.validator_set;
        require!(epoch > vs.epoch, BridgeError::EpochNotIncreasing);

        let total_power: u64 = validators.iter().map(|v| v.power).sum();
        vs.epoch = epoch;
        vs.total_power = total_power;
        vs.validators = validators;
        // bump is set on first init; on realloc it stays the same
        if vs.bump == 0 {
            vs.bump = ctx.bumps.validator_set;
        }

        Ok(())
    }

    pub fn set_withdraw_paused(ctx: Context<AdminAction>, paused: bool) -> Result<()> {
        ctx.accounts.config.withdraw_paused = paused;
        Ok(())
    }

    pub fn withdraw(
        ctx: Context<Withdraw>,
        user: [u8; 20],
        destination: [u8; 32],
        amount: u64,
        nonce: u64,
        signatures: Vec<EcdsaSignature>,
    ) -> Result<()> {
        withdraw::handle_withdraw(ctx, user, destination, amount, nonce, signatures)
    }
}

// ---------------------------------------------------------------------------
// Account constraints
// ---------------------------------------------------------------------------

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(mut)]
    pub admin: Signer<'info>,

    #[account(
        init,
        payer = admin,
        space = 8 + BridgeConfig::INIT_SPACE,
        seeds = [b"config"],
        bump,
    )]
    pub config: Account<'info, BridgeConfig>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct InitializeVault<'info> {
    #[account(mut)]
    pub admin: Signer<'info>,

    #[account(
        seeds = [b"config"],
        bump = config.bump,
        has_one = admin @ BridgeError::Unauthorized,
    )]
    pub config: Account<'info, BridgeConfig>,

    pub mint: Account<'info, Mint>,

    #[account(
        init,
        payer = admin,
        space = 8 + VaultState::INIT_SPACE,
        seeds = [b"vault_state", mint.key().as_ref()],
        bump,
    )]
    pub vault_state: Account<'info, VaultState>,

    /// CHECK: PDA used as SPL token account authority
    #[account(
        seeds = [b"vault", mint.key().as_ref()],
        bump,
    )]
    pub vault_authority: UncheckedAccount<'info>,

    #[account(
        init,
        payer = admin,
        associated_token::mint = mint,
        associated_token::authority = vault_authority,
    )]
    pub vault_token_account: Account<'info, TokenAccount>,

    pub token_program: Program<'info, Token>,
    pub associated_token_program: Program<'info, AssociatedToken>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct Deposit<'info> {
    pub user: Signer<'info>,

    #[account(
        seeds = [b"config"],
        bump = config.bump,
    )]
    pub config: Account<'info, BridgeConfig>,

    pub mint: Account<'info, Mint>,

    #[account(
        seeds = [b"vault_state", mint.key().as_ref()],
        bump = vault_state.bump,
    )]
    pub vault_state: Account<'info, VaultState>,

    /// CHECK: PDA used as SPL token account authority
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

    #[account(
        mut,
        associated_token::mint = mint,
        associated_token::authority = user,
    )]
    pub user_token_account: Account<'info, TokenAccount>,

    pub token_program: Program<'info, Token>,
}

#[derive(Accounts)]
pub struct AdminAction<'info> {
    pub admin: Signer<'info>,

    #[account(
        mut,
        seeds = [b"config"],
        bump = config.bump,
        has_one = admin @ BridgeError::Unauthorized,
    )]
    pub config: Account<'info, BridgeConfig>,
}

#[derive(Accounts)]
pub struct SetValidatorSet<'info> {
    #[account(mut)]
    pub admin: Signer<'info>,

    #[account(
        seeds = [b"config"],
        bump = config.bump,
        has_one = admin @ BridgeError::Unauthorized,
    )]
    pub config: Account<'info, BridgeConfig>,

    #[account(
        init_if_needed,
        payer = admin,
        space = 8 + ValidatorSet::INIT_SPACE,
        seeds = [b"validator_set"],
        bump,
    )]
    pub validator_set: Account<'info, ValidatorSet>,

    pub system_program: Program<'info, System>,
}
