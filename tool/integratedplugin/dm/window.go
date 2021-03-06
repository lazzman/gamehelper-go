// 窗口

package dm

import "github.com/go-ole/go-ole"

func (com *DaMo) ClientToScreen(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.iDispatch.CallMethod("ClientToScreen", hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *DaMo) EnumProcess(name string) string {
	ret, _ := com.iDispatch.CallMethod("EnumProcess", name)
	return ret.ToString()
}

func (com *DaMo) EnumWindow(parent int, title, className string, filter int) string {
	ret, _ := com.iDispatch.CallMethod("EnumWindow", parent, title, className, filter)
	return ret.ToString()
}

func (com *DaMo) EnumWindowByProcess(processName, title, className string, filter int) string {
	ret, _ := com.iDispatch.CallMethod("EnumWindowByProcess", processName, title, className, filter)
	return ret.ToString()
}

func (com *DaMo) EnumWindowByProcessId(pid int, title, className string, filter int) string {
	ret, _ := com.iDispatch.CallMethod("EnumWindowByProcessId", pid, title, className, filter)
	return ret.ToString()
}

func (com *DaMo) EnumWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2, sort int) string {
	ret, _ := com.iDispatch.CallMethod("EnumWindowSuper", spec1, flag1, type1, spec2, flag2, type2, sort)
	return ret.ToString()
}

func (com *DaMo) FindWindow(class, title string) int {
	ret, _ := com.iDispatch.CallMethod("FindWindow", class, title)
	return int(ret.Val)
}

func (com *DaMo) FindWindowByProcess(processName, class, title string) int {
	ret, _ := com.iDispatch.CallMethod("FindWindowByProcess", processName, class, title)
	return int(ret.Val)
}

func (com *DaMo) FindWindowByProcessId(processId int, class, title string) int {
	ret, _ := com.iDispatch.CallMethod("FindWindowByProcessId", processId, class, title)
	return int(ret.Val)
}

func (com *DaMo) FindWindowEx(parent int, class, title string) int {
	ret, _ := com.iDispatch.CallMethod("FindWindowEx", parent, class, title)
	return int(ret.Val)
}

func (com *DaMo) FindWindowSuper(spec1 string, flag1, type1 int, spec2 string, flag2, type2 int) int {
	ret, _ := com.iDispatch.CallMethod("FindWindowSuper", spec1, flag1, type1, spec2, flag2, type2)
	return int(ret.Val)
}

func (com *DaMo) GetClientRect(hwnd int, x1, y1, x2, y2 *int) int {
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

func (com *DaMo) GetClientSize(hwnd int, width, height *int) int {
	pWidth := ole.NewVariant(ole.VT_I4, int64(*width))
	pHeight := ole.NewVariant(ole.VT_I4, int64(*height))
	ret, _ := com.iDispatch.CallMethod("GetClientSize", hwnd, &pWidth, &pHeight)
	*width = int(pWidth.Val)
	*height = int(pHeight.Val)
	pWidth.Clear()
	pHeight.Clear()
	return int(ret.Val)
}

func (com *DaMo) GetForegroundFocus() int {
	ret, _ := com.iDispatch.CallMethod("GetForegroundFocus")
	return int(ret.Val)
}

func (com *DaMo) GetForegroundWindow() int {
	ret, _ := com.iDispatch.CallMethod("GetForegroundWindow")
	return int(ret.Val)
}

func (com *DaMo) GetMousePointWindow() int {
	ret, _ := com.iDispatch.CallMethod("GetMousePointWindow")
	return int(ret.Val)
}

func (com *DaMo) GetPointWindow(x, y int) int {
	ret, _ := com.iDispatch.CallMethod("GetPointWindow", x, y)
	return int(ret.Val)
}

func (com *DaMo) GetProcessInfo(pid int) string {
	ret, _ := com.iDispatch.CallMethod("GetProcessInfo", pid)
	return ret.ToString()
}

func (com *DaMo) GetSpecialWindow(flag int) int {
	ret, _ := com.iDispatch.CallMethod("GetSpecialWindow", flag)
	return int(ret.Val)
}

func (com *DaMo) GetWindow(hwnd, flag int) int {
	ret, _ := com.iDispatch.CallMethod("GetWindow", hwnd, flag)
	return int(ret.Val)
}

func (com *DaMo) GetWindowClass(hwnd int) string {
	ret, _ := com.iDispatch.CallMethod("GetWindowClass", hwnd)
	return ret.ToString()
}

func (com *DaMo) GetWindowProcessId(hwnd int) int {
	ret, _ := com.iDispatch.CallMethod("GetWindowProcessId", hwnd)
	return int(ret.Val)
}

func (com *DaMo) GetWindowProcessPath(hwnd int) string {
	ret, _ := com.iDispatch.CallMethod("GetWindowProcessPath", hwnd)
	return ret.ToString()
}

func (com *DaMo) GetWindowRect(hwnd int, x1, y1, x2, y2 *int) int {
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

func (com *DaMo) GetWindowState(hwnd, flag int) int {
	ret, _ := com.iDispatch.CallMethod("GetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *DaMo) GetWindowTitle(hwnd int) string {
	ret, _ := com.iDispatch.CallMethod("GetWindowTitle", hwnd)
	return ret.ToString()
}

func (com *DaMo) MoveWindow(hwnd, x, y int) int {
	ret, _ := com.iDispatch.CallMethod("MoveWindow", hwnd, x, y)
	return int(ret.Val)
}

func (com *DaMo) ScreenToClient(hwnd int, x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.iDispatch.CallMethod("ScreenToClient", hwnd, &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *DaMo) SendPaste(hwnd int) int {
	ret, _ := com.iDispatch.CallMethod("SendPaste", hwnd)
	return int(ret.Val)
}

func (com *DaMo) SendString(hwnd int, str string) int {
	ret, _ := com.iDispatch.CallMethod("SendString", hwnd, str)
	return int(ret.Val)
}

func (com *DaMo) SendString2(hwnd int, str string) int {
	ret, _ := com.iDispatch.CallMethod("SendString2", hwnd, str)
	return int(ret.Val)
}

func (com *DaMo) SendStringIme(str string) int {
	ret, _ := com.iDispatch.CallMethod("SendStringIme", str)
	return int(ret.Val)
}

func (com *DaMo) SendStringIme2(hwnd int, str string, mode int) int {
	ret, _ := com.iDispatch.CallMethod("SendStringIme2", hwnd, str, mode)
	return int(ret.Val)
}

func (com *DaMo) SetClientSize(hwnd, width, height int) int {
	ret, _ := com.iDispatch.CallMethod("SetClientSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *DaMo) SetWindowSize(hwnd, width, height int) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowSize", hwnd, width, height)
	return int(ret.Val)
}

func (com *DaMo) SetWindowState(hwnd, flag int) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowState", hwnd, flag)
	return int(ret.Val)
}

func (com *DaMo) SetWindowText(hwnd int, title string) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowText", hwnd, title)
	return int(ret.Val)
}

func (com *DaMo) SetWindowTransparent(hwnd, trans int) int {
	ret, _ := com.iDispatch.CallMethod("SetWindowTransparent", hwnd, trans)
	return int(ret.Val)
}
