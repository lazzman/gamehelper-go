// 图色

package dm

import "github.com/go-ole/go-ole"

func (com *DaMo) AppendPicAddr(picInfo string, addr, size int) string {
	ret, _ := com.iDispatch.CallMethod("AppendPicAddr", picInfo, addr, size)
	return ret.ToString()
}

func (com *DaMo) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.iDispatch.CallMethod("Capture", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *DaMo) CaptureGif(x1, y1, x2, y2 int, file string, delay, time int) int {
	ret, _ := com.iDispatch.CallMethod("CaptureGif", x1, y1, x2, y2, file, delay, time)
	return int(ret.Val)
}

func (com *DaMo) CaptureJpg(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.iDispatch.CallMethod("CaptureJpg", x1, y1, x2, y2, file, quality)
	return int(ret.Val)
}

func (com *DaMo) CapturePng(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.iDispatch.CallMethod("CapturePng", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *DaMo) CapturePre(file string) int {
	ret, _ := com.iDispatch.CallMethod("CapturePre", file)
	return int(ret.Val)
}

func (com *DaMo) CmpColor(x int, y int, color string, sim float32) int {
	ret, _ := com.iDispatch.CallMethod("CmpColor", x, y, color, sim)
	return int(ret.Val)
}

func (com *DaMo) EnableDisplayDebug(enableDebug int) int {
	ret, _ := com.iDispatch.CallMethod("EnableDisplayDebug", enableDebug)
	return int(ret.Val)
}

func (com *DaMo) EnableFindPicMultithread(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableFindPicMultithread", enable)
	return int(ret.Val)
}

func (com *DaMo) EnableGetColorByCapture(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableGetColorByCapture", enable)
	return int(ret.Val)
}

func (com *DaMo) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindColor", x1, y1, x2, y2, color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DaMo) FindColorBlock(x1, y1, x2, y2 int, color string, sim float32, count, width, height int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindColorBlock", x1, y1, x2, y2, color, sim, count, width, height, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DaMo) FindColorBlockEx(x1, y1, x2, y2 int, color string, sim float32, count, width, height int) string {
	ret, _ := com.iDispatch.CallMethod("FindColorBlockEx", x1, y1, x2, y2, color, sim, count, width, height)
	return ret.ToString()
}

func (com *DaMo) FindColorE(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindColorE", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *DaMo) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindColorEx", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *DaMo) FindMulColor(x1, y1, x2, y2 int, color string, sim float32) int {
	ret, _ := com.iDispatch.CallMethod("FindMulColor", x1, y1, x2, y2, color, sim)
	return int(ret.Val)
}

func (com *DaMo) FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindMultiColor", x1, y1, x2, y2, firstColor, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DaMo) FindMultiColorE(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindMultiColorE", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *DaMo) FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindMultiColorEx", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *DaMo) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindPic", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DaMo) FindPicE(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindPicE", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *DaMo) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindPicEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *DaMo) FindPicExS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindPicExS", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

// func (com *DaMo)FindPicMem(x1, y1, x2, y2, pic_info, delta_color,sim, dir,intX, intY)  int{}
// func (com *DaMo)FindPicMemE(x1, y1, x2, y2, pic_info, delta_color,sim, dir,intX, intY)  string{}
// func (com *DaMo)FindPicMemEx(x1, y1, x2, y2, pic_info, delta_color,sim, dir,intX, intY)  string{}

func (com *DaMo) FindPicS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindPicS", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

// func (com *DaMo)FindShape(x1, y1, x2, y2, offset_color,sim, dir,intX,intY) int{}
// func (com *DaMo)FindShapeE(x1, y1, x2, y2, offset_color,sim, dir,intX,intY) string{}
// func (com *DaMo)FindShapeEx(x1, y1, x2, y2, offset_color,sim, dir,intX,intY) string{}
// func (com *DaMo)FreePic(pic_name) int{}

func (com *DaMo) GetAveHSV(x1, y1, x2, y2 int) string {
	ret, _ := com.iDispatch.CallMethod("GetAveHSV", x1, y1, x2, y2)
	return ret.ToString()
}

func (com *DaMo) GetAveRGB(x1, y1, x2, y2 int) string {
	ret, _ := com.iDispatch.CallMethod("GetAveRGB", x1, y1, x2, y2)
	return ret.ToString()
}

func (com *DaMo) GetColor(x, y int) string {
	ret, _ := com.iDispatch.CallMethod("GetColor", x, y)
	return ret.ToString()
}

// func (com *DaMo)GetColorBGR(x,y)string{}

func (com *DaMo) GetColorHSV(x, y int) string {
	ret, _ := com.iDispatch.CallMethod("GetColorHSV", x, y)
	return ret.ToString()
}

// func (com *DaMo)GetColorNum(x1, y1, x2, y2, color, sim) int{}

func (com *DaMo) GetPicSize(picName string) string {
	ret, _ := com.iDispatch.CallMethod("GetPicSize", picName)
	return ret.ToString()
}

// func (com *DaMo)GetScreenData(x1, y1, x2, y2) int{}
// func (com *DaMo)GetScreenDataBmp(x1,y1,x2,y2,data,size) int{}
// func (com *DaMo)ImageToBmp(pic_name,bmp_name) int{}

func (com *DaMo) IsDisplayDead(x1, y1, x2, y2, time int) int {
	ret, _ := com.iDispatch.CallMethod("IsDisplayDead", x1, y1, x2, y2, time)
	return int(ret.Val)
}

// func (com *DaMo)LoadPic(pic_name) int{}
// func (com *DaMo)LoadPicByte(addr,size,pic_name) int{}
// func (com *DaMo)LoadPic(pic_name) int{}

func (com *DaMo) MatchPicName(picName string) string {
	ret, _ := com.iDispatch.CallMethod("MatchPicName", picName)
	return ret.ToString()
}

func (com *DaMo) SetExcludeRegion(mode int, info string) int {
	ret, _ := com.iDispatch.CallMethod("SetExcludeRegion", mode, info)
	return int(ret.Val)
}

func (com *DaMo) SetPicPwd(pwd string) int {
	ret, _ := com.iDispatch.CallMethod("SetExcludeRegion", pwd)
	return int(ret.Val)
}
