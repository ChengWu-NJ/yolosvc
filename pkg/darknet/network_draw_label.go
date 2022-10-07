package darknet

import (
	"fmt"
	"image"
	"log"
	"math"

	"github.com/ChengWu-NJ/yolosvc/pkg/drawbbox"
)

type labelColor struct {
	R, G, B float64
}

var (
	colors = [6][3]float64{{1, 0, 1}, {0, 0, 1}, {0, 1, 1}, {0, 1, 0}, {1, 1, 0}, {1, 0, 0}}
)

// R, G, B -- c: 2, 1, 0
func getColor(c, x, max int) float64 {
	ratio := (float64(x) / float64(max)) * 5.
	fI := math.Floor(ratio)
	i := int(fI)
	j := int(math.Ceil(ratio))
	ratio -= fI
	return (1.-ratio)*colors[i][c] + ratio*colors[j][c]
}

// Draw detected results
func (n *YOLONetwork) DrawDetectionResult(img image.Image, dr *DetectionResult) image.Image {
	boxes, err := n.convertDetectionResultToBBOX(dr)
	log.Printf(`convertDetectionResultToBBOX [%#v], len[%d] err[%v]`, boxes,len(boxes), err)
	if err != nil {
		return img
	}

	return n.labelAdder.AddOnImage(img, boxes)
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
		log.Printf(`id[%d], label[%s]`, id, label)

		box := &drawbbox.BBox{
			Left:   float64(dt.StartPoint.X),
			Top:    float64(dt.StartPoint.Y),
			Right:  float64(dt.EndPoint.X),
			Bottom: float64(dt.EndPoint.Y),
			R:      n.labelColors[id].R,
			G:      n.labelColors[id].G,
			B:      n.labelColors[id].B,
			Label:  label,
		}
		boxes = append(boxes, box)
	}

	return boxes, nil
}

func getIDAndLabel(dt *Detection) (int, string) {
	iMax, pMax := 0, float32(0.0)

	log.Printf(`dt[%#v]`, dt)
	log.Printf(`id[%#v], classNames[%#v], prob[%#v]`, len(dt.ClassIDs), len(dt.ClassNames), len(dt.Probabilities))
	for i, p := range dt.Probabilities {
		if p > pMax {
			iMax, pMax = i, p
		}
	}

	log.Printf(`iMax[%d], pMax[%f]`, iMax, pMax)

	return iMax, fmt.Sprintf(`%s: %.2f`, dt.ClassNames[iMax], dt.Probabilities[iMax])
}
