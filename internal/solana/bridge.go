package solana

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/zakir-web3/bridge/internal/contract"
)

const withdrawEIP712Type = "Withdraw(address user,bytes32 destination,bytes32 token,uint256 amount,uint256 chainId,uint64 nonce)"

// DepositConfirmer submits Solana deposit confirmations to the target chain hub.
type DepositConfirmer interface {
	DepositConfirmSolana(ctx context.Context, srcChainID *big.Int, deposits ...*contract.SolanaDeposit) error
}

// HubContract is the bridgehub surface required for Solana withdraw relay.
type HubContract interface {
	DepositConfirmer
	SubmitSignatures(ctx context.Context, chainId *big.Int, signatures ...contract.CrossChainMessage) error
	FinalizedWithdrawal(ctx context.Context, chainId *big.Int, withdrawals ...*contract.BridgeFinalizedWithdrawal) error
	GetValidatorSet() *contract.ValidatorSet
}

// TypedDataSigner signs EIP-712 payloads with the validator secp256k1 key.
type TypedDataSigner interface {
	SignTypedData(typedData apitypes.TypedData) (contract.Signature, error)
}

type Bridge struct {
	logger            zerolog.Logger
	cfg               Config
	client            *Client
	chainID           *big.Int
	hub               HubContract
	signer            TypedDataSigner
	validatorSet      contract.ValidatorSet
	verifyingContract common.Address
	domainSeparator   common.Hash
	withdrawTypeHash  []byte
}

func NewBridge(ctx context.Context, cfg Config, hub HubContract, signer TypedDataSigner) (*Bridge, error) {
	client, err := NewClient(cfg.NodeURL, cfg.CommitmentType(), &cfg.RetryConfig, nil)
	if err != nil {
		return nil, err
	}
	logger := log.With().
		Str("module", "solana").
		Uint64("chainId", cfg.ChainID.Uint64()).
		Logger()

	validatorSet := hub.GetValidatorSet()
	if validatorSet == nil {
		return nil, errors.New("validator set unavailable")
	}

	var programIDBytes [32]byte
	copy(programIDBytes[:], cfg.ProgramID.Bytes())

	verifyingContract := contract.VerifyingContractFromProgramID(programIDBytes)
	domainSeparator, err := contract.SolanaBridgeDomainSeparator(cfg.ChainID, programIDBytes)
	if err != nil {
		return nil, errors.Wrap(err, "compute solana bridge domain separator")
	}

	logger.Info().
		Str("programId", cfg.ProgramID.String()).
		Str("verifyingContract", verifyingContract.Hex()).
		Hex("domainSeparator", domainSeparator.Bytes()).
		Bool("enableWithdraw", cfg.EnableWithdraw).
		Msg("initialized Solana bridge")

	return &Bridge{
		logger:            logger,
		cfg:               cfg,
		client:            client,
		chainID:           cfg.ChainID,
		hub:               hub,
		signer:            signer,
		validatorSet:      *validatorSet,
		verifyingContract: verifyingContract,
		domainSeparator:   domainSeparator,
		withdrawTypeHash:  crypto.Keccak256([]byte(withdrawEIP712Type)),
	}, nil
}

func (b *Bridge) GetChainID() *big.Int {
	return b.chainID
}

func (b *Bridge) GetClient() *Client {
	return b.client
}

func (b *Bridge) GetProgramID() solana.PublicKey {
	return b.cfg.ProgramID
}

