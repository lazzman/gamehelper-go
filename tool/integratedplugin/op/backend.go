// 后台

package op

func (com *Opsoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.iDispatch.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	return int(ret.Val)
}

func (com *Opsoft) UnBindWindow() int {
	ret, _ := com.iDispatch.CallMethod("UnBindWindow")
	return int(ret.Val)
}
