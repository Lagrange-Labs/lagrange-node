package network

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/Lagrange-Labs/Lagrange-Node/logger"
	"github.com/Lagrange-Labs/Lagrange-Node/network/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func RunServer(cfg *ServerConfig, storage storageInterface) error {
	ctx := context.Background()

	if len(cfg.GRPCPort) == 0 {
		errMsg := fmt.Sprintf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
		logger.WithError(errors.New(errMsg)).Error("Failed to start gRPC server")
		return errors.New(errMsg)
	}

	sequencerService, err := NewSequencerService(storage)
	if err != nil {
		return err
	}

	go func() {
		_ = runGRPCServer(ctx, sequencerService, cfg.GRPCPort)
	}()

	return nil
}

// HealthChecker will provide an implementation of the HealthCheck interface.
type healthChecker struct{}

// NewHealthChecker returns a health checker according to standard package
// grpc.health.v1.
func newHealthChecker() *healthChecker {
	return &healthChecker{}
}

// HealthCheck interface implementation.

// Check returns the current status of the server for unary gRPC health requests,
// for now if the server is up and able to respond we will always return SERVING.
func (s *healthChecker) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch returns the current status of the server for stream gRPC health requests,
// for now if the server is up and able to respond we will always return SERVING.
func (s *healthChecker) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

func runGRPCServer(ctx context.Context, svc types.NetworkServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	types.RegisterNetworkServiceServer(server, svc)

	healthService := newHealthChecker()
	grpc_health_v1.RegisterHealthServer(server, healthService)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	logger.Info("gRPC Server is serving at ", port)
	return server.Serve(listen)
}
