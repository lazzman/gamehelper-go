// 其他

package dm

func (com *DaMo) DmGuard(enable, lType int) int {
	ret, _ := com.iDispatch.CallMethod("DmGuard", enable, lType)
	return int(ret.Val)
}

func (com *DaMo) DmGuardParams(cmd, subcmd, param string) string {
	ret, _ := com.iDispatch.CallMethod("DmGuardParams", cmd, subcmd, param)
	return ret.ToString()
}

func (com *DaMo) UnLoadDriver() int {
	ret, _ := com.iDispatch.CallMethod("UnLoadDriver")
	return int(ret.Val)
}
