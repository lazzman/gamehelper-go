// 系统

package dm

func (com *DaMo) Beep(f, duration int) int {
	ret, _ := com.iDispatch.CallMethod("Beep", f, duration)
	return int(ret.Val)
}

func (com *DaMo) CheckFontSmooth() int {
	ret, _ := com.iDispatch.CallMethod("CheckFontSmooth")
	return int(ret.Val)
}

func (com *DaMo) CheckUAC() int {
	ret, _ := com.iDispatch.CallMethod("CheckUAC")
	return int(ret.Val)
}

func (com *DaMo) Delay(mis int) int {
	ret, _ := com.iDispatch.CallMethod("Delay", mis)
	return int(ret.Val)
}

func (com *DaMo) Delays(misMin, misMax int) int {
	ret, _ := com.iDispatch.CallMethod("Delays", misMin, misMax)
	return int(ret.Val)
}

// DisableCloseDisplayAndSleep 不支持xp系统
func (com *DaMo) DisableCloseDisplayAndSleep() int {
	ret, _ := com.iDispatch.CallMethod("DisableCloseDisplayAndSleep")
	return int(ret.Val)
}

func (com *DaMo) DisableFontSmooth() int {
	ret, _ := com.iDispatch.CallMethod("DisableFontSmooth")
	return int(ret.Val)
}

func (com *DaMo) DisablePowerSave() int {
	ret, _ := com.iDispatch.CallMethod("DisablePowerSave")
	return int(ret.Val)
}

func (com *DaMo) DisableScreenSave() int {
	ret, _ := com.iDispatch.CallMethod("DisableScreenSave")
	return int(ret.Val)
}

func (com *DaMo) EnableFontSmooth() int {
	ret, _ := com.iDispatch.CallMethod("EnableFontSmooth")
	return int(ret.Val)
}

func (com *DaMo) ExitOs(types int) int {
	ret, _ := com.iDispatch.CallMethod("ExitOs", types)
	return int(ret.Val)
}

func (com *DaMo) GetClipboard() string {
	ret, _ := com.iDispatch.CallMethod("GetClipboard")
	return ret.ToString()
}

func (com *DaMo) GetCPUType() int {
	ret, _ := com.iDispatch.CallMethod("GetCpuType")
	return int(ret.Val)
}

func (com *DaMo) GetCPUUsage() int {
	ret, _ := com.iDispatch.CallMethod("GetCpuUsage")
	return int(ret.Val)
}

func (com *DaMo) GetDir(types int) string {
	ret, _ := com.iDispatch.CallMethod("GetDir", types)
	return ret.ToString()
}

// GetDiskModel 需要管理员权限
func (com *DaMo) GetDiskModel(index int) string {
	ret, _ := com.iDispatch.CallMethod("GetDiskModel", index)
	return ret.ToString()
}

// GetDiskReversion 需要管理员权限
func (com *DaMo) GetDiskReversion(index int) string {
	ret, _ := com.iDispatch.CallMethod("GetDiskReversion", index)
	return ret.ToString()
}

// GetDiskSerial 需要管理员权限
func (com *DaMo) GetDiskSerial(index int) string {
	ret, _ := com.iDispatch.CallMethod("GetDiskSerial", index)
	return ret.ToString()
}

func (com *DaMo) GetDisplayInfo() string {
	ret, _ := com.iDispatch.CallMethod("GetDisplayInfo")
	return ret.ToString()
}

func (com *DaMo) GetDPI() int {
	ret, _ := com.iDispatch.CallMethod("GetDPI")
	return int(ret.Val)
}

func (com *DaMo) GetLocale() int {
	ret, _ := com.iDispatch.CallMethod("GetLocale")
	return int(ret.Val)
}

func (com *DaMo) GetMachineCode() string {
	ret, _ := com.iDispatch.CallMethod("GetMachineCode")
	return ret.ToString()
}

// GetMachineCodeNoMac 需要管理员权限
func (com *DaMo) GetMachineCodeNoMac() string {
	ret, _ := com.iDispatch.CallMethod("GetMachineCodeNoMac")
	return ret.ToString()
}

func (com *DaMo) GetMemoryUsage() int {
	ret, _ := com.iDispatch.CallMethod("GetMemoryUsage")
	return int(ret.Val)
}

func (com *DaMo) GetNetTime() string {
	ret, _ := com.iDispatch.CallMethod("GetNetTime")
	return ret.ToString()
}

func (com *DaMo) GetNetTimeByIP(ip string) string {
	ret, _ := com.iDispatch.CallMethod("GetNetTimeByIp", ip)
	return ret.ToString()
}

func (com *DaMo) GetOsBuildNumber() int {
	ret, _ := com.iDispatch.CallMethod("GetOsBuildNumber")
	return int(ret.Val)
}

func (com *DaMo) GetOsType() int {
	ret, _ := com.iDispatch.CallMethod("GetOsType")
	return int(ret.Val)
}

func (com *DaMo) GetScreenDepth() int {
	ret, _ := com.iDispatch.CallMethod("GetScreenDepth")
	return int(ret.Val)
}

func (com *DaMo) GetScreenHeight() int {
	ret, _ := com.iDispatch.CallMethod("GetScreenHeight")
	return int(ret.Val)
}

func (com *DaMo) GetScreenWidth() int {
	ret, _ := com.iDispatch.CallMethod("GetScreenWidth")
	return int(ret.Val)
}

func (com *DaMo) GetTime() int {
	ret, _ := com.iDispatch.CallMethod("GetTime")
	return int(ret.Val)
}

func (com *DaMo) Is64Bit() int {
	ret, _ := com.iDispatch.CallMethod("Is64Bit")
	return int(ret.Val)
}

func (com *DaMo) IsSurrpotVt() int {
	ret, _ := com.iDispatch.CallMethod("IsSurrpotVt")
	return int(ret.Val)
}

func (com *DaMo) Play(mediaFile string) int {
	ret, _ := com.iDispatch.CallMethod("Play", mediaFile)
	return int(ret.Val)
}

func (com *DaMo) RunApp(appPath string, mode int) int {
	ret, _ := com.iDispatch.CallMethod("RunApp", appPath, mode)
	return int(ret.Val)
}

func (com *DaMo) SetClipboard(value string) int {
	ret, _ := com.iDispatch.CallMethod("SetClipboard", value)
	return int(ret.Val)
}

func (com *DaMo) SetDisplayAcceler(level int) int {
	ret, _ := com.iDispatch.CallMethod("SetDisplayAcceler", level)
	return int(ret.Val)
}

func (com *DaMo) SetLocale() int {
	ret, _ := com.iDispatch.CallMethod("SetLocale")
	return int(ret.Val)
}

func (com *DaMo) SetScreen(width, height, depth int) int {
	ret, _ := com.iDispatch.CallMethod("SetScreen", width, height, depth)
	return int(ret.Val)
}

func (com *DaMo) SetUAC(enable int) int {
	ret, _ := com.iDispatch.CallMethod("SetUAC", enable)
	return int(ret.Val)
}

func (com *DaMo) ShowTaskBarIcon(hwnd, isShow int) int {
	ret, _ := com.iDispatch.CallMethod("ShowTaskBarIcon", hwnd, isShow)
	return int(ret.Val)
}

func (com *DaMo) Stop(id int) int {
	ret, _ := com.iDispatch.CallMethod("Stop", id)
	return int(ret.Val)
}
