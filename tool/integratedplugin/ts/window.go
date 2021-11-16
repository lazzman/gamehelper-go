// 窗口

package ts

import "github.com/go-ole/go-ole"

func (com *TianShiPlug) EnumProcess(name string) string {
	ret, _ := com.iDispatch.CallMethod("EnumProcess", name)
	return ret.ToString()
}

func (com *TianShiPlug) EnumWindow(parent int, title, className string, filter int) string {
	ret, _ := com.iDispatch.CallMethod("EnumWindow", parent, title, className, filter)
	return ret.ToString()
}

func (com *TianShiPlug) EnumWindowByProcess(processName, title, className string, filter int) string {
	ret, _ := com.iDispatch.CallMethod("EnumWindowByProcess", processName, title, className, filter)
	return ret.ToString()
}

func (com *TianShiPlug) FindWindow(class, title string) int {
	ret, _ := com.iDispatch.CallMethod("FindWindow", class, title)
	return int(ret.Val)
}

func (com *TianShiPlug) FindWindowByProcess(processName, class, title string) int {
	ret, _ := com.iDispatch.CallMethod("FindWindowByProcess", processName, class, title)
	return int(ret.Val)
}

func (com *TianShiPlug) FindWindowEx(parent int, class, title string) int {
	ret, _ := com.iDispatch.CallMethod("FindWindowEx", parent, class, title)
	return int(ret.Val)
}

func (com *TianShiPlug) GetClientRect(hwnd int, x1, y1, x2, y2 *int) int {
	intx1 := ole.NewVariant(ole.VT_I4, int64(*x1))
	inty1 := ole.NewVariant(ole.VT_I4, int64(*y1))
	intx2 := ole.NewVariant(ole.VT_I4, int64(*x2))
	inty2 := ole.NewVariant(ole.VT_I4, int64(*y2))
	ret, _ := com.iDispatch.CallMethod("GetClientRect", hwnd, &intx1, &inty1, &intx2, &inty2)
	*x1 = int(intx1.Val)
	*y1 = int(inty1.Val)
	*x2 = int(intx2.Val)
	*y2 = int(inty2.Val)
	// 清除对象变量内存
	intx1.Clear()
	inty1.Clear()
	intx2.Clear()
	inty2.Clear()
	return int(ret.Val)
}

func (com *TianShiPlug) GetClientSize(hwnd int, width, height *int) int {
	pWidth := ole.NewVariant(ole.VT_I4, int64(*width))
	pHeight := ole.NewVariant(ole.VT_I4, int64(*height))
	ret, _ := com.iDispatch.CallMethod("GetClientSize", hwnd, &pWidth, &pHeight)
	*width = int(pWidth.Val)
	*height = int(pHeight.Val)
	pWidth.Clear()
	pHeight.Clear()
	return int(ret.Val)
}

func (com *TianShiPlug) GetForegroundFocus() int {
	ret, _ := com.iDispatch.CallMethod("GetForegroundFocus")
	return int(ret.Val)
}

func (com *TianShiPlug) GetForegroundWindow() int {
	ret, _ := com.iDispatch.CallMethod("GetForegroundWindow")
	return int(ret.Val)
}

func (com *TianShiPlug) GetMousePointWindow() int {
	ret, _ := com.iDispatch.CallMethod("GetMousePointWindow")
	return int(ret.Val)
}

func (com *TianShiPlug) GetPointWindow(x, y int) int {
	ret, _ := com.iDispatch.CallMethod("GetPointWindow", x, y)
	return int(ret.Val)
}

func (com *TianShiPlug) GetProcessInfo(pid int) string {
	ret, _ := com.iDispatch.CallMethod("GetProcessInfo", pid)
	return ret.ToString()
}

func (com *TianShiPlug) GetSpecialWindow(flag int) int {
	ret, _ := com.iDispatch.CallMethod("GetSpecialWindow", flag)
	return int(ret.Val)
}

func (com *TianShiPlug) GetWindow(hwnd, flag int) int {
	ret, _ := com.iDispatch.CallMethod("GetWindow", hwnd, flag)
	return int(ret.Val)
}

func (com *TianShiPlug) GetWindowClass(hwnd int) string {
	ret, _ := com.iDispatch.CallMethod("GetWindowClass", hwnd)
	return ret.ToString()
}

func (com *TianShiPlug) GetWindowProcessId(hwnd int) int {
	ret, _ := com.iDispatch.CallMethod("GetWindowProcessId", hwnd)
	return int(ret.Val)
}

func (com *TianShiPlug) GetWindowProcessPath(hwnd int) string {
	ret, _ := com.iDispatch.CallMethod("GetWindowProcessPath", hwnd)
	return ret.ToString()
}

func (com *TianShiPlug) GetWindowRect(hwnd int, x1, y1, x2, y2 *int) int {
	intx1 := ole.NewVariant(ole.VT_I4, int64(*x1))
	inty1 := ole.NewVariant(ole.VT_I4, int64(*y1))
	intx2 := ole.NewVariant(ole.VT_I4, int64(*x2))
	inty2 := ole.NewVariant(ole.VT_I4, int64(*y2))
	ret, _ := com.iDispatch.CallMethod("GetWindowRect", hwnd, &intx1, &inty1, &intx2, &inty2)
	*x1 = int(intx1.Val)
	*y1 = int(inty1.Val)
	*x2 = int(intx2.Val)
	*y2 = int(inty2.Val)
	intx1.Clear()
	inty1.Clear()
	intx2.Clear()
	inty2.Clear()
	return int(ret.Val)
}

func (com *TianShiPlug) GetWindowState(hwnd, flag int) int {
	ret, _ := com.iDispatch.CallMethod("GetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *TianShiPlug) GetWindowTitle(hwnd int) string {
	ret, _ := com.iDispatch.CallMethod("GetWindowTitle", hwnd)
	return ret.ToString()
}

func (com *TianShiPlug) MoveWindow(hwnd, x, y int) int {
	ret, _ := com.iDispatch.CallMethod("MoveWindow", hwnd, x, y)
	return int(ret.Val)
}

func (com *TianShiPlug) ScreenToClient(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.iDispatch.CallMethod("ScreenToClient", hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *TianShiPlug) SendPaste(hwnd int) int {
	ret, _ := com.iDispatch.CallMethod("SendPaste", hwnd)
	return int(ret.Val)
}

func (com *TianShiPlug) SetClientSize(hwnd, width, height int) int {
	ret, _ := com.iDispatch.CallMethod("SetClientSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *TianShiPlug) SetWindowSize(hwnd, width, height int) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *TianShiPlug) SetWindowState(hwnd, flag int) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *TianShiPlug) SetWindowText(hwnd int, title string) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowText", hwnd, title)
	return int(ret.Val)
}

func (com *TianShiPlug) SetWindowTransparent(hwnd, trans int) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowTransparent", hwnd, trans)
	return int(ret.Val)
}
