// 后台

package lw

func (com *LeWan) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	displayParam := getLWDisplayParam(display)
	mouseParam := getLWMouseParam(mouse)
	keypadParam := getLWKeypadParam(keypad)
	ret, _ := com.iDispatch.CallMethod("BindWindow", hwnd, displayParam, mouseParam, keypadParam, 0, mode)
	return int(ret.Val)
}

func getLWDisplayParam(display string) int {
	switch display {
	case "normal":
		return 0
	case "opengl":
		return 5
	case "gdi":
		return 1
	case "gdi2":
		return 2
	case "dx2":
		return 3
	case "dx":
		return 4
	default:
		return 0
	}
}

func getLWMouseParam(mouse string) int {
	switch mouse {
	case "normal":
		return 0
	case "windows":
		return 1
	case "dx":
		return 3
	default:
		return 0
	}
}

func getLWKeypadParam(keypad string) int {
	switch keypad {
	case "normal":
		return 0
	case "windows":
		return 1
	case "dx":
		return 3
	default:
		return 0
	}
}

func (com *LeWan) DownCPU(rate int) int {
	ret, _ := com.iDispatch.CallMethod("DownCpu", rate)
	return int(ret.Val)
}

func (com *LeWan) EnableRealKeypad(enable int) int {
	return 1
}

func (com *LeWan) EnableRealMouse(enable, mousedelay, mousestep int) int {
	ret, _ := com.iDispatch.CallMethod("EnableRealMouse", enable, mousedelay, mousestep)
	return int(ret.Val)
}

func (com *LeWan) IsBind(hwnd int) int {
	ret, _ := com.iDispatch.CallMethod("IsBind", hwnd)
	return int(ret.Val)
}

func (com *LeWan) LockInput(lock int) int {
	ret, _ := com.iDispatch.CallMethod("LockInput", lock)
	return int(ret.Val)
}

func (com *LeWan) UnBindWindow() int {
	ret, _ := com.iDispatch.CallMethod("UnBindWindow")
	return int(ret.Val)
}
