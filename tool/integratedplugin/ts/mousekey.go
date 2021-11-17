// 键鼠

package ts

import "github.com/go-ole/go-ole"

func (com *TianShiPlug) SetMouseSpeed(speed int) int {
	return 1
}

func (com *TianShiPlug) EnableMouseAccuracy(enable int) int {
	return 1
}

func (com *TianShiPlug) GetCursorPos(x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.iDispatch.CallMethod("GetCursorPos", &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *TianShiPlug) GetCursorShape() string {
	ret, _ := com.iDispatch.CallMethod("GetCursorShape")
	return ret.ToString()
}

func (com *TianShiPlug) KeyDown(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

func (com *TianShiPlug) KeyDownChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyDownChar", keyStr)
	return int(ret.Val)
}

func (com *TianShiPlug) KeyPress(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

func (com *TianShiPlug) KeyPressChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyPressChar", keyStr)
	return int(ret.Val)
}

// KeyPressStr 根据指定的字符串序列，依次按顺序按下其中的字符
func (com *TianShiPlug) KeyPressStr(keyStr string, delay int) int {
	ret, _ := com.iDispatch.CallMethod("KeyPressStr", keyStr, delay)
	return int(ret.Val)
}

func (com *TianShiPlug) KeyUp(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

func (com *TianShiPlug) KeyUpChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyUpChar", keyStr)
	return int(ret.Val)
}

// LeftClick 按下鼠标左键
func (com *TianShiPlug) LeftClick() int {
	ret, _ := com.iDispatch.CallMethod("LeftClick")
	return int(ret.Val)
}

// LeftDoubleClick 双击鼠标左键
func (com *TianShiPlug) LeftDoubleClick() int {
	ret, _ := com.iDispatch.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

// LeftDown 按住鼠标左键
func (com *TianShiPlug) LeftDown() int {
	ret, _ := com.iDispatch.CallMethod("LeftDown")
	return int(ret.Val)
}

func (com *TianShiPlug) LeftUp() int {
	ret, _ := com.iDispatch.CallMethod("LeftUp")
	return int(ret.Val)
}

func (com *TianShiPlug) MiddleClick() int {
	ret, _ := com.iDispatch.CallMethod("MiddleClick")
	return int(ret.Val)
}

func (com *TianShiPlug) MoveR(rx, ry int) int {
	ret, _ := com.iDispatch.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

func (com *TianShiPlug) MoveTo(x, y int) int {
	ret, _ := com.iDispatch.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

func (com *TianShiPlug) RightClick() int {
	ret, _ := com.iDispatch.CallMethod("RightClick")
	return int(ret.Val)
}

func (com *TianShiPlug) RightDown() int {
	ret, _ := com.iDispatch.CallMethod("RightDown")
	return int(ret.Val)
}

func (com *TianShiPlug) RightUp() int {
	ret, _ := com.iDispatch.CallMethod("RightUp")
	return int(ret.Val)
}

func (com *TianShiPlug) SetKeypadDelay(types string, delay int) int {
	ret, _ := com.iDispatch.CallMethod("SetKeypadDelay", types, delay)
	return int(ret.Val)
}

func (com *TianShiPlug) SetMouseDelay(types string, delay int) int {
	ret, _ := com.iDispatch.CallMethod("SetMouseDelay", types, delay)
	return int(ret.Val)
}

func (com *TianShiPlug) SetSimMode(mode int) int {
	ret, _ := com.iDispatch.CallMethod("SetSimMode", mode)
	return int(ret.Val)
}

func (com *TianShiPlug) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.iDispatch.CallMethod("WaitKey", vkCode, timeOut)
	return int(ret.Val)
}

func (com *TianShiPlug) WheelDown() int {
	ret, _ := com.iDispatch.CallMethod("WheelDown")
	return int(ret.Val)
}

func (com *TianShiPlug) WheelUp() int {
	ret, _ := com.iDispatch.CallMethod("WheelUp")
	return int(ret.Val)
}