func (b *Bridge) Withdraw(ctx context.Context, _ *big.Int, withdraws ...*contract.BridgeHubWithdraw) error {
	signatures := make([]contract.CrossChainMessage, 0, len(withdraws))
	var programIDBytes [32]byte
	copy(programIDBytes[:], b.cfg.ProgramID.Bytes())

	for _, w := range withdraws {
		if w.ChainId != nil && w.ChainId.Cmp(b.chainID) != 0 {
			continue
		}
		typedData := contract.BridgeHubWithdrawSolanaTypedData(w, b.chainID, programIDBytes)
		domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
		if err != nil {
			return errors.Wrap(err, "hash domain separator for solana withdraw")
		}
		if !bytes.Equal(domainSeparator, b.domainSeparator.Bytes()) {
			return errors.Errorf("domain separator mismatch: expected %s, got %s",
				b.domainSeparator.Hex(), domainSeparator.String())
		}
		messageRawData, err := typedData.EncodeData(typedData.PrimaryType, typedData.Message, 1)
		if err != nil {
			return errors.Wrap(err, "encode data for solana withdraw")
		}
		signature, err := b.signer.SignTypedData(typedData)
		if err != nil {
			return err
		}
		signatures = append(signatures, contract.CrossChainMessage{
			DomainSeparator: common.BytesToHash(domainSeparator),
			Signature:       signature,
			MessageRawData:  messageRawData,
		})
		b.logger.Info().
			Str("user", w.User.Hex()).
			Uint64("nonce", w.Nonce).
			Msg("solana withdraw signature created")
	}
	if len(signatures) == 0 {
		return nil
	}
	return b.hub.SubmitSignatures(ctx, b.chainID, signatures...)
}

func (b *Bridge) BridgeSignatureSubmitted(ctx context.Context, _ *big.Int, signs ...*contract.MessageSignature) error {
	if !b.cfg.EnableWithdraw {
		b.logger.Info().Msg("solana withdraw is disabled, skipping")
		return nil
	}
	if len(b.cfg.FeePayerKey) == 0 {
		return errors.New("fee_payer_key is required to submit solana withdraw transactions")
	}

	for _, sign := range signs {
		if !b.haveEnoughPower(sign.TotalPower) {
			b.logger.Debug().
				Uint64("totalPower", sign.TotalPower).
				Msg("not enough power for solana withdraw")
			continue
		}
		if len(sign.RawData) < 32 || !bytes.Equal(sign.RawData[:32], b.withdrawTypeHash) {
			continue
		}

		user, destination, token, amount, chainID, nonce, err := contract.ParseWithdrawMessageRawData(sign.RawData)
		if err != nil {
			return errors.Wrap(err, "parse solana withdraw message")
		}
		if chainID.Cmp(b.chainID) != 0 {
			b.logger.Debug().
				Str("expectedChainId", b.chainID.String()).
				Str("actualChainId", chainID.String()).
				Msg("withdraw chain id mismatch, skipping")
			continue
		}

		mint := TokenBytesToMint(token)
		if !b.cfg.IsBridgeMint(mint) {
			b.logger.Warn().Str("mint", mint.String()).Msg("skipping non-whitelisted withdraw mint")
			continue
		}

		splAmount, err := AmountToUint64(amount)
		if err != nil {
			return errors.Wrap(err, "convert withdraw amount")
		}

		ecdsaSigs, err := b.validatorSignatures(sign)
		if err != nil {
			return errors.Wrap(err, "collect validator signatures")
		}
		if len(ecdsaSigs) == 0 {
			b.logger.Warn().Uint64("nonce", nonce).Msg("no non-empty validator signatures for solana withdraw")
			continue
		}

		var userBytes [20]byte
		copy(userBytes[:], user.Bytes())
		params := WithdrawParams{
			User:        userBytes,
			Destination: DestinationBytesToPubkey(destination),
			Mint:        mint,
			Amount:      splAmount,
			Nonce:       nonce,
			Signatures:  ecdsaSigs,
		}

		if err := b.sendWithdraw(ctx, params); err != nil {
			return errors.Wrap(err, "send solana withdraw")
		}
		b.logger.Info().
			Str("user", user.Hex()).
			Str("mint", mint.String()).
			Str("destination", params.Destination.String()).
			Uint64("amount", splAmount).
			Uint64("nonce", nonce).
			Msg("solana withdraw transaction submitted")
	}
	return nil
}

