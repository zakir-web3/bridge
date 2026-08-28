// 引入 Anchor 框架常用类型、宏和 trait（Context、Result、require!、#[program] 等）
use anchor_lang::prelude::*;
// 引入 SPL Associated Token Account 程序类型，用于创建/校验 ATA
use anchor_spl::associated_token::AssociatedToken;
// 引入 SPL Token 程序：token 模块做 CPI 转账，Mint/TokenAccount 等是账户类型
use anchor_spl::token::{self, Mint, Token, TokenAccount, TransferChecked};

// 声明本 program 的 Program ID（公钥）；须与 Anchor.toml、部署 keypair 一致
declare_id!("C4YxxrnCKnE4hVdTPcmTZN6yuHp5U9xVXRs3VanEeYfq");

// 标记下方 mod 为 Anchor program：其中 pub fn 会暴露为链上可调用的 instruction
#[program]
pub mod bridge {
    // 引入父模块中的账户结构体、错误码、事件等
    use super::*;

    // 初始化全局桥配置账户（仅部署后调用一次）
    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        // 获取可变的 config 账户引用
        let config = &mut ctx.accounts.config;
        // 将调用者 admin 的公钥写入配置，作为后续管理员校验依据
        config.admin = ctx.accounts.admin.key();
        // 初始状态：桥未暂停，允许 deposit
        config.paused = false;
        // 保存 config PDA 的 bump seed，后续指令用 seeds+bump 验证同一 PDA
        config.bump = ctx.bumps.config;
        // 返回 Ok(()) 表示 instruction 成功
        Ok(())
    }

    // 为某一种 SPL mint 初始化金库（每种代币各调用一次）
    pub fn initialize_vault(ctx: Context<InitializeVault>) -> Result<()> {
        // 获取可变的 vault_state 账户引用
        let vault_state = &mut ctx.accounts.vault_state;
        // 记录该金库对应的 SPL mint 公钥
        vault_state.mint = ctx.accounts.mint.key();
        // 保存 vault_state PDA 的 bump，供后续 deposit 校验
        vault_state.bump = ctx.bumps.vault_state;
        Ok(())
    }

    // 用户 deposit：将 SPL 代币锁入金库，并发出跨链 relayer 可扫描的事件
    pub fn deposit(ctx: Context<Deposit>, destination: [u8; 20], amount: u64) -> Result<()> {
        // 只读引用全局配置
        let config = &ctx.accounts.config;
        // 若桥已暂停则拒绝 deposit
        require!(!config.paused, BridgeError::Paused);
        // 金额必须大于 0
        require!(amount > 0, BridgeError::ZeroAmount);

        // 引用 mint 账户，后续转账与事件需要 decimals 和 mint 地址
        let mint = &ctx.accounts.mint;
        // CPI 调用 SPL Token 的 transfer_checked（校验 decimals，比 transfer 更安全）
        token::transfer_checked(
            // 构造跨程序调用上下文：指定 token_program 与 TransferChecked 账户布局
            CpiContext::new(
                // SPL Token program 的 AccountInfo
                ctx.accounts.token_program.to_account_info(),
                TransferChecked {
                    // 转出方：用户的 SPL 代币账户
                    from: ctx.accounts.user_token_account.to_account_info(),
                    // 代币 mint 账户
                    mint: mint.to_account_info(),
                    // 转入方：金库 SPL 代币账户
                    to: ctx.accounts.vault_token_account.to_account_info(),
                    // 转出授权方：用户（须签名）
                    authority: ctx.accounts.user.to_account_info(),
                },
            ),
            // 转账数量（SPL 最小单位，如 6 decimals 下 100000 = 0.1 token）
            amount,
            // mint.decimals，transfer_checked 用于核对精度
            mint.decimals,
        )?; // ? 将 CPI 错误向上传播为 instruction 失败

        // 发出 Anchor 事件，序列化进交易 logs（relayer 解析 Program data）
        emit!(DepositEvent {
            user: ctx.accounts.user.key(),       // Solana 存款方公钥
            destination,                         // 目标链 EVM 收款地址（20 字节）
            mint: mint.key(),                    // 被锁定的 SPL mint
            amount,                              // 锁定数量
        });

        Ok(())
    }

    // 管理员暂停桥：禁止新的 deposit
    pub fn pause(ctx: Context<AdminAction>) -> Result<()> {
        ctx.accounts.config.paused = true;
        Ok(())
    }

    // 管理员恢复桥：允许 deposit
    pub fn unpause(ctx: Context<AdminAction>) -> Result<()> {
        ctx.accounts.config.paused = false;
        Ok(())
    }
}

// 为 Initialize 指令定义必须传入的账户及约束；Anchor 在进 handler 前自动校验
#[derive(Accounts)]
pub struct Initialize<'info> {
    // admin 须签名且账户可变（作为 init 的 payer 付 rent）
    #[account(mut)]
    pub admin: Signer<'info>,

    // 创建 config PDA 账户
    #[account(
        init,                                    // 新建账户（若已存在则失败）
        payer = admin,                           // rent 由 admin 支付
        space = 8 + BridgeConfig::INIT_SPACE,    // 8 字节账户 discriminator + 结构体大小
        seeds = [b"config"],                     // PDA seed：固定字节 "config"
        bump                                      // 自动选合法 bump 并写入 ctx.bumps.config
    )]
    pub config: Account<'info, BridgeConfig>,

    // Solana System Program，init 创建账户时需要
    pub system_program: Program<'info, System>,
}

