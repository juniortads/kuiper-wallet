package implementation

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/juniortads/kuiper-wallet/services/account"
	"github.com/segmentio/ksuid"
	"time"
)

const (
	Ok = "Ok"
	Blocked = "Blocked"
	Cancelled = "Cancelled"
)

type service struct {
	repository account.Repository
	logger     log.Logger
}

func NewService(rep account.Repository, logger log.Logger) *service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) CreateAccount(ctx context.Context, account account.Account) (string, error) {
	logger := log.With(s.logger, "method", "CreateAccount")
	account.Id = ksuid.New().String()
	account.CreationDate = time.Now().String()
	account.State = Ok

	resp, err := s.repository.CreateAccount(ctx, account)

	if err != nil {
		level.Error(logger).Log("err", err)
		return err.Error(),err
	}
	return resp.(string), nil
}
