package op

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"os/exec"
)

type Opsoft struct {
	iDispatch *ole.IDispatch
	IUnknown  *ole.IUnknown
}

var modulePathAbs string

func (com *Opsoft) RegModule(modulePath string, regCode string, verInfo string) (err error) {
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
	com.IUnknown, err = oleutil.CreateObject("op.opsoft")
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

func (com *Opsoft) UnRegModule() {
	com.IUnknown.Release()
	com.iDispatch.Release()
	ole.CoUninitialize()
	// 卸载模块
	cmd := exec.Command("regsvr32", modulePathAbs, "/u", "/s")
	cmd.Run()
}
