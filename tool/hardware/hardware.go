//go:build windows

package hardware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/yusufpapurcu/wmi"
)

var (
	kernel = syscall.NewLazyDLL("Kernel32.dll")
)

func main() {
	fmt.Printf("开机时长:%s\n", GetStartTime())
	fmt.Printf("当前用户:%s\n", GetUserName())
	fmt.Printf("当前系统:%s\n", runtime.GOOS)
	fmt.Printf("系统版本:%s\n", GetSystemVersion())
	fmt.Printf("%s\n", GetBiosInfo())
	fmt.Printf("Motherboard:\t%s\n", GetMotherboardInfo())

	fmt.Printf("CPU:\t%s\n", GetCpuInfo())
	fmt.Printf("Memory:\t%s\n", GetMemory())
	fmt.Printf("Disk:\t%v\n", GetDiskInfo())
	fmt.Printf("Interfaces:\t%v\n", GetInfs())
}

// GetStartTime 开机时间
func GetStartTime() string {
	GetTickCount := kernel.NewProc("GetTickCount")
	r, _, _ := GetTickCount.Call()
	if r == 0 {
		return ""
	}
	ms := time.Duration(r * 1000 * 1000)
	return ms.String()
}

// GetUserName 当前用户名
func GetUserName() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	user := syscall.StringToUTF16Ptr("USERNAME")
	domain := syscall.StringToUTF16Ptr("USERDOMAIN")
	r, err := syscall.GetEnvironmentVariable(user, &buffer[0], size)
	if err != nil {
		return ""
	}
	buffer[r] = '@'
	old := r + 1
	if old >= size {
		return syscall.UTF16ToString(buffer[:r])
	}
	r, err = syscall.GetEnvironmentVariable(domain, &buffer[old], size-old)
	return syscall.UTF16ToString(buffer[:old+r])
}

// GetSystemVersion 系统版本
func GetSystemVersion() string {
	version, err := syscall.GetVersion()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d.%d (%d)", byte(version), uint8(version>>8), version>>16)
}

type diskusage struct {
	Path  string `json:"path"`
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
}

func usage(getDiskFreeSpaceExW *syscall.LazyProc, path string) (diskusage, error) {
	lpFreeBytesAvailable := int64(0)
	var info = diskusage{Path: path}
	disk, _, err := getDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(info.Path))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&(info.Total))),
		uintptr(unsafe.Pointer(&(info.Free))))
	if disk != 0 {
		err = nil
	}
	return info, err
}

// GetDiskInfo 硬盘信息
func GetDiskInfo() (infos []diskusage) {
	GetLogicalDriveStringsW := kernel.NewProc("GetLogicalDriveStringsW")
	GetDiskFreeSpaceExW := kernel.NewProc("GetDiskFreeSpaceExW")
	lpBuffer := make([]byte, 254)
	disk, _, _ := GetLogicalDriveStringsW.Call(
		uintptr(len(lpBuffer)),
		uintptr(unsafe.Pointer(&lpBuffer[0])))
	if disk == 0 {
		return
	}
	for _, v := range lpBuffer {
		if v >= 65 && v <= 90 {
			path := string(v) + ":"
			if path == "A:" || path == "B:" {
				continue
			}
			info, err := usage(GetDiskFreeSpaceExW, string(v)+":")
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}
	}
	return infos
}

// GetCpuInfo CPU信息
func GetCpuInfo() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	var index = uint32(copy(buffer, syscall.StringToUTF16("Num:")) - 1)
	nums := syscall.StringToUTF16Ptr("NUMBER_OF_PROCESSORS")
	arch := syscall.StringToUTF16Ptr("PROCESSOR_ARCHITECTURE")
	r, err := syscall.GetEnvironmentVariable(nums, &buffer[index], size-index)
	if err != nil {
		return ""
	}
	index += r
	index += uint32(copy(buffer[index:], syscall.StringToUTF16(" Arch:")) - 1)
	r, err = syscall.GetEnvironmentVariable(arch, &buffer[index], size-index)
	if err != nil {
		return syscall.UTF16ToString(buffer[:index])
	}
	index += r
	return syscall.UTF16ToString(buffer[:index+r])
}

