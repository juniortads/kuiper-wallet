package transaction

import (
	"context"
	"github.com/amzn/ion-go/ion"
	"time"
)

type Transaction struct {
	Id		         	 string `ion:"Id"`
	TrackingId       	 string `ion:"TrackingId"`
	Notes            	 string `ion:"Notes"`
	TransactionValue 	 Value  `ion:"TransactionValue"`
	SourceAccountId      string `ion:"SourceAccountId"`
	BallastAccountId     string `ion:"BallastAccountId"`
	MetadataId		 	 string `ion:"MetadataId"`
	TransactionType      string `ion:"TransactionType"`
	DestinationHolder    string `ion:"DestinationHolder"`
	CreationDateTime     time.Time `ion:"CreationDateTime" type:"timestamp"`
}

type Value struct {
	Currency string `ion:"Currency"`
	Amount *ion.Decimal `ion:"Amount"`
}

type Repository interface {
	CreateTransaction(ctx context.Context, transaction Transaction) (interface{}, error)
}
