// 系统

package lw

func (com *LeWan) CheckUAC() int {
	ret, _ := com.iDispatch.CallMethod("CheckUAC")
	return int(ret.Val)
}

func (com *LeWan) CheckFontSmooth() int {
	ret, _ := com.iDispatch.CallMethod("CheckFontSmooth")
	return int(ret.Val)
}

func (com *LeWan) Delay(mis int) int {
	ret, _ := com.iDispatch.CallMethod("Delay")
	return int(ret.Val)
}

func (com *LeWan) DisableFontSmooth() int {
	ret, _ := com.iDispatch.CallMethod("DisableFontSmooth")
	return int(ret.Val)
}

func (com *LeWan) EnableIme(enable int) int {
	return 1
}
