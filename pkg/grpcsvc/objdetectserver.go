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
	Detector *darknet.YOLONetwork
}

func NewServer(ctx context.Context) (pb.ObjDetectServer, error) {
	return newServer(ctx)
}

func newServer(ctx context.Context) (pb.ObjDetectServer, error) {
	detector, err := NewDetector()
	if err != nil {
		return nil, err
	}

	return &server{
		ctx:      ctx,
		Detector: detector,
	}, nil
}

func NewDetector() (*darknet.YOLONetwork, error) {
	detector := &darknet.YOLONetwork{
		GPUDeviceIndex:           0,
		NetworkConfigurationFile: config.GlobalConfig.DarknetConfigFile,
		WeightsFile:              config.GlobalConfig.DarknetWeightsFile,
		Threshold:                config.GlobalConfig.DetectThreshold,
		ClassNames:               config.GlobalConfig.GetClassNames(),
		Classes:                  config.GlobalConfig.GetClassNumber(),
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

	outputCh := make(chan *pb.JpgBytes, 1)
	defer close(outputCh)

	var err error
	s.Detector.DetectAndLabelOnJpeg(jpgBytes, &outputCh)

	return <-outputCh, err
}

func (s *server) DetectJpgStream(stream pb.ObjDetect_DetectJpgStreamServer) error {
	stmCtx := stream.Context()
	labelCh := make(chan *pb.JpgBytes, 100)
	outputCh := make(chan *pb.JpgBytes, 100)
	defer func() {
		// take the initiative to release resources which could ease the GC
		close(outputCh)
		close(labelCh)
	}()

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
				case labelCh <- jpgBytes:
				default: // just skip away if dataCh is full
				}
			}
		}
	}()

	// label routine
	go func() {
		defer exitCancel()

		for {
			select {
			case <-stmCtx.Done():
				return

			case <-exitCtx.Done():
				return

			case _jpgBytes := <-labelCh:
				go s.Detector.DetectAndLabelOnJpeg(_jpgBytes, &outputCh)
			}
		}
	}()

	// sending loop
	for {
		select {
		case <-stmCtx.Done():
			if stmCtx.Err() == io.EOF {
				return nil
			}
			return stmCtx.Err()

		case <-exitCtx.Done():
			return nil

		case _jpgBytes := <-outputCh:
			_ = stream.Send(_jpgBytes)
		}
	}
}
