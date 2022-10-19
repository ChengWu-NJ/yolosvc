package main

import (
	"bufio"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gookit/slog"

	"github.com/ChengWu-NJ/yolosvc/pkg/config"
	"github.com/ChengWu-NJ/yolosvc/pkg/darknet"
)

const (
	DARKNET_THRESHOLD = 0.8
)

var (
	n *darknet.YOLONetwork
)

func main() {

	n := &darknet.YOLONetwork{
		GPUDeviceIndex:           0,
		NetworkConfigurationFile: config.GlobalConfig.DarknetConfigFile,
		WeightsFile:              config.GlobalConfig.DarknetWeightsFile,
		Threshold:                config.GlobalConfig.DetectThreshold,
		ClassNames:               config.GlobalConfig.GetClassNames(),
		Classes:                  config.GlobalConfig.GetClassNumber(),
	}

	if err := n.Init(); err != nil {
		slog.Fatal(err)
		return
	}
	defer n.Close()

	batchFile, err := os.Open(`/apps/yolosvc/test.txt`)
	if err != nil {
		slog.Fatal(err)
		return
	}
	defer batchFile.Close()

	i := 0
	oks := 0
	scanner := bufio.NewScanner(batchFile)
	ts0 := time.Now()
	for scanner.Scan() {
		fnImg := scanner.Text()
		fnImg = strings.TrimSpace(fnImg)
		if fnImg == "" {
			continue
		}

		slog.Infof(`------- %000d. %s --------`, i, fnImg)
		if err := detectImgFile(fnImg); err != nil {
			slog.Error(err)
		} else {
			oks++
		}

		i++
	}
	ts1 := time.Now()

	timeLen := ts1.Sub(ts0)
	slog.Infof(`total %d images are detected in %v seconds. each image need %v`, oks, timeLen, timeLen/time.Duration(oks))

	/*
	   srcImg = n.DrawDetectionResult(srcImg, dr)
	   outBuf := &bytes.Buffer{}
	   err = jpeg.Encode(outBuf, srcImg, nil)

	   	if err != nil {
	   		printError(err)
	   		return
	   	}

	   err = os.WriteFile(`/dev/shm/bboxedImg.jpg`, outBuf.Bytes(), 0660)
	   log.Println(`save err:`, err)
	*/
}

func detectImgFile(fnImg string) error {
	infile, err := os.Open(fnImg)
	if err != nil {
		return err
	}
	defer infile.Close()

	srcImg, err := jpeg.Decode(infile)
	if err != nil {
		return err
	}

	imgDarknet, err := darknet.Image2Float32(srcImg)
	if err != nil {
		return err
	}
	defer imgDarknet.Close()

	dr, err := n.Detect(imgDarknet)
	if err != nil {
		return err
	}

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

	return nil
}
