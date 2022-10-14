package grpcsvc

import (
	"context"
	"net"
	"time"

	"github.com/ChengWu-NJ/yolosvc/pkg/pb"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
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

	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			// After a duration of this time if the server doesn't see any activity it
			// pings the client to see if the transport is still alive.
			// If set below 1s, a minimum value of 1s will be used instead.
			// The default value is 2 hours which is too long.
			Time: 60 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			// to set MinTime is for avoiding "too many pings" error
			// MinTime is the minimum amount of time a client should wait before sending
			// a keepalive ping. for example: when set 10 seconds, server will close
			// the connection if the client pings >= 2 in 10 seconds.
			// the minimum value of "Time" (check no active from server) of client is 10 seconds.
			MinTime:             10 * time.Second,
			PermitWithoutStream: true,
		}),
	)

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
