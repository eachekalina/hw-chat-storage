package grpc

import (
	"context"
	"fmt"
	"github.com/central-university-dev/2024-spring-ab-go-hw-3-eachekalina/internal/pb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
)

func Run(ctx context.Context, addr string, msgHandler pb.MessageStorageServer, userHandler pb.UserStorageServer) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	server := grpc.NewServer()
	pb.RegisterMessageStorageServer(server, msgHandler)
	pb.RegisterUserStorageServer(server, userHandler)

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := server.Serve(lis)
		if err != nil {
			return fmt.Errorf("server: %w", err)
		}
		return nil
	})

	<-ctx.Done()
	server.GracefulStop()

	return eg.Wait()
}
