package transaction

import (
	"context"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	Add(ctx context.Context, ts *Transaction) error

	Get(ctx context.Context, transactionUUID uuid.UUID) (*Transaction, error)

	Update(
		ctx context.Context,
		transactionUUID uuid.UUID,
		updateFn func(ctx context.Context, ts *Transaction) (*Transaction, error),
	) error
}
