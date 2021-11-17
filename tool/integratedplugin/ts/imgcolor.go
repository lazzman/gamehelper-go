// 图色

package ts

import "github.com/go-ole/go-ole"

func (com *TianShiPlug) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.iDispatch.CallMethod("Capture", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *TianShiPlug) CaptureGif(x1, y1, x2, y2 int, file string, delay, time int) int {
	ret, _ := com.iDispatch.CallMethod("CaptureGif", x1, y1, x2, y2, file, delay, time)
	return int(ret.Val)
}

func (com *TianShiPlug) CaptureJpg(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.iDispatch.CallMethod("CaptureJpg", x1, y1, x2, y2, file, quality)
	return int(ret.Val)
}

func (com *TianShiPlug) CapturePng(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.iDispatch.CallMethod("CapturePng", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *TianShiPlug) CmpColor(x int, y int, color string, sim float32) int {
	ret, _ := com.iDispatch.CallMethod("CmpColor", x, y, color, sim)
	return int(ret.Val)
}

func (com *TianShiPlug) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindColor", x1, y1, x2, y2, color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	if int(ret.Val) == -1 {
		return 0
	} else {
		return 1
	}
}

func (com *TianShiPlug) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindColorEx", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *TianShiPlug) FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindMultiColor", x1, y1, x2, y2, firstColor, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *TianShiPlug) FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindMultiColorEx", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *TianShiPlug) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindPic", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *TianShiPlug) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindPicEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *TianShiPlug) FindPicExS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindPicExS", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *TianShiPlug) FindPicS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindPicS", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *TianShiPlug) GetColor(x, y int) string {
	ret, _ := com.iDispatch.CallMethod("GetColor", x, y)
	return ret.ToString()
}

// func (com *TianShiPlug)GetScreenData(x1, y1, x2, y2) int{}

func (com *TianShiPlug) MatchPicName(picName string) string {
	ret, _ := com.iDispatch.CallMethod("MatchPicName", picName)
	return ret.ToString()
}

func (com *TianShiPlug) SetPicPwd(pwd string) int {
	ret, _ := com.iDispatch.CallMethod("SetPicPwd", pwd)
	return int(ret.Val)
}
