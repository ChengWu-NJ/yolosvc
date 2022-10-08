package drawbbox

import "testing"

func TestColors(t *testing.T) {
	Classes := 8
	for classID := 0; classID < Classes; classID++ {
		offset := (classID * 123457) % Classes
		R := GetColor(2, offset, Classes)
		G := GetColor(1, offset, Classes)
		B := GetColor(0, offset, Classes)
		t.Logf(`%d - oR[%f] oG[%f] oB[%f]`, classID, R, G, B)

		rate, cR, cG, cB := GetConstrastColor(R,G,B)
		t.Logf(`%d - cR[%f] cG[%f] cB[%f] r[%f]`, classID, cR, cG, cB, rate)
	}
}

/* test output sample:
0 - oR[0.930000] oG[0.070000] oB[0.930000]
0 - cR[0.047059] cG[0.066667] cB[0.047059] r[5.371975]
1 - oR[0.392500] oG[0.607500] oB[0.930000]
1 - cR[0.392157] cG[0.015686] cB[0.047059] r[4.757698]
2 - oR[0.715000] oG[0.070000] oB[0.285000]
2 - cR[0.713725] cG[0.890196] cB[0.517647] r[4.504064]
3 - oR[0.177500] oG[0.070000] oB[0.822500]
3 - cR[0.176471] cG[0.772549] cB[0.819608] r[4.711427]
4 - oR[0.930000] oG[0.500000] oB[0.500000]
4 - cR[0.047059] cG[0.027451] cB[0.498039] r[5.889958]
5 - oR[0.177500] oG[0.822500] oB[0.070000]
5 - cR[0.176471] cG[0.054902] cB[0.066667] r[8.721551]
6 - oR[0.715000] oG[0.285000] oB[0.070000]
6 - cR[0.713725] cG[0.988235] cB[0.831373] r[4.526691]
7 - oR[0.392500] oG[0.930000] oB[0.607500]
7 - cR[0.392157] cG[0.047059] cB[0.603922] r[6.866886]
*/