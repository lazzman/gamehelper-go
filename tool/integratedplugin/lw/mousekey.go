// 键鼠

package lw

func (com *LeWan) SetMouseSpeed(speed int) int {
	return 1
}

func (com *LeWan) EnableMouseAccuracy(enable int) int {
	return 1
}

func (com *LeWan) KeyDown(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

func (com *LeWan) KeyPress(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

func (com *LeWan) KeyDownChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyDown", VKCODEMAP[keyStr])
	return int(ret.Val)
}

func (com *LeWan) KeyUpChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyUp", VKCODEMAP[keyStr])
	return int(ret.Val)
}

func (com *LeWan) KeyPressChar(keyStr string) int {
	com.KeyDownChar(keyStr)
	com.KeyUpChar(keyStr)
	return 1
}

// KeyPressStr 根据指定的字符串序列，依次按顺序按下其中的字符
func (com *LeWan) KeyPressStr(keyStr string, delay int) int {
	ret, _ := com.iDispatch.CallMethod("KeyPressStr", keyStr, delay)
	return int(ret.Val)
}

func (com *LeWan) KeyUp(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

// LeftClick 按下鼠标左键
func (com *LeWan) LeftClick() int {
	ret, _ := com.iDispatch.CallMethod("LeftClick")
	return int(ret.Val)
}

// LeftDoubleClick 双击鼠标左键
func (com *LeWan) LeftDoubleClick() int {
	ret, _ := com.iDispatch.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

// LeftDown 按住鼠标左键
func (com *LeWan) LeftDown() int {
	ret, _ := com.iDispatch.CallMethod("LeftDown")
	return int(ret.Val)
}

func (com *LeWan) LeftUp() int {
	ret, _ := com.iDispatch.CallMethod("LeftUp")
	return int(ret.Val)
}

func (com *LeWan) MiddleClick() int {
	ret, _ := com.iDispatch.CallMethod("MiddleClick")
	return int(ret.Val)
}

func (com *LeWan) MoveR(rx, ry int) int {
	ret, _ := com.iDispatch.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

func (com *LeWan) MoveTo(x, y int) int {
	ret, _ := com.iDispatch.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

func (com *LeWan) RightClick() int {
	ret, _ := com.iDispatch.CallMethod("RightClick")
	return int(ret.Val)
}

func (com *LeWan) RightDown() int {
	ret, _ := com.iDispatch.CallMethod("RightDown")
	return int(ret.Val)
}

func (com *LeWan) RightUp() int {
	ret, _ := com.iDispatch.CallMethod("RightUp")
	return int(ret.Val)
}

func (com *LeWan) SetKeypadDelay(types string, delay int) int {
	ret, _ := com.iDispatch.CallMethod("SetKeypadDelay", types, delay)
	return int(ret.Val)
}

func (com *LeWan) SetMouseDelay(types string, delay int) int {
	ret, _ := com.iDispatch.CallMethod("SetMouseDelay", types, delay)
	return int(ret.Val)
}

func (com *LeWan) SetSimMode(mode int) int {
	ret, _ := com.iDispatch.CallMethod("SetSimMode", mode)
	return int(ret.Val)
}

func (com *LeWan) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.iDispatch.CallMethod("WaitKey", vkCode, timeOut)
	return int(ret.Val)
}

func (com *LeWan) WheelDown() int {
	ret, _ := com.iDispatch.CallMethod("WheelDown")
	return int(ret.Val)
}

func (com *LeWan) WheelUp() int {
	ret, _ := com.iDispatch.CallMethod("WheelUp")
	return int(ret.Val)
}
