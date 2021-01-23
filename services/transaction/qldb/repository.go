package qldb

import (
	"context"
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

func (repo *repository) CreateTransaction(ctx context.Context, transaction transaction.Transaction) error {
	_, err := repo.db.Execute(context.Background(), func(txn qldbdriver.Transaction) (interface{}, error){
		result, err := txn.Execute("SELECT * FROM WalletTransaction WHERE TrackingID = ?", transaction.TrackingID)

		if err != nil {
			return nil, err
		}
		if result.Next(txn) {
			// Document already exists, no need to insert
		}else {
			transaction := map[string]interface{}{
				"ID": transaction.ID,
				"TrackingID": transaction.TrackingID,
				"TransactionValue": transaction.TransactionValue,
				"AccountID": transaction.AccountID,
				"Notes": transaction.Notes,
			}
			_, err = txn.Execute("INSERT INTO WalletTransaction ?", transaction)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})

	if err != nil {
		return err
	}
	return nil
}
