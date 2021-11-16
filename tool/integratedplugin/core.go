package integratedplugin

import (
	_ "embed"
	"github.com/lazzman/gamehelper-go/tool/integratedplugin/dm"
	"github.com/lazzman/gamehelper-go/tool/integratedplugin/op"
	"github.com/lazzman/gamehelper-go/tool/integratedplugin/ts"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

const (
	DM = "dm" //大漠插件（推荐）
	OP = "op" //OP插件（目前插件不完善）
	TS = "ts" //天使插件（太老win10bug多）
)

//go:embed com/dm.dll
var dmDll []byte

//go:embed com/op.dll
var opDll []byte

//go:embed com/ts.dll
var tsDll []byte

type IntegratedPlugin interface {
	RegModule(modulePath string, regCode string, verInfo string) error //注册模块
	UnRegModule()                                                      //卸载模块
	Version() string                                                   //打印插件版本信息
	SetPath(path string) int                                           //设置全局路径,设置了此路径后,所有接口调用中,相关的文件都相对于此路径. 比如图片,字库等.
	EnablePicCache(enable int) int                                     //设置是否开启或者关闭插件内部的图片缓存机制. (默认是关闭).
	BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int
	UnBindWindow() int
	SetWindowState(hwnd, flag int) int
	SetWindowSize(hwnd, width, height int) int
	GetClientSize(hwnd int, width, height *int) int
	FindWindow(class, title string) int
	FindWindowEx(parent int, class, title string) int
	MoveWindow(hwnd, x, y int) int
	SetKeypadDelay(types string, delay int) int
	SetMouseDelay(types string, delay int) int
	EnableIme(enable int) int
	EnableRealKeypad(enable int) int
	EnableRealMouse(enable, mousedelay, mousestep int) int
	LeftClick() int
	LeftDoubleClick() int
	KeyPressChar(keyStr string) int
	KeyPressStr(keyStr string, delay int) int
	KeyDownChar(keyStr string) int
	KeyUpChar(keyStr string) int
	SetMouseSpeed(speed int) int
	EnableMouseAccuracy(enable int) int
	MoveR(rx, ry int) int
	MoveTo(x, y int) int
	WheelUp() int
	WheelDown() int
	FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int
	FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string
	FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int
	FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string
	Capture(x1, y1, x2, y2 int, file string) int
}

func Instance(plug string, args ...string) IntegratedPlugin {
	var iPlugin IntegratedPlugin
	if strings.EqualFold(plug, DM) {
		err := ioutil.WriteFile(filepath.Join("plugin.dll"), dmDll, 0777)
		if err != nil {
			log.Panicln("导出插件失败", err)
		}
		abs, err := filepath.Abs(filepath.Join("plugin.dll"))
		if err != nil {
			log.Panicln("综合插件路径错误", err)
		}
		iPlugin = &dm.DaMo{}
		err = iPlugin.RegModule(abs, args[0], args[1])
		if err != nil {
			log.Panicln("综合插件初始化失败", err)
		}
	} else if strings.EqualFold(plug, OP) {
		err := ioutil.WriteFile(filepath.Join("plugin.dll"), opDll, 0777)
		if err != nil {
			log.Panicln("导出插件失败", err)
		}
		abs, err := filepath.Abs(filepath.Join("plugin.dll"))
		if err != nil {
			log.Panicln("综合插件路径错误", err)
		}
		iPlugin = &op.Opsoft{}
		err = iPlugin.RegModule(abs, "", "")
		if err != nil {
			log.Panicln("综合插件初始化失败", err)
		}
	} else if strings.EqualFold(plug, TS) {
		err := ioutil.WriteFile(filepath.Join("plugin.dll"), tsDll, 0777)
		if err != nil {
			log.Panicln("导出插件失败", err)
		}
		abs, err := filepath.Abs(filepath.Join("plugin.dll"))
		if err != nil {
			log.Panicln("综合插件路径错误", err)
		}
		iPlugin = &ts.TianShiPlug{}
		err = iPlugin.RegModule(abs, "", "")
		if err != nil {
			log.Panicln("综合插件初始化失败", err)
		}
	}
	return iPlugin
}
