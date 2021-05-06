package grpc

import (
	"context"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/juniortads/kuiper-wallet/services/account"
	"github.com/juniortads/kuiper-wallet/services/account/transport"
	"github.com/juniortads/kuiper-wallet/services/account/transport/pb"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	createAccount kitgrpc.Handler
	logger         log.Logger
}

func NewGRPCServer(
	endpoints transport.Endpoints, options []kitgrpc.ServerOption,
	logger log.Logger,
) pb.AccountServer {
	return &grpcServer{
		createAccount: kitgrpc.NewServer(
			endpoints.CreateAccount, decodeCreateAccountRequest, encodeCreateAccountResponse, options...,
		),
		logger: logger,
	}
}

func (s *grpcServer) CreateAccount(ctx oldcontext.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	_, rep, err := s.createAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateAccountResponse), nil
}

func decodeCreateAccountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateAccountRequest)

	return transport.CreateAccountRequest{
		Account: account.Account{
			Name:             req.Name,
			TrackingId:       req.TrackingId,
			ExternalId:       req.ExternalId,
			DocumentNumber:   req.DocumentNumber,
			CompanyId:        req.CompanyId,
			BallastAccountId: req.BallastAccountId,
		},
	}, nil
}

func encodeCreateAccountResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.CreateAccountResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.CreateAccountResponse{ Id: res.Id }, nil
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
