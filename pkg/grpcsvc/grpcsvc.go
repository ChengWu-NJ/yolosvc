package grpcsvc


import (
	"context"
	"net"

	"google.golang.org/grpc"
	"github.com/ChengWu-NJ/yolosvc/pkg/pb"
	"github.com/gookit/slog"
)

// Run starts the gRPC service.
func Run(ctx context.Context, network, address string) error {
	slog.Infof("grpc starting listening at %s", address)

	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			slog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	ds, err := newServer(ctx)
	if err != nil {
		return err
	}

	pb.RegisterObjDetectServer(s, ds)

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()

	slog.Debug(`grpc starting to serve...`)
	return s.Serve(l)
}
