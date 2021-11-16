// 文字识别

package dm

import "github.com/go-ole/go-ole"

// AddDict 给指定的字库中添加一条字库信息
func (com *DaMo) AddDict(index int, dictInfo string) int {
	ret, _ := com.iDispatch.CallMethod("AddDict", index, dictInfo)
	return int(ret.Val)
}

// ClearDict 清空指定的字库
func (com *DaMo) ClearDict(index int) int {
	ret, _ := com.iDispatch.CallMethod("ClearDict", index)
	return int(ret.Val)
}

func (com *DaMo) EnableShareDict(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableShareDict", enable)
	return int(ret.Val)
}

func (com *DaMo) FetchWord(x1, y1, x2, y2 int, color, word string) string {
	ret, _ := com.iDispatch.CallMethod("FetchWord", x1, y1, x2, y2, color, word)
	return ret.ToString()
}

func (com *DaMo) FindStr(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindStr", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DaMo) FindStrE(x1, y1, x2, y2 int, str string, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrE", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) FindStrEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) FindStrExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrExS", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) FindStrFast(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindStrFast", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DaMo) FindStrFastE(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrFastE", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) FindStrFastEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrFastEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) FindStrFastExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrFastExS", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) FindStrFastS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindStrFastS", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *DaMo) FindStrS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindStrS", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *DaMo) FindStrWithFont(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, fontName string, fontSize, flag int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindStrWithFont", x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DaMo) FindStrWithFontE(x1, y1, x2, y2 int, str, colorFormat string, sim float32, fontName string, fontSize, flag int) string {
	ret, _ := com.iDispatch.CallMethod("FindStrWithFontE", x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *DaMo) FindStrWithFontEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32, fontName string, fontSize, flag int) string {
	ret, _ := com.iDispatch.CallMethod("FindStrWithFontEx", x1, y1, x2, y2, str, colorFormat, sim, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *DaMo) GetDict(index, fontIndex int) string {
	ret, _ := com.iDispatch.CallMethod("GetDict", index, fontIndex)
	return ret.ToString()
}

func (com *DaMo) GetDictCount(index int) int {
	ret, _ := com.iDispatch.CallMethod("GetDictCount", index)
	return int(ret.Val)
}

func (com *DaMo) GetDictInfo(str, fontName string, fontSize, flag int) string {
	ret, _ := com.iDispatch.CallMethod("GetDictInfo", str, fontName, fontSize, flag)
	return ret.ToString()
}

func (com *DaMo) GetNowDict() int {
	ret, _ := com.iDispatch.CallMethod("GetNowDict")
	return int(ret.Val)
}

func (com *DaMo) GetResultCount(str string) int {
	ret, _ := com.iDispatch.CallMethod("GetResultCount", str)
	return int(ret.Val)
}

func (com *DaMo) GetResultPos(str string, index int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("GetResultPos", str, index, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DaMo) GetWordResultCount(str string) int {
	ret, _ := com.iDispatch.CallMethod("GetWordResultCount", str)
	return int(ret.Val)
}

func (com *DaMo) GetWordResultPos(str string, index int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("GetWordResultPos", str, index, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DaMo) GetWordResultStr(str string, index int) string {
	ret, _ := com.iDispatch.CallMethod("GetWordResultCount", str, index)
	return ret.ToString()
}

func (com *DaMo) GetWords(x1, y1, x2, y2 int, color string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("GetWords", x1, y1, x2, y2, color, sim)
	return ret.ToString()
}

func (com *DaMo) GetWordsNoDict(x1, y1, x2, y2 int, color string) string {
	ret, _ := com.iDispatch.CallMethod("GetWordsNoDict", x1, y1, x2, y2, color)
	return ret.ToString()
}

func (com *DaMo) Ocr(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("Ocr", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) OcrEx(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("OcrEx", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) OcrExOne(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("OcrExOne", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) OcrInFile(x1, y1, x2, y2 int, picName, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("OcrInFile", x1, y1, x2, y2, picName, colorFormat, sim)
	return ret.ToString()
}

func (com *DaMo) SaveDict(index int, file string) int {
	ret, _ := com.iDispatch.CallMethod("SaveDict", index, file)
	return int(ret.Val)
}

func (com *DaMo) SetColGapNoDict(colGap int) int {
	ret, _ := com.iDispatch.CallMethod("SetColGapNoDict", colGap)
	return int(ret.Val)
}

func (com *DaMo) SetDict(index int, file string) int {
	ret, _ := com.iDispatch.CallMethod("SetDict", index, file)
	return int(ret.Val)
}

func (com *DaMo) SetDictMem(index, addr, size int) int {
	ret, _ := com.iDispatch.CallMethod("SetDictMem", index, addr, size)
	return int(ret.Val)
}

func (com *DaMo) SetDictPwd(pwd string) int {
	ret, _ := com.iDispatch.CallMethod("SetDictPwd", pwd)
	return int(ret.Val)
}

func (com *DaMo) SetExactOcr(exactOcr int) int {
	ret, _ := com.iDispatch.CallMethod("SetExactOcr", exactOcr)
	return int(ret.Val)
}

func (com *DaMo) SetMinColGap(minColGap int) int {
	ret, _ := com.iDispatch.CallMethod("SetMinColGap", minColGap)
	return int(ret.Val)
}

func (com *DaMo) SetMinRowGap(minRowGap int) int {
	ret, _ := com.iDispatch.CallMethod("SetMinRowGap", minRowGap)
	return int(ret.Val)
}

func (com *DaMo) SetRowGapNoDict(RowGap int) int {
	ret, _ := com.iDispatch.CallMethod("SetRowGapNoDict", RowGap)
	return int(ret.Val)
}

func (com *DaMo) SetWordGap(wordGap int) int {
	ret, _ := com.iDispatch.CallMethod("SetWordGap", wordGap)
	return int(ret.Val)
}

func (com *DaMo) SetWordGapNoDict(wordGap int) int {
	ret, _ := com.iDispatch.CallMethod("SetWordGapNoDict", wordGap)
	return int(ret.Val)
}

func (com *DaMo) SetWordLineHeight(lineHeight int) int {
	ret, _ := com.iDispatch.CallMethod("SetWordLineHeight", lineHeight)
	return int(ret.Val)
}

func (com *DaMo) SetWordLineHeightNoDict(lineHeight int) int {
	ret, _ := com.iDispatch.CallMethod("SetWordLineHeightNoDict", lineHeight)
	return int(ret.Val)
}

func (com *DaMo) UseDict(index int) int {
	ret, _ := com.iDispatch.CallMethod("UseDict", index)
	return int(ret.Val)
}
