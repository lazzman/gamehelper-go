// 基本

package ts

func (com *TianShiPlug) Version() string {
	ver, _ := com.iDispatch.CallMethod("Ver")
	return ver.ToString()
}

func (com *TianShiPlug) EnablePicCache(enable int) int {
	return 1
}

func (com *TianShiPlug) GetBasePath() string {
	ret, _ := com.iDispatch.CallMethod("GetBasePath")
	return ret.ToString()
}

func (com *TianShiPlug) GetPath() string {
	ret, _ := com.iDispatch.CallMethod("GetPath")
	return ret.ToString()
}

func (com *TianShiPlug) SetPath(path string) int {
	ret, _ := com.iDispatch.CallMethod("SetPath", path)
	return int(ret.Val)
}

func (com *TianShiPlug) SetShowErrorMsg(show int) int {
	ret, _ := com.iDispatch.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

func (com *TianShiPlug) Reg(regCode string, verInfo string) int {
	ret, _ := com.iDispatch.CallMethod("Reg", regCode, verInfo)
	return int(ret.Val)
}
