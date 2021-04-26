package account

import (
	"context"
	"github.com/amzn/ion-go/ion"
	"time"
)

type Account struct {
	Id               string          `ion:"Id"`
	Name             string          `ion:"Name"`
	TrackingId       string          `ion:"TrackingId"`
	ExternalId       string          `ion:"ExternalId"`
	DocumentNumber   string          `ion:"DocumentNumber"`
	CompanyId        string          `ion:"CompanyId"`
	BallastAccountId string          `ion:"BallastAccountId"`
	MetadataId       string          `ion:"MetadataId"`
	State            string          `ion:"State"`
	CreationDateTime time.Time       `ion:"CreationDateTime" type:"timestamp"`

}
type AvailableBalance struct {
	Balance    *ion.Decimal `ion:"Balance"`
	AccountId  string       `ion:"AccountId"`
	Currency   string       `ion:"Currency"`
	MetadataId string       `ion:"MetadataId"`
}

type Repository interface {
	CreateAccount(ctx context.Context, account Account) (interface{}, error)
}
