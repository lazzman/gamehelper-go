// 文字识别

package lw

func (com *LeWan) FindStr(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	ret, _ := com.iDispatch.CallMethod("FindStr", x1, y1, x2, y2, str, colorFormat, sim)
	*intX = com.GetX()
	*intY = com.GetY()
	return int(ret.Val)
}

func (com *LeWan) FindStrEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *LeWan) FindStrExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrExS", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *LeWan) FindStrFast(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	ret, _ := com.iDispatch.CallMethod("FindStrFast", x1, y1, x2, y2, str, colorFormat, sim)
	*intX = com.GetX()
	*intY = com.GetY()
	return int(ret.Val)
}

func (com *LeWan) FindStrFastEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrFastEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *LeWan) Ocr(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("Ocr", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *LeWan) SetDict(index int, file string) int {
	ret, _ := com.iDispatch.CallMethod("SetDict", index, file)
	return int(ret.Val)
}
