package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"github.com/awslabs/amazon-qldb-driver-go/qldbdriver"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/juniortads/kuiper-wallet/services/transaction"
	transactionsvc "github.com/juniortads/kuiper-wallet/services/transaction/implementation"
	"github.com/juniortads/kuiper-wallet/services/transaction/qldb"
	"github.com/juniortads/kuiper-wallet/services/transaction/transport"
	grpctransport "github.com/juniortads/kuiper-wallet/services/transaction/transport/grpc"
	"github.com/juniortads/kuiper-wallet/services/transaction/transport/pb"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	port = ":50051"
)

func main()  {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "transaction",
			"ts", log.DefaultTimestampUTC,
			"clr", log.DefaultCaller,
		)
	}
	awsSession := session.Must(session.NewSession(aws.NewConfig().WithRegion("us-east-1")))
	qldbSession := qldbsession.New(awsSession)

	driver, err := qldbdriver.New(
		"kuiper",
		qldbSession,
		func(options *qldbdriver.DriverOptions) {
			options.LoggerVerbosity = qldbdriver.LogInfo
		})
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	var svc transaction.Service
	{
		repository, err := qldb.New(driver, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = transactionsvc.NewService(repository, logger)
	}

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(svc)
	}
	var (
		//ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{}
		transactionService  = grpctransport.NewGRPCServer(endpoints, serverOptions, logger)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	var g group.Group
	{
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", port)
			pb.RegisterTransactionServer(grpcServer, transactionService)
			return grpcServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	level.Error(logger).Log("exit", g.Run())
}
