package wallet

import (
	"context"

	"github.com/google/uuid"
)

type WalletRepository interface {
	Add(ctx context.Context, wallet *Wallet) error

	Get(ctx context.Context, walletUUID uuid.UUID) (*Wallet, error)

	Update(
		ctx context.Context,
		walletUUID uuid.UUID,
		updateFn func(ctx context.Context, wallet *Wallet) (*Wallet, error),
	) error
}
