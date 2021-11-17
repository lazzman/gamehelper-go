// 汇编

package lw

func (com *LeWan) AsmAdd(asmIns string) int {
	ret, _ := com.iDispatch.CallMethod("AsmAdd", asmIns)
	return int(ret.Val)
}

func (com *LeWan) AsmCall(hwnd, mode int) int {
	ret, _ := com.iDispatch.CallMethod("AsmCall", hwnd, mode)
	return int(ret.Val)
}

func (com *LeWan) AsmClear() int {
	ret, _ := com.iDispatch.CallMethod("AsmClear")
	return int(ret.Val)
}

func (com *LeWan) Assemble(timeOut int64, is64bit int) string {
	ret, _ := com.iDispatch.CallMethod("Assemble", timeOut, is64bit)
	return ret.ToString()
}
