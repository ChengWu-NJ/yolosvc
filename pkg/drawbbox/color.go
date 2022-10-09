package drawbbox

import "math"

var (
	colors = [6][3]float64{{1, 0, 1}, {0, 0, 1}, {0, 1, 1}, {0, 1, 0}, {1, 1, 0}, {1, 0, 0}}
)

// R, G, B -- c: 2, 1, 0
func GetColor(c, x, max int) float64 {
	ratio := (float64(x) / float64(max)) * 5.
	n := int(math.Floor(ratio))
	ratio -= float64(n)
	i := n % 6
	j := (n + max/2) % 6

	r := (1.-ratio)*colors[i][c] + ratio*colors[j][c]
	// reduce to 0.07 ~ 0.93 to avoid pure colors
	r = r*0.86 + 0.07
	return r
}

// R, G, B -- 0 ~ 1
func luminanceOfRGB255(r255, g255, b255 float64) float64 {
	return 0.2126*cg(r255) + 0.7152*cg(g255) + 0.0722*cg(b255)
}

func cg(c float64) float64 {
	if c <= 10. {
		return c / 3294.
	}

	return math.Pow((c/269. + 0.0513), 2.4)
}

// https://ux.stackexchange.com/questions/82056/how-to-measure-the-contrast-between-any-given-color-and-white/82068#82068
func GetConstrastColor(r, g, b float64) (rate, fR, fG, fB float64) {
	r255, g255, b255 := r*255., g*255., b*255.
	origL := luminanceOfRGB255(r255, g255, b255)

	ir255, ig255, ib255 := int(r255), int(g255), int(b255)

	// 17*15 = 255: do 17 loops at step 15 on r, g, b
	iCstR255, iCstG255, iCstB255 := 0, 0, 0
	fCstR255, fCstG255, fCstB255 := float64(0.), float64(0.), float64(0.)
	var cstL, L1, L2 float64
	for stepR := 0; stepR < 17; stepR++ {
		for stepG := 0; stepG < 17; stepG++ {
			for stepB := 0; stepB < 17; stepB++ {
				iCstR255 = (ir255 + stepR*15) % 255
				iCstG255 = (ig255 + stepG*15) % 255
				iCstB255 = (ib255 + stepB*15) % 255

				fCstR255, fCstG255, fCstB255 = float64(iCstR255), float64(iCstG255), float64(iCstB255)
				fR, fG, fB = fCstR255/255., fCstG255/255., fCstB255/255.
				cstL = luminanceOfRGB255(fCstR255, fCstG255, fCstB255)
				if cstL > origL {
					L1 = cstL
					L2 = origL
				} else {
					L1 = origL
					L2 = cstL
				}

				rate = (L1 + 0.05) / (L2 + 0.05)
				if rate > 5. {
					return rate, fR, fG, fB
				}
			}
		}
	}

	return 999, 0., 0., 0. // black (0,0,0)
}