func (b *Bridge) ProcessTransaction(ctx context.Context, sig solana.Signature, tx *rpc.GetTransactionResult) error {
	if finalized, err := ParseWithdrawFinalizedFromTransaction(b.cfg.ProgramID, tx, sig); err == nil {
		return b.handleWithdrawFinalized(ctx, finalized)
	}

	parsed, err := ParseDepositFromTransaction(b.cfg.ProgramID, tx, sig)
	if err != nil {
		b.logger.Debug().
			Err(err).
			Str("signature", sig.String()).
			Msg("skip non-bridge Solana transaction")
		return nil
	}
	mintPub := solana.PublicKeyFromBytes(parsed.Token[:])
	if !b.cfg.IsBridgeMint(mintPub) {
		b.logger.Warn().Str("mint", mintPub.String()).Msg("skipping non-whitelisted mint")
		return nil
	}
	parsed.ChainID = b.chainID
	b.logger.Info().
		Str("signature", sig.String()).
		Str("mint", mintPub.String()).
		Str("destination", parsed.Destination.Hex()).
		Str("amount", parsed.Amount.String()).
		Uint64("blockNumber", parsed.BlockNumber).
		Msg("solana deposit detected")
	return b.hub.DepositConfirmSolana(ctx, b.chainID, &parsed.SolanaDeposit)
}

func (b *Bridge) handleWithdrawFinalized(ctx context.Context, parsed *ParsedWithdrawFinalized) error {
	mintPub := TokenBytesToMint(parsed.Token)
	if !b.cfg.IsBridgeMint(mintPub) {
		b.logger.Warn().Str("mint", mintPub.String()).Msg("skipping non-whitelisted withdraw finalized mint")
		return nil
	}
	b.logger.Info().
		Str("signature", parsed.Signature.String()).
		Str("user", parsed.User.Hex()).
		Str("mint", mintPub.String()).
		Uint64("nonce", parsed.Nonce).
		Str("amount", parsed.Amount.String()).
		Msg("solana withdraw finalized detected")
	return b.hub.FinalizedWithdrawal(ctx, b.chainID, parsed.ToBridgeFinalizedWithdrawal())
}

func (b *Bridge) sendWithdraw(ctx context.Context, params WithdrawParams) error {
	withdrawIx, err := BuildWithdrawInstruction(b.cfg.ProgramID, b.cfg.FeePayerKey.PublicKey(), params)
	if err != nil {
		return err
	}
	instructions := PrependComputeBudgetInstructions(len(params.Signatures), b.cfg.ComputeUnitPrice)
	instructions = append(instructions, withdrawIx)

	blockhash, err := b.client.GetLatestBlockhash(ctx)
	if err != nil {
		return err
	}

	solTx, err := solana.NewTransaction(
		instructions,
		blockhash,
		solana.TransactionPayer(b.cfg.FeePayerKey.PublicKey()),
	)
	if err != nil {
		return errors.Wrap(err, "build solana withdraw transaction")
	}

	_, err = solTx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key.Equals(b.cfg.FeePayerKey.PublicKey()) {
			return &b.cfg.FeePayerKey
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "sign solana withdraw transaction")
	}

	if b.cfg.NoSend {
		b.logger.Info().Msg("no_send enabled, skipping solana withdraw broadcast")
		return nil
	}

	sig, err := b.client.SendTransaction(ctx, solTx)
	if err != nil {
		return err
	}
	b.logger.Info().Str("signature", sig.String()).Msg("solana withdraw transaction sent")
	return nil
}

func (b *Bridge) validatorSignatures(sign *contract.MessageSignature) ([]EcdsaSignature, error) {
	rawSigs := contract.NonEmptyValidatorSignatures(sign.Signatures)
	out := make([]EcdsaSignature, 0, len(rawSigs))
	for _, sig := range rawSigs {
		converted, err := SignatureToEcdsa(sig)
		if err != nil {
			return nil, err
		}
		out = append(out, converted)
	}
	return out, nil
}

func (b *Bridge) haveEnoughPower(cumulativePower uint64) bool {
	totalPower := uint64(0)
	for _, power := range b.validatorSet.Powers {
		totalPower += power
	}
	return 3*cumulativePower > 2*totalPower
}
