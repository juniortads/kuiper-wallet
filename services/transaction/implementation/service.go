package implementation

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/juniortads/kuiper-wallet/services/transaction"
	"github.com/segmentio/ksuid"
)

type service struct {
	repository transaction.Repository
	logger     log.Logger
}

func NewService(rep transaction.Repository, logger log.Logger) *service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) CreateTransaction(ctx context.Context, transaction transaction.Transaction) (string, error) {
	logger := log.With(s.logger, "method", "CreateTransaction")
	transaction.ID = ksuid.New().String()

	//TODO: refactoring accounting id by account
	transaction.AccountingID = ksuid.New().String()

	resp, err := s.repository.CreateTransaction(ctx, transaction)

	if err != nil {
		level.Error(logger).Log("err", err)
		return err.Error(),err
	}
	return resp.(string), nil
}
