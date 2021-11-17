// 图色

package lw

func (com *LeWan) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.iDispatch.CallMethod("Capture", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *LeWan) CmpColor(x int, y int, color string, sim float32) int {
	ret, _ := com.iDispatch.CallMethod("CmpColor", x, y, color, sim)
	return int(ret.Val)
}

func (com *LeWan) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int {
	ret, _ := com.iDispatch.CallMethod("FindColor", x1, y1, x2, y2, color, sim, dir)
	*intX = com.GetX()
	*intY = com.GetY()
	return int(ret.Val)
}

func (com *LeWan) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindColorEx", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *LeWan) FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	ret, _ := com.iDispatch.CallMethod("FindMultiColor", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	*intX = com.GetX()
	*intY = com.GetY()
	return int(ret.Val)
}

func (com *LeWan) FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindMultiColorEx", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *LeWan) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	ret, _ := com.iDispatch.CallMethod("FindPic", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	*intX = com.GetX()
	*intY = com.GetY()
	return int(ret.Val)
}

func (com *LeWan) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.iDispatch.CallMethod("FindPicEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *LeWan) GetColor(x, y int) string {
	ret, _ := com.iDispatch.CallMethod("GetColor", x, y)
	return ret.ToString()
}

// func (com *LeWan)GetScreenData(x1, y1, x2, y2) int{}

func (com *LeWan) MatchPicName(picName string) string {
	ret, _ := com.iDispatch.CallMethod("MatchPicName", picName)
	return ret.ToString()
}

func (com *LeWan) SetPicPwd(pwd string) int {
	ret, _ := com.iDispatch.CallMethod("SetPicPwd", pwd)
	return int(ret.Val)
}
