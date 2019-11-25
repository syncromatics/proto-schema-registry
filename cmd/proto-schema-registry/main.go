package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/syncromatics/proto-schema-registry/internal/log"
	"github.com/syncromatics/proto-schema-registry/internal/service"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	settings, err := getSettingsFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	log.Info(settings.String())

	server := grpc.NewServer()
	service.RegisterService(server)

	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(hostServer(ctx, server, settings.Port))

	eventChan := make(chan os.Signal)
	signal.Notify(eventChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-eventChan:
	case <-ctx.Done():
	}

	cancel()

	if err := grp.Wait(); err != nil {
		panic(err)
	}
}

func hostServer(ctx context.Context, server *grpc.Server, port int) func() error {
	cancel := make(chan struct{})
	go func() {
		select {
		case <-ctx.Done():
			server.GracefulStop()
			return
		case <-cancel:
			return
		}
	}()

	return func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			close(cancel)
			return errors.Wrap(err, "failed to listen")
		}

		reflection.Register(server)

		log.Info("starting service", "port", port)
		if err := server.Serve(lis); err != nil {
			close(cancel)
			return errors.Wrap(err, "failed to serve")
		}
		log.Info("service stopped")

		return nil
	}
}
