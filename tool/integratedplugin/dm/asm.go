// 汇编

package dm

func (com *DaMo) AsmAdd(asmIns string) int {
	ret, _ := com.iDispatch.CallMethod("AsmAdd", asmIns)
	return int(ret.Val)
}

func (com *DaMo) AsmCall(hwnd, mode int) int {
	ret, _ := com.iDispatch.CallMethod("AsmCall", hwnd, mode)
	return int(ret.Val)
}

func (com *DaMo) AsmCallEx(hwnd, mode int, baseAddr string) int64 {
	ret, _ := com.iDispatch.CallMethod("AsmCallEx", hwnd, mode, baseAddr)
	return ret.Val
}

func (com *DaMo) AsmClear() int {
	ret, _ := com.iDispatch.CallMethod("AsmClear")
	return int(ret.Val)
}

func (com *DaMo) AsmSetTimeout(timeOut, param int) int {
	ret, _ := com.iDispatch.CallMethod("AsmSetTimeout", timeOut, param)
	return int(ret.Val)
}

func (com *DaMo) Assemble(timeOut int64, is64bit int) string {
	ret, _ := com.iDispatch.CallMethod("Assemble", timeOut, is64bit)
	return ret.ToString()
}

func (com *DaMo) DisAssemble(asmCode string, baseAddr int64, is64bit int) string {
	ret, _ := com.iDispatch.CallMethod("DisAssemble", asmCode, baseAddr, is64bit)
	return ret.ToString()
}
