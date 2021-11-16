// 键鼠

package dm

import "github.com/go-ole/go-ole"

func (com *DaMo) EnableMouseAccuracy(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnableMouseAccuracy", enable)
	return int(ret.Val)
}

func (com *DaMo) GetCursorPos(x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.iDispatch.CallMethod("GetCursorPos", &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *DaMo) GetCursorShape() string {
	ret, _ := com.iDispatch.CallMethod("GetCursorShape")
	return ret.ToString()
}

func (com *DaMo) GetCursorShapeEx(types int) string {
	ret, _ := com.iDispatch.CallMethod("GetCursorShapeEx", types)
	return ret.ToString()
}

func (com *DaMo) GetCursorSpot() string {
	ret, _ := com.iDispatch.CallMethod("GetCursorSpot")
	return ret.ToString()
}

func (com *DaMo) GetKeyState(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("GetKeyState", vkCode)
	return int(ret.Val)
}

func (com *DaMo) GetMouseSpeed() int {
	ret, _ := com.iDispatch.CallMethod("GetMouseSpeed")
	return int(ret.Val)
}

func (com *DaMo) KeyDown(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

func (com *DaMo) KeyDownChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyDownChar", keyStr)
	return int(ret.Val)
}

func (com *DaMo) KeyPress(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

func (com *DaMo) KeyPressChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyPressChar", keyStr)
	return int(ret.Val)
}

// KeyPressStr 根据指定的字符串序列，依次按顺序按下其中的字符
func (com *DaMo) KeyPressStr(keyStr string, delay int) int {
	ret, _ := com.iDispatch.CallMethod("KeyPressStr", keyStr, delay)
	return int(ret.Val)
}

func (com *DaMo) KeyUp(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

func (com *DaMo) KeyUpChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyUpChar", keyStr)
	return int(ret.Val)
}

// LeftClick 按下鼠标左键
func (com *DaMo) LeftClick() int {
	ret, _ := com.iDispatch.CallMethod("LeftClick")
	return int(ret.Val)
}

// LeftDoubleClick 双击鼠标左键
func (com *DaMo) LeftDoubleClick() int {
	ret, _ := com.iDispatch.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

// LeftDown 按住鼠标左键
func (com *DaMo) LeftDown() int {
	ret, _ := com.iDispatch.CallMethod("LeftDown")
	return int(ret.Val)
}

func (com *DaMo) LeftUp() int {
	ret, _ := com.iDispatch.CallMethod("LeftUp")
	return int(ret.Val)
}

func (com *DaMo) MiddleClick() int {
	ret, _ := com.iDispatch.CallMethod("MiddleClick")
	return int(ret.Val)
}

func (com *DaMo) MiddleDown() int {
	ret, _ := com.iDispatch.CallMethod("MiddleDown")
	return int(ret.Val)
}

func (com *DaMo) MiddleUp() int {
	ret, _ := com.iDispatch.CallMethod("MiddleUp")
	return int(ret.Val)
}

func (com *DaMo) MoveR(rx, ry int) int {
	ret, _ := com.iDispatch.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

func (com *DaMo) MoveTo(x, y int) int {
	ret, _ := com.iDispatch.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

func (com *DaMo) MoveToEx(x, y, w, h int) string {
	ret, _ := com.iDispatch.CallMethod("MoveToEx", x, y, w, h)
	return ret.ToString()
}

func (com *DaMo) RightClick() int {
	ret, _ := com.iDispatch.CallMethod("RightClick")
	return int(ret.Val)
}

func (com *DaMo) RightDown() int {
	ret, _ := com.iDispatch.CallMethod("RightDown")
	return int(ret.Val)
}

func (com *DaMo) RightUp() int {
	ret, _ := com.iDispatch.CallMethod("RightUp")
	return int(ret.Val)
}

func (com *DaMo) SetKeypadDelay(types string, delay int) int {
	ret, _ := com.iDispatch.CallMethod("SetKeypadDelay", types, delay)
	return int(ret.Val)
}

func (com *DaMo) SetMouseDelay(types string, delay int) int {
	ret, _ := com.iDispatch.CallMethod("SetMouseDelay", types, delay)
	return int(ret.Val)
}

func (com *DaMo) SetMouseSpeed(speed int) int {
	ret, _ := com.iDispatch.CallMethod("SetMouseSpeed", speed)
	return int(ret.Val)
}

func (com *DaMo) SetSimMode(mode int) int {
	ret, _ := com.iDispatch.CallMethod("SetSimMode", mode)
	return int(ret.Val)
}

func (com *DaMo) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.iDispatch.CallMethod("SetSimMode", vkCode, timeOut)
	return int(ret.Val)
}

func (com *DaMo) WheelDown() int {
	ret, _ := com.iDispatch.CallMethod("WheelDown")
	return int(ret.Val)
}

func (com *DaMo) WheelUp() int {
	ret, _ := com.iDispatch.CallMethod("WheelUp")
	return int(ret.Val)
}
