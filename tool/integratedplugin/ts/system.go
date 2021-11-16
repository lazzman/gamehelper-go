// 系统

package ts

func (com *TianShiPlug) CheckUAC() int {
	ret, _ := com.iDispatch.CallMethod("CheckUAC")
	return int(ret.Val)
}

func (com *TianShiPlug) CheckFontSmooth() int {
	ret, _ := com.iDispatch.CallMethod("CheckFontSmooth")
	return int(ret.Val)
}

func (com *TianShiPlug) Delay(mis int) int {
	ret, _ := com.iDispatch.CallMethod("NextDelay")
	return int(ret.Val)
}

func (com *TianShiPlug) DisableFontSmooth() int {
	ret, _ := com.iDispatch.CallMethod("DisableFontSmooth")
	return int(ret.Val)
}

func (com *TianShiPlug) EnableIme(enable int) int {
	return 1
}
