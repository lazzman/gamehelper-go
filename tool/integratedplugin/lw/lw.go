//go:build windows

package lw

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"os/exec"
)

type LeWan struct {
	iDispatch *ole.IDispatch
	IUnknown  *ole.IUnknown
}

var VKCODEMAP = map[string]int{
	"1":      49,
	"2":      50,
	"3":      51,
	"4":      52,
	"5":      53,
	"6":      54,
	"7":      55,
	"8":      56,
	"9":      57,
	"0":      48,
	"-":      189,
	"=":      187,
	"back":   8,
	"a":      65,
	"b":      66,
	"c":      67,
	"d":      68,
	"e":      69,
	"f":      70,
	"g":      71,
	"h":      72,
	"i":      73,
	"j":      74,
	"k":      75,
	"l":      76,
	"m":      77,
	"n":      78,
	"o":      79,
	"p":      80,
	"q":      81,
	"r":      82,
	"s":      83,
	"t":      84,
	"u":      85,
	"v":      86,
	"w":      87,
	"x":      88,
	"y":      89,
	"z":      90,
	"ctrl":   17,
	"alt":    18,
	"shift":  16,
	"win":    91,
	"space":  32,
	"cap":    20,
	"tab":    9,
	"~":      192,
	"esc":    27,
	"enter":  13,
	"up":     38,
	"down":   40,
	"left":   37,
	"right":  39,
	"option": 93,
	"print":  44,
	"delete": 46,
	"home":   36,
	"end":    35,
	"pgup":   33,
	"pgdn":   34,
	"f1":     112,
	"f2":     113,
	"f3":     114,
	"f4":     115,
	"f5":     116,
	"f6":     117,
	"f7":     118,
	"f8":     119,
	"f9":     120,
	"f10":    121,
	"f11":    122,
	"f12":    123,
	"[":      219,
	"]":      221,
	"\\":     220,
	";":      186,
	"'":      222,
	":":      188,
	".":      190,
	"/":      191,
}

var modulePathAbs string

func (com *LeWan) RegModule(modulePath string, regCode string, verInfo string) (err error) {
	// 卸载模块
	cmd := exec.Command("regsvr32", modulePath, "/u", "/s")
	err = cmd.Run()

	// 注册模块
	cmd = exec.Command("regsvr32", modulePath, "/s")
	err = cmd.Run()
	if err != nil {
		return
	}
	modulePathAbs = modulePath

	// 创建COM对象
	err = ole.CoInitialize(0)
	if err != nil {
		return
	}
	com.IUnknown, err = oleutil.CreateObject("lw.lwsoft3")
	if err != nil {
		return
	}

	// 查询接口
	com.iDispatch, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return
	}

	return
}

func (com *LeWan) UnRegModule() {
	com.IUnknown.Release()
	com.iDispatch.Release()
	ole.CoUninitialize()
	// 卸载模块
	cmd := exec.Command("regsvr32", modulePathAbs, "/u", "/s")
	cmd.Run()
}
