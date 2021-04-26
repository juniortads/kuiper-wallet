package qldb

import (
	"context"
	"github.com/amzn/ion-go/ion"
	"github.com/awslabs/amazon-qldb-driver-go/qldbdriver"
	"github.com/go-kit/kit/log"
	"github.com/juniortads/kuiper-wallet/services/account"
)

type repository struct {
	db *qldbdriver.QLDBDriver
	logger log.Logger
}

func New(db *qldbdriver.QLDBDriver, logger log.Logger) (*repository, error) {
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "amazonqldb"),
	}, nil
}

func (repo *repository) CreateAccount(ctx context.Context, acc account.Account) (interface{}, error) {
	resp, err := repo.db.Execute(context.Background(), func(txn qldbdriver.Transaction) (interface{}, error){
		result, err := repo.checkIfThereIsAccountByTrackingId(txn, acc.TrackingId)
		if err != nil {
			return nil, err
		}
		if result != "" {
			return result, nil
		} else {
			resp, err := repo.addAccount(txn, acc)
			if err != nil {
				return nil, err
			}
			err = repo.updateMetadataId(txn, "Account", resp, acc.Id)
			if err != nil {
				return nil, err
			}
			respAvailableBalance, err := repo.addAvailableBalance(txn, account.AvailableBalance{
				Balance:    ion.MustParseDecimal("0"),
				AccountId:  acc.Id,
				Currency:   "",
				MetadataId: "",
			})
			if err != nil {
				return nil, err
			}
			err = repo.updateMetadataId(txn, "AvailableBalance", respAvailableBalance, acc.Id)
			if err != nil {
				return nil, err
			}
			return acc.Id, nil
		}
	})

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (repo *repository) addAvailableBalance(txn qldbdriver.Transaction, availableBalance account.AvailableBalance) (interface{}, error) {
	resp, err := txn.Execute("INSERT INTO AvailableBalance ?", availableBalance)
	for resp.Next(txn) {
		var decoded map[string]interface{}
		err = ion.Unmarshal(resp.GetCurrentData(), &decoded)
		if err != nil {
			return nil, err
		}

		return decoded["documentId"], nil
	}
	return nil, err
}

func (repo *repository) checkIfThereIsAccountByTrackingId(txn qldbdriver.Transaction, trackingId string)(string, error)  {
	result, err := txn.Execute("SELECT * FROM Account WHERE trackingID = ?", trackingId)

	if err != nil {
		return "", err
	}
	if result.Next(txn) {
		ionBinary := result.GetCurrentData()
		temp := new(account.Account)
		err = ion.Unmarshal(ionBinary, temp)
		if err != nil {
			return "", err
		}
		return temp.Id, nil
	}
	return "", err
}

func (repo *repository) addAccount(txn qldbdriver.Transaction, acc interface{}) (interface{}, error) {
	resp, err := txn.Execute("INSERT INTO Account ?", acc)
	for resp.Next(txn) {
		var decoded map[string]interface{}
		err = ion.Unmarshal(resp.GetCurrentData(), &decoded)
		if err != nil {
			return nil, err
		}

		return decoded["documentId"], nil
	}
	return nil, err
}

func (repo *repository) updateMetadataId(txn qldbdriver.Transaction, table string, documentId interface{}, accountId string) error {
	_, err := txn.Execute("UPDATE "+table+" SET metadataID = ? WHERE id = ?", documentId, accountId)
	if err != nil {
		return err
	}
	return nil
}