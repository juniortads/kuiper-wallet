package grpc

import (
	"context"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/juniortads/kuiper-wallet/services/transaction"
	"github.com/juniortads/kuiper-wallet/services/transaction/transport"
	"github.com/juniortads/kuiper-wallet/services/transaction/transport/pb"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
type grpcServer struct {
	createTransaction kitgrpc.Handler
	logger         log.Logger
}

func NewGRPCServer(
	endpoints transport.Endpoints, options []kitgrpc.ServerOption,
	logger log.Logger,
) pb.TransactionServer {
	return &grpcServer{
		createTransaction: kitgrpc.NewServer(
			endpoints.CreateTransaction, decodeCreateCustomerRequest, encodeCreateCustomerResponse, options...,
		),
		logger: logger,
	}
}

func (s *grpcServer) CreateTransaction(ctx oldcontext.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	_, rep, err := s.createTransaction.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateTransactionResponse), nil
}

func decodeCreateCustomerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateTransactionRequest)
	return transport.CreateTransactionRequest{
		Transaction: transaction.Transaction{
			ID:               req.Id,
			Notes:            req.Notes,
			TransactionValue: transaction.Value{
				Currency: req.TransactionValue.Currency,
				//Amount:   req.TransactionValue.Amount,
			},
			AccountID:        req.AccountId,
			TrackingID:       req.TrackingId,
		},
	}, nil
}

func encodeCreateCustomerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.CreateTransactionResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.CreateTransactionResponse{ Id: res.Id }, nil
	}
	return nil, err
}

func getError(err error) error {
	switch err {
	case nil:
		return nil
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}