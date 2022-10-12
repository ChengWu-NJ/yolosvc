package darknet

// #include <stdlib.h>
//
// #include <darknet.h>
//
// #include "network.h"
import "C"
import (
	"errors"
	"time"
	"unsafe"

	"github.com/ChengWu-NJ/yolosvc/pkg/config"
	"github.com/ChengWu-NJ/yolosvc/pkg/drawbbox"
)

// YOLONetwork represents a neural network using YOLO.
type YOLONetwork struct {
	GPUDeviceIndex           int
	NetworkConfigurationFile string
	WeightsFile              string
	Threshold                float32

	ClassNames []string
	Classes    int

	cNet                *C.network
	hierarchalThreshold float32
	nms                 float32

	labelAdder *drawbbox.LabelAdder
	boxColors  map[int]*labelColor
	txtColors  map[int]*labelColor
}

var (
	errNetworkNotInit      = errors.New("network not initialised")
	errUnableToInitNetwork = errors.New("unable to initialise")
)

// Init the network.
func (n *YOLONetwork) Init() error {
	nCfg := C.CString(n.NetworkConfigurationFile)
	defer C.free(unsafe.Pointer(nCfg))
	wFile := C.CString(n.WeightsFile)
	defer C.free(unsafe.Pointer(wFile))
	// GPU device ID must be set before `load_network()` is invoked.
	C.cuda_set_device(C.int(n.GPUDeviceIndex))
	n.cNet = C.load_network(nCfg, wFile, 0)
	if n.cNet == nil {
		return errUnableToInitNetwork
	}
	//C.srand(2222222)
	n.hierarchalThreshold = 0.5
	n.nms = 0.45
	//metadata := C.get_metadata(nCfg)
	//n.Classes = int(metadata.classes)
	//n.ClassNames = makeClassNames(metadata.names, n.Classes)

	n.labelAdder = drawbbox.NewLabelAdder()
	n.boxColors = make(map[int]*labelColor)
	n.txtColors = make(map[int]*labelColor)

	// set label colors of classes
	var oR, oG, oB, cR, cG, cB float64
	for _, cls := range config.GlobalConfig.ObjClasses{
		n.boxColors[cls.Id] = &labelColor{
			R: float64(cls.LabelColorR)/255.,
			G: float64(cls.LabelColorG)/255.,
			B: float64(cls.LabelColorB)/255.,
		}

		_, cR, cG, cB = drawbbox.GetConstrastColor(oR, oG, oB)
		n.txtColors[cls.Id] = &labelColor{
			R: cR,
			G: cG,
			B: cB,
		}
	}

	return nil
}

// Close and release resources.
func (n *YOLONetwork) Close() error {
	if n.cNet == nil {
		return errNetworkNotInit
	}
	C.free_network_ptr(n.cNet)
	n.cNet = nil
	return nil
}

// Detect specified image
func (n *YOLONetwork) Detect(img *DarknetImage) (*DetectionResult, error) {
	if n.cNet == nil {
		return nil, errNetworkNotInit
	}
	startTime := time.Now()
	result := C.perform_network_detect(n.cNet, &img.image, C.int(n.Classes), C.float(n.Threshold), C.float(n.hierarchalThreshold), C.float(n.nms), C.int(0))
	endTime := time.Now()
	ds := makeDetections(img, result.detections, int(result.detections_len), n.Threshold, n.Classes, n.ClassNames)
	C.free_detections(result.detections, result.detections_len)
	endTimeOverall := time.Now()
	out := DetectionResult{
		Detections:           ds,
		NetworkOnlyTimeTaken: endTime.Sub(startTime),
		OverallTimeTaken:     endTimeOverall.Sub(startTime),
	}
	return &out, nil
}