type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

// GetMemory 内存信息
func GetMemory() string {
	GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return ""
	}
	return fmt.Sprint(memInfo.ullTotalPhys / (1024 * 1024))
}

type infInfo struct {
	Name string
	Ipv4 []string
	Ipv6 []string
}

// GetInfs 网卡信息
func GetInfs() []infInfo {
	inf, err := net.Interfaces()
	if err != nil {
		return []infInfo{}
	}
	var is = make([]infInfo, len(inf))
	for i, v := range inf {
		ips, err := v.Addrs()
		if err != nil {
			continue
		}
		is[i].Name = v.Name
		for _, ip := range ips {
			if strings.Contains(ip.String(), ":") {
				is[i].Ipv6 = append(is[i].Ipv6, ip.String())
			} else {
				is[i].Ipv4 = append(is[i].Ipv4, ip.String())
			}
		}
	}
	return is
}

type Motherboard struct {
	Caption                 string
	ConfigOptions           []string
	CreationClassName       string
	Description             string
	Manufacturer            string
	Model                   string
	Name                    string
	OtherIdentifyingInfo    string
	PartNumber              string
	Product                 string
	RequirementsDescription string
	SerialNumber            string
	SKU                     string
	SlotLayout              string
	Status                  string
	Tag                     string
	Version                 string
}

// GetMotherboardInfo 主板信息
func GetMotherboardInfo() *Motherboard {
	var s []Motherboard
	err := wmi.Query("SELECT Caption,ConfigOptions,CreationClassName,Description,Manufacturer,Model,Name,OtherIdentifyingInfo,PartNumber,Product,RequirementsDescription,SerialNumber,SKU,SlotLayout,Status,Tag,Version FROM Win32_BaseBoard", &s)
	if err != nil {
		return nil
	}
	return &s[0]
}

type Bios struct {
	BIOSVersion        []string
	BuildNumber        string
	Caption            string
	CodeSet            string
	CurrentLanguage    string
	Description        string
	IdentificationCode string
	LanguageEdition    string
	ListOfLanguages    []string
	Manufacturer       string
	Name               string
	OtherTargetOS      string
	SerialNumber       string
	SMBIOSBIOSVersion  string
	SoftwareElementID  string
	Status             string
	Version            string
}

// GetBiosInfo BIOS信息
func GetBiosInfo() *Bios {
	var s []Bios
	err := wmi.Query("SELECT BIOSVersion,BuildNumber,Caption,CodeSet,CurrentLanguage,Description,IdentificationCode,LanguageEdition,ListOfLanguages,Manufacturer,Name,OtherTargetOS,SerialNumber,SMBIOSBIOSVersion,SoftwareElementID,Status,Version FROM Win32_BIOS", &s) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return nil
	}
	return &s[0]
}

type ComputerSystemProduct struct {
	Caption           string
	Description       string
	IdentifyingNumber string
	Name              string
	SKUNumber         string
	Vendor            string
	Version           string
	UUID              string
}

func GetComputerSystemProduct() *ComputerSystemProduct {
	var s []ComputerSystemProduct
	err := wmi.Query("SELECT Caption,Description,IdentifyingNumber,Name,SKUNumber,Vendor,Version,UUID FROM Win32_ComputerSystemProduct", &s) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return nil
	}
	return &s[0]
}

func GetUUID() string {
	computerSystemProduct := GetComputerSystemProduct()
	uuid := computerSystemProduct.UUID
	identifyingNumber := computerSystemProduct.IdentifyingNumber
	hash := sha256.New()
	if uuid != "" {
		hash.Write([]byte(uuid))
	}
	if identifyingNumber != "" {
		hash.Write([]byte(identifyingNumber))
	}
	hash.Write([]byte("GameHelper"))
	return hex.EncodeToString(hash.Sum(nil))
}
