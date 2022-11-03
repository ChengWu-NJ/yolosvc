package darknet

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/ChengWu-NJ/yolosvc/pkg/drawbbox"
	"github.com/ChengWu-NJ/yolosvc/pkg/pb"
	"github.com/gookit/slog"
)

type labelColor struct {
	R, G, B float64
}

func (n *YOLONetwork) DetectAndLabelOnJpeg(jpgBytes *pb.JpgBytes, outputCh *chan *pb.JpgBytes) {
	buf := bytes.NewBuffer(jpgBytes.JpgData)

	srcImg, err := jpeg.Decode(buf)
	if err != nil {
		slog.Trace(err)
		*outputCh <- jpgBytes
		return
	}

	imgDarknet, err := Image2Float32(srcImg)
	if err != nil {
		slog.Trace(err)
		*outputCh <- jpgBytes
		return
	}
	defer imgDarknet.Close()

	results, err := n.Detect(imgDarknet)
	if err != nil {
		slog.Trace(err)
		*outputCh <- jpgBytes
		return
	}

	srcImg = n.DrawDetectionResult(srcImg, results, jpgBytes.SrcTs)

	if err := jpeg.Encode(buf, srcImg, nil); err != nil {
		slog.Trace(err)
		*outputCh <- jpgBytes
		return
	}

	*outputCh <- &pb.JpgBytes{
		SrcID:   jpgBytes.SrcID,
		SrcTs:   jpgBytes.SrcTs,
		JpgData: buf.Bytes(),
	}
}

// Draw detected results
func (n *YOLONetwork) DrawDetectionResult(img image.Image, dr *DetectionResult, nowTs int64) image.Image {
	boxes, err := n.convertDetectionResultToBBOX(dr)
	if err != nil {
		return img
	}

	return n.labelAdder.AddOnImage(img, boxes, nowTs)
}

func (n *YOLONetwork) convertDetectionResultToBBOX(dr *DetectionResult) ([]*drawbbox.BBox, error) {
	if dr == nil || dr.Detections == nil {
		return nil, fmt.Errorf(`arguments nil`)
	}

	boxes := make([]*drawbbox.BBox, 0)

	for _, dt := range dr.Detections {
		if dt == nil || dt.ClassIDs == nil || len(dt.ClassIDs) == 0 ||
			dt.ClassNames == nil || len(dt.ClassNames) == 0 ||
			dt.Probabilities == nil || len(dt.Probabilities) == 0 {
			continue
		}

		id, label := getIDAndLabel(dt)
		slog.Printf(`id[%d], label[%s]`, id, label)

		box := &drawbbox.BBox{
			Left:   float64(dt.StartPoint.X),
			Top:    float64(dt.StartPoint.Y),
			Right:  float64(dt.EndPoint.X),
			Bottom: float64(dt.EndPoint.Y),
			BoxR:   n.boxColors[id].R,
			BoxG:   n.boxColors[id].G,
			BoxB:   n.boxColors[id].B,
			TxtR:   n.txtColors[id].R,
			TxtG:   n.txtColors[id].G,
			TxtB:   n.txtColors[id].B,
			Label:  label,
		}
		boxes = append(boxes, box)
	}

	return boxes, nil
}

func getIDAndLabel(dt *Detection) (int, string) {
	iMax, pMax := 0, float32(0.0)

	for i, p := range dt.Probabilities {
		if p > pMax {
			iMax, pMax = i, p
		}
	}

	return dt.ClassIDs[iMax], fmt.Sprintf(`%s: %.2f`, dt.ClassNames[iMax], dt.Probabilities[iMax])
}
