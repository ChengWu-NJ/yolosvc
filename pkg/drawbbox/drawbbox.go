package drawbbox

import (
	"bytes"
	"image"
	"image/jpeg"
	"math"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

type LabelAdder struct {
	dc   *gg.Context
	font *truetype.Font
}

func NewLabelAdder() *LabelAdder {
	adder := &LabelAdder{}

	var err error
	adder.font, err = truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}

	return adder
}

type BBox struct {
	Left, Top, Right, Bottom float64
	BoxR, BoxG, BoxB         float64
	TxtR, TxtG, TxtB         float64
	Label                    string
}

func (a *LabelAdder) AddOnJpgBytes(jpgBytes []byte, boxes []*BBox) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewBuffer(jpgBytes))
	if err != nil {
		return nil, err
	}

	img = a.AddOnImage(img, boxes)

	outBuf := &bytes.Buffer{}
	err = jpeg.Encode(outBuf, img, nil)
	if err != nil {
		return nil, err
	}

	return outBuf.Bytes(), nil
}

func (a *LabelAdder) AddOnImage(img image.Image, boxes []*BBox) image.Image {
	imgHigh := img.Bounds().Dy()
	fontHeight := float64(imgHigh) / 1000. * 18.
	a.dc = gg.NewContextForImage(img)
	a.dc.SetLineWidth(fontHeight / 10.)

	for _, b := range boxes {
		a.dc.SetRGBA(b.BoxR, b.BoxG, b.BoxR, 1.)
		// draw box
		a.dc.DrawRectangle(b.Left, b.Top, math.Abs(b.Right-b.Left), math.Abs(b.Bottom-b.Top))
		a.dc.Stroke()
		a.dc.ClearPath()

		// draw background of label
		//a.dc.SetRGBA(b.BoxR, b.BoxG, b.BoxR, 0.5)
		bgW, bgH := a.dc.MeasureString(b.Label)
		bgW, bgH = (bgW+2.)*1.1, bgH+2.
		x_bg, y_bg := b.Left-fontHeight/20., b.Top-bgH-fontHeight/20
		a.dc.DrawRoundedRectangle(x_bg, y_bg, bgW, bgH, fontHeight/8)
		a.dc.Fill()
		a.dc.ClearPath()

		// draw text of label
		face := truetype.NewFace(a.font, &truetype.Options{Size: fontHeight})
		a.dc.SetFontFace(face)
		x_txt, y_txt := b.Left+1., b.Top-1-fontHeight/20.

		a.dc.SetRGBA(b.TxtR, b.TxtG, b.TxtB, 1.)
		a.dc.DrawString(b.Label, x_txt, y_txt)
	}

	return a.dc.Image()
}
