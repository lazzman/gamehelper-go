// 基本

package dm

func (com *DaMo) Version() string {
	ver, _ := com.iDispatch.CallMethod("Ver")
	return ver.ToString()
}

func (com *DaMo) EnablePicCache(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnablePicCache", enable)
	return int(ret.Val)
}

func (com *DaMo) GetBasePath() string {
	ret, _ := com.iDispatch.CallMethod("GetBasePath")
	return ret.ToString()
}

func (com *DaMo) GetDmCount() int {
	ret, _ := com.iDispatch.CallMethod("GetDmCount")
	return int(ret.Val)
}

func (com *DaMo) GetID() int {
	ret, _ := com.iDispatch.CallMethod("GetID")
	return int(ret.Val)
}

func (com *DaMo) GetLastError() int {
	ret, _ := com.iDispatch.CallMethod("GetLastError")
	return int(ret.Val)
}

func (com *DaMo) GetPath() string {
	ret, _ := com.iDispatch.CallMethod("GetPath")
	return ret.ToString()
}

func (com *DaMo) Reg(regCode string, verInfo string) int {
	ret, _ := com.iDispatch.CallMethod("Reg", regCode, verInfo)
	return int(ret.Val)
}

func (com *DaMo) RegEx(regCode, verInfo, ip string) int {
	ret, _ := com.iDispatch.CallMethod("RegEx", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *DaMo) RegExNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.iDispatch.CallMethod("RegExNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *DaMo) RegNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.iDispatch.CallMethod("RegNoMac", regCode, verInfo, ip)
	return int(ret.Val)
}

func (com *DaMo) SetDisplayInput(mode string) int {
	ret, _ := com.iDispatch.CallMethod("SetDisplayInput", mode)
	return int(ret.Val)
}

func (com *DaMo) SetEnumWindowDelay(delay int) int {
	ret, _ := com.iDispatch.CallMethod("SetEnumWindowDelay", delay)
	return int(ret.Val)
}

func (com *DaMo) SetPath(path string) int {
	ret, _ := com.iDispatch.CallMethod("SetPath", path)
	return int(ret.Val)
}

func (com *DaMo) SetShowErrorMsg(show int) int {
	ret, _ := com.iDispatch.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

func (com *DaMo) SpeedNormalGraphic(enable int) int {
	ret, _ := com.iDispatch.CallMethod("SpeedNormalGraphic", enable)
	return int(ret.Val)
}
