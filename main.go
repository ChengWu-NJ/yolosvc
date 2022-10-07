package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/ChengWu-NJ/yolosvc/pkg/darknet"
)

const (
	DARKNET_CONFIG    = `/home/ubuntu/devapps/yolo_dataset/yolov7_test.cfg`
	DARKNET_WEIGHTS   = `/home/ubuntu/devapps/yolo_dataset/backup/yolov7_final.weights`
	DARKNET_THRESHOLD = 0.8
)

func printError(err error) {
	log.Println("error:", err)
}

func main() {

	n := darknet.YOLONetwork{
		GPUDeviceIndex:           0,
		NetworkConfigurationFile: DARKNET_CONFIG,
		WeightsFile:              DARKNET_WEIGHTS,
		Threshold:                DARKNET_THRESHOLD,
		ClassNames: []string{
			`yellowGourami`,
			`blueGourami`,
			`redFighter`,
			`originalFighter`,
			`fox`,
			`sucker`,
			`tiger`,
			`WCMM`,
		},
		Classes: 8,
	}
	if err := n.Init(); err != nil {
		printError(err)
		return
	}
	defer n.Close()

	infile, err := os.Open(`/home/ubuntu/devapps/yolo_dataset/img/valid/09061104.jpg`)
	if err != nil {
		panic(err.Error())
	}
	defer infile.Close()

	srcImg, err := jpeg.Decode(infile)
	if err != nil {
		panic(err.Error())
	}

	imgDarknet, err := darknet.Image2Float32(srcImg)
	if err != nil {
		panic(err.Error())
	}
	defer imgDarknet.Close()

	dr, err := n.Detect(imgDarknet)
	if err != nil {
		printError(err)
		return
	}

	log.Println("Network-only time taken:", dr.NetworkOnlyTimeTaken)
	log.Println("Overall time taken:", dr.OverallTimeTaken, len(dr.Detections))
	for _, d := range dr.Detections {
		for i := range d.ClassIDs {
			bBox := d.BoundingBox
			fmt.Printf("%s (%d): %.4f%% | start point: (%d,%d) | end point: (%d, %d)\n",
				d.ClassNames[i], d.ClassIDs[i],
				d.Probabilities[i],
				bBox.StartPoint.X, bBox.StartPoint.Y,
				bBox.EndPoint.X, bBox.EndPoint.Y,
			)
		}
	}

	srcImg = n.DrawDetectionResult(srcImg, dr)
	outBuf := &bytes.Buffer{}
	err = jpeg.Encode(outBuf, srcImg, nil)
	if err != nil {
		printError(err)
		return
	}

	err = os.WriteFile(`/dev/shm/bboxedImg.jpg`, outBuf.Bytes(), 0660)	
	log.Println(`save err:`, err)
}
