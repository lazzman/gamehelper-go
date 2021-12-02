package utils

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetHWInfo(t *testing.T) {
	fmt.Printf("开机时长:\t%s\n", GetStartTime())
	fmt.Printf("当前用户:\t%s\n", GetUserName())
	fmt.Printf("当前系统:\t%s\n", runtime.GOOS)
	fmt.Printf("系统版本:\t%s\n", GetSystemVersion())
	fmt.Printf("CPU:\t%s\n", GetCpuInfo())
	fmt.Printf("Memory:\t%s\n", GetMemory())
	fmt.Printf("Disk:\t%v\n", GetDiskInfo())
	fmt.Printf("Interfaces:\t%v\n", GetInfs())
	fmt.Printf("BIOS:\t%v\n", GetBiosInfo())
	fmt.Printf("Motherboard:\t%v\n", GetMotherboardInfo())
	fmt.Printf("ComputerSystemProduct:\t%v\n", GetComputerSystemProduct())
	fmt.Printf("UUID:\t%s\n", GetUUID())
}
