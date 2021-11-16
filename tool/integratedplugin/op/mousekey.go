// 键鼠

package op

import "github.com/go-ole/go-ole"

func (com *Opsoft) GetCursorPos(x, y *int) int {
	intx := ole.NewVariant(ole.VT_I4, int64(*x))
	inty := ole.NewVariant(ole.VT_I4, int64(*y))
	ret, _ := com.iDispatch.CallMethod("GetCursorPos", &intx, &inty)
	*x = int(intx.Val)
	*y = int(inty.Val)
	intx.Clear()
	inty.Clear()
	return int(ret.Val)
}

func (com *Opsoft) GetKeyState(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("GetKeyState", vkCode)
	return int(ret.Val)
}

func (com *Opsoft) KeyDown(vkCode string) int {
	ret, _ := com.iDispatch.CallMethod("KeyDown", vkCode)
	return int(ret.Val)
}

func (com *Opsoft) KeyDownChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyDownChar", keyStr)
	return int(ret.Val)
}

func (com *Opsoft) KeyPress(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyPress", vkCode)
	return int(ret.Val)
}

func (com *Opsoft) KeyPressChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyPressChar", keyStr)
	return int(ret.Val)
}

func (com *Opsoft) KeyUp(vkCode int) int {
	ret, _ := com.iDispatch.CallMethod("KeyUp", vkCode)
	return int(ret.Val)
}

func (com *Opsoft) KeyUpChar(keyStr string) int {
	ret, _ := com.iDispatch.CallMethod("KeyUpChar", keyStr)
	return int(ret.Val)
}

func (com *Opsoft) LeftClick() int {
	ret, _ := com.iDispatch.CallMethod("LeftClick")
	return int(ret.Val)
}

func (com *Opsoft) LeftDoubleClick() int {
	ret, _ := com.iDispatch.CallMethod("LeftDoubleClick")
	return int(ret.Val)
}

func (com *Opsoft) LeftDown() int {
	ret, _ := com.iDispatch.CallMethod("LeftDown")
	return int(ret.Val)
}

func (com *Opsoft) LeftUp() int {
	ret, _ := com.iDispatch.CallMethod("LeftUp")
	return int(ret.Val)
}

func (com *Opsoft) MiddleClick() int {
	ret, _ := com.iDispatch.CallMethod("MiddleClick")
	return int(ret.Val)
}

func (com *Opsoft) MiddleDown() int {
	ret, _ := com.iDispatch.CallMethod("MiddleDown")
	return int(ret.Val)
}

func (com *Opsoft) MiddleUp() int {
	ret, _ := com.iDispatch.CallMethod("MiddleUp")
	return int(ret.Val)
}

func (com *Opsoft) MoveR(rx, ry int) int {
	ret, _ := com.iDispatch.CallMethod("MoveR", rx, ry)
	return int(ret.Val)
}

func (com *Opsoft) MoveTo(x, y int) int {
	ret, _ := com.iDispatch.CallMethod("MoveTo", x, y)
	return int(ret.Val)
}

func (com *Opsoft) MoveToEx(x, y, w, h int) string {
	ret, _ := com.iDispatch.CallMethod("MoveToEx", x, y, w, h)
	return ret.ToString()
}

func (com *Opsoft) RightClick() int {
	ret, _ := com.iDispatch.CallMethod("RightClick")
	return int(ret.Val)
}

func (com *Opsoft) RightDown() int {
	ret, _ := com.iDispatch.CallMethod("RightDown")
	return int(ret.Val)
}

func (com *Opsoft) RightUp() int {
	ret, _ := com.iDispatch.CallMethod("RightUp")
	return int(ret.Val)
}

func (com *Opsoft) WaitKey(vkCode, timeOut int) int {
	ret, _ := com.iDispatch.CallMethod("SetSimMode", vkCode, timeOut)
	return int(ret.Val)
}

func (com *Opsoft) WheelDown() int {
	ret, _ := com.iDispatch.CallMethod("WheelDown")
	return int(ret.Val)
}

func (com *Opsoft) WheelUp() int {
	ret, _ := com.iDispatch.CallMethod("WheelUp")
	return int(ret.Val)
}
