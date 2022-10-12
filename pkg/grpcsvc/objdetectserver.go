package grpcsvc

import (
	"context"
	"fmt"
	"io"

	"github.com/ChengWu-NJ/yolosvc/pkg/config"
	"github.com/ChengWu-NJ/yolosvc/pkg/darknet"
	"github.com/ChengWu-NJ/yolosvc/pkg/pb"
	"github.com/gookit/slog"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Implements of EchoServiceServer

type server struct {
	ctx context.Context
	pb.ObjDetectServer
	detector *darknet.YOLONetwork
}

func newServer(ctx context.Context) (pb.ObjDetectServer, error) {
	detector, err := newDetector()
	if err != nil {
		return nil, err
	}

	return &server{
		ctx:      ctx,
		detector: detector,
	}, nil
}

func newDetector() (*darknet.YOLONetwork, error) {
	detector := &darknet.YOLONetwork{
		GPUDeviceIndex:           0,
		NetworkConfigurationFile: config.GlobalConfig.DarknetConfigFile,
		WeightsFile:              config.GlobalConfig.DarknetWeightsFile,
		Threshold:                config.GlobalConfig.DetectThreshold,
		ClassNames:               config.GlobalConfig.ClassNames,
		Classes:                  len(config.GlobalConfig.ClassNames),
	}

	if err := detector.Init(); err != nil {
		return nil, err
	}

	return detector, nil
}

func (s *server) Healthz(ctx context.Context, emp *emptypb.Empty) (*pb.HealthzResponse, error) {
	return &pb.HealthzResponse{
		State: "ok",
		Htime: timestamppb.Now(),
	}, nil
}

func (s *server) DetectOneJpg(ctx context.Context, jpgBytes *pb.JpgBytes) (*pb.JpgBytes, error) {
	if jpgBytes == nil || len(jpgBytes.JpgData) == 0 {
		return &pb.JpgBytes{}, fmt.Errorf(`got an empty JpgBytes argument`)
	}

	var err error
	jpgBytes.JpgData, err = s.detector.DetectAndLabelOnJpeg(jpgBytes.JpgData)

	return jpgBytes, err
}

func (s *server) DetectJpgStream(stream pb.ObjDetect_DetectJpgStreamServer) error {
	stmCtx := stream.Context()
	dataCh := make(chan *pb.JpgBytes, 100)
	defer close(dataCh) // take the initiative to release resources which could ease the GC

	exitCtx, exitCancel := context.WithCancel(s.ctx)
	defer exitCancel()

	// receiving loop goroutine
	go func() {
		defer exitCancel()

		for {
			select {
			case <-stmCtx.Done():
				return

			case <-exitCtx.Done():
				return

			default:
				jpgBytes, err := stream.Recv()
				if err != nil {
					if err == io.EOF {
						// closed by the client
						return
					}

					slog.Error(err)
					// tolerate other errors, and let stmCtx.Done() to deal with the network errors
					break // break out from select{}
				}

				if jpgBytes == nil || len(jpgBytes.JpgData) == 0 {
					break // break out to ignor wrong data
				}
				select {
				case dataCh <- jpgBytes:
				default: // just skip away if dataCh is full
				}
			}
		}
	}()

	// sending loop
	var err error
	for {
		select {
		case <-stmCtx.Done():
			if stmCtx.Err() == io.EOF {
				return nil
			}
			return stmCtx.Err()

		case <-exitCtx.Done():
			return nil

	case _jpgBytes := <-dataCh:
			_jpgBytes.JpgData, err = s.detector.DetectAndLabelOnJpeg(_jpgBytes.JpgData)
			if err != nil {
				slog.Error(err)
				break // break out select{}
			}

			// ignore error check, and let stmCtx.Done() to deal with possible network errors
			_ = stream.Send(_jpgBytes)
		}
	}
}
