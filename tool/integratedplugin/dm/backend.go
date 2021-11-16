// 后台

package dm

func (com *DaMo) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.iDispatch.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return int(ret.Val)
}

func (com *DaMo) BindWindowEx(hwnd int, display string, mouse string, keypad string, public string, mode int) int {
	ret, _ := com.iDispatch.CallMethod("BindWindowEx", hwnd, display, mouse, keypad, public, mode)
	return int(ret.Val)
}

func (com *DaMo) DownCPU(rate int) int {
	ret, _ := com.iDispatch.CallMethod("DownCpu", rate)
	return int(ret.Val)
}

func (com *DaMo) EnableBind(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableBind", enable)
	return int(ret.Val)
}

func (com *DaMo) EnableFakeActive(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableFakeActive", enable)
	return int(ret.Val)
}

func (com *DaMo) EnableIme(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableIme", enable)
	return int(ret.Val)
}

func (com *DaMo) EnableKeypadMsg(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableKeypadMsg", enable)
	return int(ret.Val)
}

func (com *DaMo) EnableKeypadPatch(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableKeypadPatch", enable)
	return int(ret.Val)
}

func (com *DaMo) EnableKeypadSync(enable, timeOut int) int {
	ret, _ := com.iDispatch.CallMethod("EnableKeypadSync", enable, timeOut)
	return int(ret.Val)
}

func (com *DaMo) EnableMouseMsg(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableMouseMsg", enable)
	return int(ret.Val)
}

func (com *DaMo) EnableMouseSync(enable, timeOut int) int {
	ret, _ := com.iDispatch.CallMethod("EnableMouseSync", enable, timeOut)
	return int(ret.Val)
}

func (com *DaMo) EnableRealKeypad(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableRealKeypad", enable)
	return int(ret.Val)
}

func (com *DaMo) EnableRealMouse(enable, mousedelay, mousestep int) int {
	ret, _ := com.iDispatch.CallMethod("EnableRealMouse", enable, mousedelay, mousestep)
	return int(ret.Val)
}

func (com *DaMo) EnableSpeedDx(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableSpeedDx", enable)
	return int(ret.Val)
}

func (com *DaMo) ForceUnBindWindow() int {
	ret, _ := com.iDispatch.CallMethod("ForceUnBindWindow")
	return int(ret.Val)
}

func (com *DaMo) GetBindWindow() int {
	ret, _ := com.iDispatch.CallMethod("GetBindWindow")
	return int(ret.Val)
}

func (com *DaMo) GetFps() int {
	ret, _ := com.iDispatch.CallMethod("GetFps")
	return int(ret.Val)
}

func (com *DaMo) HackSpeed(rate int) int {
	ret, _ := com.iDispatch.CallMethod("HackSpeed", rate)
	return int(ret.Val)
}

func (com *DaMo) IsBind(hwnd int) int {
	ret, _ := com.iDispatch.CallMethod("IsBind", hwnd)
	return int(ret.Val)
}

func (com *DaMo) LockDisplay(lock int) int {
	ret, _ := com.iDispatch.CallMethod("LockDisplay", lock)
	return int(ret.Val)
}

func (com *DaMo) LockInput(lock int) int {
	ret, _ := com.iDispatch.CallMethod("LockInput", lock)
	return int(ret.Val)
}

func (com *DaMo) LockMouseRect(x1, y1, x2, y2 int) int {
	ret, _ := com.iDispatch.CallMethod("LockMouseRect", x1, y1, x2, y2)
	return int(ret.Val)
}

func (com *DaMo) SetAero(enable int) int {
	ret, _ := com.iDispatch.CallMethod("SetAero", enable)
	return int(ret.Val)
}

func (com *DaMo) SetDisplayDelay(time int) int {
	ret, _ := com.iDispatch.CallMethod("SetDisplayDelay", time)
	return int(ret.Val)
}

func (com *DaMo) SetDisplayRefreshDelay(time int) int {
	ret, _ := com.iDispatch.CallMethod("SetDisplayRefreshDelay", time)
	return int(ret.Val)
}

func (com *DaMo) SwitchBindWindow(hwnd int) int {
	ret, _ := com.iDispatch.CallMethod("SwitchBindWindow", hwnd)
	return int(ret.Val)
}

func (com *DaMo) UnBindWindow() int {
	ret, _ := com.iDispatch.CallMethod("UnBindWindow")
	return int(ret.Val)
}
