// 窗口

package lw

func (com *LeWan) EnumWindow(parent int, title, className string, filter int) string {
	ret, _ := com.iDispatch.CallMethod("EnumWindow", parent, title, className)
	return ret.ToString()
}

func (com *LeWan) FindWindow(class, title string) int {
	ret, _ := com.iDispatch.CallMethod("FindWindow", title, class)
	return int(ret.Val)
}

func (com *LeWan) FindWindowEx(parent int, class, title string) int {
	ret, _ := com.iDispatch.CallMethod("FindWindow", title, class, "", 0, parent)
	return int(ret.Val)
}

func (com *LeWan) GetClientSize(hwnd int, width, height *int) int {
	ret, _ := com.iDispatch.CallMethod("GetClientSize", hwnd)
	*width = com.GetX()
	*height = com.GetY()
	return int(ret.Val)
}

func (com *LeWan) GetMousePointWindow() int {
	ret, _ := com.iDispatch.CallMethod("GetMousePointWindow")
	return int(ret.Val)
}

func (com *LeWan) GetPointWindow(x, y int) int {
	ret, _ := com.iDispatch.CallMethod("GetPointWindow", x, y)
	return int(ret.Val)
}

func (com *LeWan) GetWindow(hwnd, flag int) int {
	ret, _ := com.iDispatch.CallMethod("GetWindow", hwnd, flag)
	return int(ret.Val)
}

func (com *LeWan) GetWindowClass(hwnd int) string {
	ret, _ := com.iDispatch.CallMethod("GetWindowClass", hwnd)
	return ret.ToString()
}

func (com *LeWan) GetWindowProcessId(hwnd int) int {
	ret, _ := com.iDispatch.CallMethod("GetWindowProcessId", hwnd)
	return int(ret.Val)
}

func (com *LeWan) GetWindowProcessPath(hwnd int) string {
	ret, _ := com.iDispatch.CallMethod("GetWindowProcessPath", hwnd)
	return ret.ToString()
}

func (com *LeWan) GetWindowState(hwnd, flag int) int {
	ret, _ := com.iDispatch.CallMethod("GetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *LeWan) GetWindowTitle(hwnd int) string {
	ret, _ := com.iDispatch.CallMethod("GetWindowTitle", hwnd)
	return ret.ToString()
}

func (com *LeWan) MoveWindow(hwnd, x, y int) int {
	ret, _ := com.iDispatch.CallMethod("MoveWindow", hwnd, x, y)
	return int(ret.Val)
}

func (com *LeWan) ScreenToClient(hwnd int, x, y *int) int {
	ret, _ := com.iDispatch.CallMethod("ScreenToClient", hwnd)
	*x = com.GetX()
	*y = com.GetY()
	return int(ret.Val)
}

func (com *LeWan) SetWindowSize(hwnd, width, height int) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *LeWan) SetWindowState(hwnd, flag int) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *LeWan) SetWindowText(hwnd int, title string) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowText", hwnd, title)
	return int(ret.Val)
}

func (com *LeWan) SetWindowTransparent(hwnd, trans int) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowTransparent", hwnd, trans)
	return int(ret.Val)
}
