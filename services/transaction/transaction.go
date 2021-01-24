package transaction

import (
	"context"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID		         string `ion:"id"`
	Notes            string `ion:"notes"`
	TransactionValue Value  `ion:"transactionValue"`
	AccountID        string `ion:"accountID"`
	TrackingID       string `ion:"trackingID"`
	AccountingID     string `ion:"accountingID"`
	MetadataID		 string `ion:"metadataID"`
}

type Value struct {
	Currency string
	Amount decimal.Decimal
}

type Repository interface {
	CreateTransaction(ctx context.Context, transaction Transaction) (interface{}, error)
}
