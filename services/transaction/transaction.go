package transaction

import (
	"context"
	"github.com/amzn/ion-go/ion"
)

type Transaction struct {
	ID		         string `ion:"id"`
	Notes            string `ion:"notes"`
	TransactionValue Value  `ion:"transactionValue"`
	AccountID        string `ion:"accountID"`
	TrackingID       string `ion:"trackingID"`
	AccountingID     string `ion:"accountingID"`
	MetadataID		 string `ion:"metadataID"`
	TransactionType  string `ion:"transactionType"`
}

type Value struct {
	Currency string `ion:"currency"`
	Amount *ion.Decimal `ion:"amount"`
}

type Repository interface {
	CreateTransaction(ctx context.Context, transaction Transaction) (interface{}, error)
}
