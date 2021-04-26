package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/juniortads/kuiper-wallet/services/account"
)

type Endpoints struct {
	CreateAccount endpoint.Endpoint
}

func MakeEndpoints(s account.Service) Endpoints {
	return Endpoints{
		CreateAccount: makeCreateAccountEndpoint(s),
	}
}

func makeCreateAccountEndpoint(s account.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		id, err := s.CreateAccount(ctx, req.Account)

		return CreateAccountResponse{ Err: err, Id: id}, nil
	}
}
