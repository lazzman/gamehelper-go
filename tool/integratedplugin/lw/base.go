// 基本

package lw

func (com *LeWan) Version() string {
	ver, _ := com.iDispatch.CallMethod("ver")
	return ver.ToString()
}

func (com *LeWan) EnablePicCache(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnablePicCache", enable)
	return int(ret.Val)
}

func (com *LeWan) GetPath() string {
	ret, _ := com.iDispatch.CallMethod("GetPath")
	return ret.ToString()
}

func (com *LeWan) SetPath(path string) int {
	ret, _ := com.iDispatch.CallMethod("SetPath", path)
	return int(ret.Val)
}

func (com *LeWan) SetShowErrorMsg(show int) int {
	ret, _ := com.iDispatch.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

func (com *LeWan) GetX() int {
	ret, _ := com.iDispatch.CallMethod("GetX")
	return int(ret.Val)
}

func (com *LeWan) GetY() int {
	ret, _ := com.iDispatch.CallMethod("GetY")
	return int(ret.Val)
}

func (com *LeWan) GetIdx() int {
	ret, _ := com.iDispatch.CallMethod("GetIdx")
	return int(ret.Val)
}
