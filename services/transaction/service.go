package transaction

import "context"

type Service interface {
	CreateTransaction(ctx context.Context, transaction Transaction) (string, error)
}