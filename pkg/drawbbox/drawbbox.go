package drawbbox

import (
	"bytes"
	"image"
	"image/jpeg"
	"math"
	"time"

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

func (a *LabelAdder) AddOnJpgBytes(jpgBytes []byte, boxes []*BBox, nowTs int64) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewBuffer(jpgBytes))
	if err != nil {
		return nil, err
	}

	img = a.AddOnImage(img, boxes, nowTs)

	outBuf := &bytes.Buffer{}
	err = jpeg.Encode(outBuf, img, nil)
	if err != nil {
		return nil, err
	}

	return outBuf.Bytes(), nil
}

func (a *LabelAdder) AddOnImage(img image.Image, boxes []*BBox, nowTs int64) image.Image {
	nowLabel := time.Unix(0, nowTs).Format("2006-01-02 15:04:05.00") + " UTC"

	imgHigh := img.Bounds().Dy()
	// adopt fontHeight as the unit
	fontHeight := float64(imgHigh) / 1000. * 18.
	face0 := truetype.NewFace(a.font, &truetype.Options{Size: fontHeight * 2})
	face1 := truetype.NewFace(a.font, &truetype.Options{Size: fontHeight})

	a.dc = gg.NewContextForImage(img)
	a.dc.SetLineWidth(fontHeight / 10.)

	// draw current time label
	a.dc.SetFontFace(face0)
	a.dc.SetRGB(0.5, 0.5, 0.5)
	nowLableWidth, nowLabelHeight := a.dc.MeasureString(nowLabel)
	// from (1.8, 1.8) with (w+0.2, h+0.2)
	a.dc.DrawRoundedRectangle(1.8*fontHeight, 1.8*fontHeight,
		nowLableWidth+0.4*fontHeight, nowLabelHeight+0.4*fontHeight, 0.1*fontHeight)
	a.dc.Fill()
	a.dc.ClearPath()
	a.dc.SetRGB(0.9, 0.9, 0.9)
	a.dc.DrawString(nowLabel, 2.*fontHeight, 2.8*fontHeight)

	for _, b := range boxes {
		a.dc.SetRGB(b.BoxR, b.BoxG, b.BoxR)
		// draw box
		a.dc.DrawRectangle(b.Left, b.Top, math.Abs(b.Right-b.Left), math.Abs(b.Bottom-b.Top))
		a.dc.Stroke()
		a.dc.ClearPath()

		// draw background of label
		//a.dc.SetRGBA(b.BoxR, b.BoxG, b.BoxR, 0.5)
		a.dc.SetFontFace(face1)
		bgW, bgH := a.dc.MeasureString(b.Label)
		bgW, bgH = (bgW+2.)*1.1, bgH+2.
		x_bg, y_bg := b.Left-fontHeight/20., b.Top-bgH-fontHeight/20
		a.dc.DrawRoundedRectangle(x_bg, y_bg, bgW, bgH, fontHeight/8)
		a.dc.Fill()
		a.dc.ClearPath()

		// draw text of label
		x_txt, y_txt := b.Left+1., b.Top-1-fontHeight/20.

		a.dc.SetRGB(b.TxtR, b.TxtG, b.TxtB)
		a.dc.DrawString(b.Label, x_txt, y_txt)
	}

	return a.dc.Image()
}
