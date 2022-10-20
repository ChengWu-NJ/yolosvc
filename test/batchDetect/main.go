package main

import (
	"bufio"
	"context"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gookit/slog"

	"github.com/ChengWu-NJ/yolosvc/pkg/darknet"
	"github.com/ChengWu-NJ/yolosvc/pkg/grpcsvc"
)

const (
	DARKNET_THRESHOLD = 0.8
)

func main() {
	ctx := context.Background()

	n, err := grpcsvc.NewDetector()
	if err != nil {
		slog.Error(err)
		return
	}

	if n == nil {
		slog.Error(`n is nil`)
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

		_ts0 := time.Now()
		if fnImg == "" {
			continue
		}

		slog.Infof(`------- %d. %s --------`, i, fnImg)
		if err := detectImgFile(fnImg, n); err != nil {
			slog.Error(err)
		}
		oks++

		_ts1 := time.Now()
		_timeLen := _ts1.Sub(_ts0)
		slog.Infof(`%d. %s detection spends %v seconds`, i, fnImg, _timeLen)

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

	<-ctx.Done()
}

func detectImgFile(fnImg string, n *darknet.YOLONetwork) error {
	if n == nil {
		return fmt.Errorf(`n is nil`)
	}

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
