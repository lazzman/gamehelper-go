// 基本

package op

func (com *Opsoft) GetBasePath() string {
	ret, _ := com.iDispatch.CallMethod("GetBasePath")
	return ret.ToString()
}

func (com *Opsoft) GetID() int {
	ret, _ := com.iDispatch.CallMethod("GetID")
	return int(ret.Val)
}

func (com *Opsoft) GetLastError() int {
	ret, _ := com.iDispatch.CallMethod("GetLastError")
	return int(ret.Val)
}

func (com *Opsoft) GetPath() string {
	ret, _ := com.iDispatch.CallMethod("GetPath")
	return ret.ToString()
}

func (com *Opsoft) SetPath(path string) int {
	ret, _ := com.iDispatch.CallMethod("SetPath", path)
	return int(ret.Val)
}

func (com *Opsoft) SetShowErrorMsg(show int) int {
	ret, _ := com.iDispatch.CallMethod("SetShowErrorMsg", show)
	return int(ret.Val)
}

func (com *Opsoft) Version() string {
	ver, _ := com.iDispatch.CallMethod("Ver")
	return ver.ToString()
}

func (com *Opsoft) EnablePicCache(enable int) int {
	ret, _ := com.iDispatch.CallMethod("EnablePicCache", enable)
	return int(ret.Val)
}
