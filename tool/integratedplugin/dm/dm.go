//go:build windows

package dm

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"os/exec"
)

type DaMo struct {
	iDispatch *ole.IDispatch
	IUnknown  *ole.IUnknown
}

var modulePathAbs string

func (com *DaMo) RegModule(modulePath string, regCode string, verInfo string) (err error) {
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
	com.IUnknown, err = oleutil.CreateObject("dm.dmsoft")
	if err != nil {
		return
	}

	// 查询接口
	com.iDispatch, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return
	}

	// 激活码注册
	ret := com.Reg(regCode, verInfo)
	if int(ret) != 1 {
		return fmt.Errorf("大漠插件注册失败，返回值：%v", ret)
	}

	return
}

func (com *DaMo) UnRegModule() {
	com.IUnknown.Release()
	com.iDispatch.Release()
	ole.CoUninitialize()
	// 卸载模块
	cmd := exec.Command("regsvr32", modulePathAbs, "/u", "/s")
	cmd.Run()
}
