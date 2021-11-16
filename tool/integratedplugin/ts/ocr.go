// 文字识别

package ts

import "github.com/go-ole/go-ole"

// ClearDict 清空指定的字库
func (com *TianShiPlug) ClearDict(index int) int {
	ret, _ := com.iDispatch.CallMethod("ClearDict", index)
	return int(ret.Val)
}

func (com *TianShiPlug) FindStr(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindStr", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *TianShiPlug) FindStrEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *TianShiPlug) FindStrExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrExS", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *TianShiPlug) FindStrFast(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindStrFast", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *TianShiPlug) FindStrFastEx(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrFastEx", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *TianShiPlug) FindStrFastExS(x1, y1, x2, y2 int, str, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("FindStrFastExS", x1, y1, x2, y2, str, colorFormat, sim)
	return ret.ToString()
}

func (com *TianShiPlug) FindStrFastS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindStrFastS", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *TianShiPlug) FindStrS(x1, y1, x2, y2 int, str string, colorFormat string, sim float32, intX *int, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.iDispatch.CallMethod("FindStrS", x1, y1, x2, y2, str, colorFormat, sim, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *TianShiPlug) GetNowDict() int {
	ret, _ := com.iDispatch.CallMethod("GetNowDict")
	return int(ret.Val)
}

func (com *TianShiPlug) Ocr(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("Ocr", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *TianShiPlug) OcrEx(x1, y1, x2, y2 int, colorFormat string, sim float32) string {
	ret, _ := com.iDispatch.CallMethod("OcrEx", x1, y1, x2, y2, colorFormat, sim)
	return ret.ToString()
}

func (com *TianShiPlug) SetDict(index int, file string) int {
	ret, _ := com.iDispatch.CallMethod("SetDict", index, file)
	return int(ret.Val)
}

func (com *TianShiPlug) SetDictPwd(pwd string) int {
	ret, _ := com.iDispatch.CallMethod("SetDictPwd", pwd)
	return int(ret.Val)
}
