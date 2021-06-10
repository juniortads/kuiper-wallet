package qldb

import (
	"context"
	"errors"
	"github.com/amzn/ion-go/ion"
	"github.com/awslabs/amazon-qldb-driver-go/qldbdriver"
	"github.com/go-kit/kit/log"
	"github.com/juniortads/kuiper-wallet/services/transaction"
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

func (repo *repository) CreateTransaction(ctx context.Context, transact transaction.Transaction) (interface{}, error) {
	resp, err := repo.db.Execute(context.Background(), func(txn qldbdriver.Transaction) (interface{}, error){
		result, err := repo.checkIfThereIsTransactionByTrackingId(txn, transact.TrackingId)
		if err != nil {
			return nil, err
		}
		if result != "" {
			return result, nil
		} else {
			result, err := repo.checkIfThereIsBalanceAvailable(txn, transact.SourceAccountId,transact.TransactionValue.Amount)
			if err != nil {
				return nil, err
			}
			if result {
				resp, err := repo.addTransaction(txn, transact)
				if err != nil {
					return nil, err
				}
				err = repo.updateMetadataId(txn, resp, transact.Id)
				if err != nil {
					return nil, err
				}
				return transact.Id, nil
			}
			return nil, errors.New("insufficient balance")
		}
	})

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (repo *repository) addTransaction(txn qldbdriver.Transaction, transact transaction.Transaction) (interface{}, error) {
	resp, err := txn.Execute("INSERT INTO Transactions ?", transact)
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

func (repo *repository) updateMetadataId(txn qldbdriver.Transaction, documentId interface{}, transactionId string) error {
	_, err := txn.Execute("UPDATE Transactions SET MetadataID = ? WHERE Id = ?", documentId, transactionId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) checkIfThereIsTransactionByTrackingId(txn qldbdriver.Transaction, trackingId string)(string, error)  {
	result, err := txn.Execute("SELECT * FROM Transactions WHERE TrackingID = ?", trackingId)

	if err != nil {
		return "", err
	}
	if result.Next(txn) {
		ionBinary := result.GetCurrentData()
		temp := new(transaction.Transaction)
		err = ion.Unmarshal(ionBinary, temp)
		if err != nil {
			return "", err
		}
		return temp.Id, nil
	}
	return "", err
}

func (repo *repository) checkIfThereIsBalanceAvailable(txn qldbdriver.Transaction, accountId string, transactionValue *ion.Decimal)(bool, error)  {
	result, err := txn.Execute("SELECT Balance FROM AvailableBalance WHERE AccountId = ?", accountId)

	if err != nil {
		return false, err
	}
	if result.Next(txn) {
		ionBinary := result.GetCurrentData()
		temp := new(transaction.AvailableBalance)
		err = ion.Unmarshal(ionBinary, temp)
		if err != nil {
			return false, err
		}
		result := temp.Balance.Cmp(transactionValue)

		if result == 1 || result == 0 {
			return true, nil
		}
	}
	return false, err
}