// initialize_vault 指令的账户布局与约束
#[derive(Accounts)]
pub struct InitializeVault<'info> {
    #[account(mut)]
    pub admin: Signer<'info>,

    // 校验 config PDA，且 config.admin 必须等于本指令的 admin 签名者
    #[account(
        seeds = [b"config"],
        bump = config.bump,
        has_one = admin @ BridgeError::Unauthorized
    )]
    pub config: Account<'info, BridgeConfig>,

    // SPL 代币 mint 账户（只读校验类型）
    pub mint: Account<'info, Mint>,

    // 为该 mint 创建 vault_state PDA，标记「此 mint 金库已初始化」
    #[account(
        init,
        payer = admin,
        space = 8 + VaultState::INIT_SPACE,
        seeds = [b"vault_state", mint.key().as_ref()], // seed 含 mint，每种币一个
        bump
    )]
    pub vault_state: Account<'info, VaultState>,

    // 金库 authority PDA：无私钥，仅作为 SPL 金库账户的 owner
    #[account(
        seeds = [b"vault", mint.key().as_ref()],
        bump
    )]
    /// CHECK: PDA used as SPL token account authority
    pub vault_authority: UncheckedAccount<'info>, // 不解析内部数据，只验证是正确 PDA

    // 创建金库的 Associated Token Account，mint 与 authority 须匹配上面账户
    #[account(
        init,
        payer = admin,
        associated_token::mint = mint,
        associated_token::authority = vault_authority
    )]
    pub vault_token_account: Account<'info, TokenAccount>,

    pub token_program: Program<'info, Token>,
    pub associated_token_program: Program<'info, AssociatedToken>,
    pub system_program: Program<'info, System>,
}

// deposit 指令的账户布局与约束
#[derive(Accounts)]
pub struct Deposit<'info> {
    // 用户必须签名（作为 SPL 转出授权方）
    pub user: Signer<'info>,

    #[account(
        seeds = [b"config"],
        bump = config.bump
    )]
    pub config: Account<'info, BridgeConfig>,

    pub mint: Account<'info, Mint>,

    // 校验该 mint 对应的 vault_state 已初始化
    #[account(
        seeds = [b"vault_state", mint.key().as_ref()],
        bump = vault_state.bump
    )]
    pub vault_state: Account<'info, VaultState>,

    #[account(
        seeds = [b"vault", mint.key().as_ref()],
        bump
    )]
    /// CHECK: PDA used as SPL token account authority
    pub vault_authority: UncheckedAccount<'info>,

    // 金库代币账户：须可变，且 ATA 的 mint/authority 与 vault_authority 一致
    #[account(
        mut,
        associated_token::mint = mint,
        associated_token::authority = vault_authority
    )]
    pub vault_token_account: Account<'info, TokenAccount>,

    // 用户代币账户：须可变，且须属于 user 对该 mint 的 ATA
    #[account(
        mut,
        associated_token::mint = mint,
        associated_token::authority = user
    )]
    pub user_token_account: Account<'info, TokenAccount>,

    // deposit 通过 CPI 调用 SPL Token，须传入 Token program
    pub token_program: Program<'info, Token>,
}

// pause / unpause 共用账户约束
#[derive(Accounts)]
pub struct AdminAction<'info> {
    pub admin: Signer<'info>,

    #[account(
        mut,                                     // paused 字段需要写入
        seeds = [b"config"],
        bump = config.bump,
        has_one = admin @ BridgeError::Unauthorized
    )]
    pub config: Account<'info, BridgeConfig>,
}

// 标记为 Anchor 可序列化账户；链上 data 前 8 字节为 account:BridgeConfig discriminator
#[account]
// 自动计算 INIT_SPACE（各字段占用字节），供 init 时 space 使用
#[derive(InitSpace)]
pub struct BridgeConfig {
    pub admin: Pubkey,   // 32 字节管理员公钥
    pub paused: bool,    // 1 字节暂停标志
    pub bump: u8,        // 1 字节 PDA bump
}

#[account]
#[derive(InitSpace)]
pub struct VaultState {
    pub mint: Pubkey,    // 该金库对应的 SPL mint
    pub bump: u8,        // vault_state PDA 的 bump
}

// Anchor 事件：emit! 后写入 logs，discriminator 为 sha256("event:DepositEvent")[0..8]
#[event]
pub struct DepositEvent {
    pub user: Pubkey,           // 32 字节 Solana 存款用户
    pub destination: [u8; 20],  // 20 字节 EVM 目标地址
    pub mint: Pubkey,           // 32 字节 SPL mint
    pub amount: u64,            // 8 字节 little-endian 数量
}

// 自定义错误码；require! 失败时返回对应错误及 #[msg] 文案
#[error_code]
pub enum BridgeError {
    #[msg("Bridge is paused")]
    Paused,       // deposit 时 config.paused == true
    #[msg("Amount must be greater than zero")]
    ZeroAmount,   // amount == 0
    #[msg("Unauthorized")]
    Unauthorized, // 非 config.admin 调用管理员指令
}
