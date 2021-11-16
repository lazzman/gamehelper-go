// 后台

package ts

func (com *TianShiPlug) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.iDispatch.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return int(ret.Val)
}

func (com *TianShiPlug) DownCPU(rate int) int {
	ret, _ := com.iDispatch.CallMethod("DownCpu", rate)
	return int(ret.Val)
}

func (com *TianShiPlug) EnableRealKeypad(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableRealKeypad", enable)
	return int(ret.Val)
}

func (com *TianShiPlug) EnableRealMouse(enable, mousedelay, mousestep int) int {
	ret, _ := com.iDispatch.CallMethod("EnableRealMouse", enable, mousedelay, mousestep)
	return int(ret.Val)
}

func (com *TianShiPlug) IsBind(hwnd int) int {
	ret, _ := com.iDispatch.CallMethod("IsBind", hwnd)
	return int(ret.Val)
}

func (com *TianShiPlug) LockInput(lock int) int {
	ret, _ := com.iDispatch.CallMethod("LockInput", lock)
	return int(ret.Val)
}

func (com *TianShiPlug) UnBindWindow() int {
	ret, _ := com.iDispatch.CallMethod("UnBindWindow")
	return int(ret.Val)
}
