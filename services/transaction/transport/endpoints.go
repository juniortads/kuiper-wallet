package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/juniortads/kuiper-wallet/services/transaction"
)

type Endpoints struct {
	CreateTransaction endpoint.Endpoint
}

func MakeEndpoints(s transaction.Service) Endpoints {
	return Endpoints{
		CreateTransaction: makeCreateTransactionEndpoint(s),
	}
}

func makeCreateTransactionEndpoint(s transaction.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTransactionRequest)
		id, err := s.CreateTransaction(ctx, req.Transaction)

		return CreateTransactionResponse{ Err: err, Id: id}, nil
	}
}