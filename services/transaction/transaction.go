package transaction

import (
	"context"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID, Notes        string
	TransactionValue Value
	AccountID        string
	TrackingID       string
	AccountingID     string
}

type Value struct {
	Currency string
	Amount decimal.Decimal
}

type Repository interface {
	CreateTransaction(ctx context.Context, transaction Transaction) error
}
