//go:build windows

package ts

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"os/exec"
)

type TianShiPlug struct {
	iDispatch *ole.IDispatch
	IUnknown  *ole.IUnknown
}

var modulePathAbs string

func (com *TianShiPlug) RegModule(modulePath string, regCode string, verInfo string) (err error) {
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
	com.IUnknown, err = oleutil.CreateObject("ts.tssoft")
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

func (com *TianShiPlug) UnRegModule() {
	com.IUnknown.Release()
	com.iDispatch.Release()
	ole.CoUninitialize()
	// 卸载模块
	cmd := exec.Command("regsvr32", modulePathAbs, "/u", "/s")
	cmd.Run()
}
