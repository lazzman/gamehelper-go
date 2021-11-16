// winapi

package op

func (com *Opsoft) RunApp(path string, mode int) int {
	ret, _ := com.iDispatch.CallMethod("RunApp", path, mode)
	return int(ret.Val)
}

func (com *Opsoft) WinExec(path string, dis int) int {
	ret, _ := com.iDispatch.CallMethod("WinExec", path, dis)
	return int(ret.Val)
}

func (com *Opsoft) GetCmdStr(path string, time int) string {
	ret, _ := com.iDispatch.CallMethod("GetCmdStr", path, time)
	return ret.ToString()
}
