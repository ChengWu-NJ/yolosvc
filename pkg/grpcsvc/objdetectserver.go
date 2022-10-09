package grpcsvc

import (
	"context"
	"fmt"

	"github.com/ChengWu-NJ/yolosvc/pkg/config"
	"github.com/ChengWu-NJ/yolosvc/pkg/darknet"
	"github.com/ChengWu-NJ/yolosvc/pkg/pb"
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

	//s.detector.Detect()

	return nil, nil
}

func (s *server) DetectJpgStream(pb.ObjDetect_DetectJpgStreamServer) error {
	// TODO...
	return nil
}
