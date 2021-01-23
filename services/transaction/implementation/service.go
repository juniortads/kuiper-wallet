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
	uuid := ksuid.New()
	id := uuid.String()
	transaction.ID = id

	if err := s.repository.CreateTransaction(ctx, transaction); err != nil {
		level.Error(logger).Log("err", err)
		return err.Error(),err
	}
	return id, nil
}
