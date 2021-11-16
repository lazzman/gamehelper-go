package integratedplugin

import (
	"gamehelper/tool/globalhotkey"
	"gamehelper/utils"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

// TestGameAimBooster http://www.aimbooster.com/
func TestGameAimBooster(t *testing.T) {
	// 选择测试的插件
	iPlugin := Instance(DM)
	// 绑定窗口
	hwnd := iPlugin.FindWindow("", "AimBooster - Google Chrome")
	if hwnd == 0 {
		println("查找窗口失败")
		return
	} else {
		println("查找窗口成功")
	}
	bind := iPlugin.BindWindow(hwnd, "normal", "windows", "windows", 0)
	if bind == 0 {
		println("绑定窗口失败")
		return
	} else {
		println("绑定窗口成功")
	}
	iPlugin.EnablePicCache(0)
	// 最大化窗口置顶 1920*1080分辨率
	iPlugin.SetWindowSize(hwnd, 1920, 1080)
	iPlugin.SetWindowState(hwnd, 1)

	flag := true
	ss := func() {
		// 找色
		var x, y int
		for flag {
			if iPlugin.FindColor(659, 369, 1261, 790, "ff954e-050505", 1.0, 0, &x, &y) == 1 {
				iPlugin.MoveTo(x, y)
				iPlugin.LeftClick()
				//iPlugin.Capture(659, 369, 1261, 790, "screen.bmp")
				//return
			}
			utils.DelayMills(100)
		}
	}

	runtime.LockOSThread()
	globalhotkey.RegHotkey(&globalhotkey.Hotkey{
		Modifiers: nil,
		KeyCode:   globalhotkey.VK_HOME,
		Handle: func() {
			flag = true
			ss()
		},
	})
	globalhotkey.RegHotkey(&globalhotkey.Hotkey{
		Modifiers: nil,
		KeyCode:   globalhotkey.VK_END,
		Handle: func() {
			flag = false
		},
	})
	globalhotkey.ListenBlock()
	runtime.UnlockOSThread()
}

// TestOps http://www.aimbooster.com/
func TestOps(t *testing.T) {
	// 选择测试的插件
	iPlugin := Instance(TS)
	// 绑定窗口
	hwnd := iPlugin.FindWindow("", "AimBooster - Google Chrome")
	if hwnd == 0 {
		println("查找窗口失败")
		return
	} else {
		println("查找窗口成功")
	}
	bind := iPlugin.BindWindow(hwnd, "normal", "normal", "normal", 0)
	if bind == 0 {
		println("绑定窗口失败")
		return
	} else {
		println("绑定窗口成功")
	}
	iPlugin.EnablePicCache(0)
	// 最大化窗口置顶 1920*1080分辨率
	iPlugin.SetWindowSize(hwnd, 1920, 1080)
	iPlugin.SetWindowState(hwnd, 1)
	iPlugin.EnableRealMouse(1, 10, 50)
	//保证大漠插件与天使插件鼠标MoveR效果相同需要将鼠标精准度开关关闭
	//iPlugin.SetMouseSpeed(6)      // 保持移动速度是居中的位置
	iPlugin.EnableMouseAccuracy(0) // 关闭精确度开关

	ss := func() {
		iPlugin.MoveTo(1000, 500)
		iPlugin.MoveR(500, 0)
		iPlugin.MoveR(0, 500)
		iPlugin.MoveR(-500, 0)
		iPlugin.MoveR(0, -500)
		iPlugin.MoveR(-500, -500)
		iPlugin.MoveR(500, 500)
	}

	runtime.LockOSThread()
	globalhotkey.RegHotkey(&globalhotkey.Hotkey{
		Modifiers: nil,
		KeyCode:   globalhotkey.VK_HOME,
		Handle: func() {
			ss()
		},
	})
	globalhotkey.ListenBlock()
	runtime.UnlockOSThread()
}

func TestTheDivision2(t *testing.T) {
	// 选择测试的插件
	ip := Instance(TS)
	title := "Tom Clancy's The Division 2"
	hwnd := ip.FindWindow("", title)
	for ; hwnd <= 0; hwnd = ip.FindWindow("", title) {
		time.Sleep(time.Second * 1)
	}
	// 绑定窗口
	if ip.BindWindow(hwnd, "normal", "normal", "normal", 0) == 0 {
		println("目标窗口绑定失败，脚本退出")
	}
	// 校验窗口大小
	println("目标分辨率须设置为1366*768")
	var width, height int
	if ip.GetClientSize(hwnd, &width, &height) == 0 {
		println("获取目标窗口大小失败，脚本退出")
	}
	// 插件配置
	// 设置全局资源路径
	ip.SetPath(filepath.Join("resource"))
	// 关闭图片缓存
	ip.EnablePicCache(0)
	// 关闭输入法
	ip.EnableIme(0)
	// 键盘延迟不能过小，否则会出现按键失灵
	ip.SetKeypadDelay("normal", 30)
	// 开启鼠键模拟
	ip.EnableRealMouse(1, 10, 50)
	ip.EnableRealKeypad(0)
	// 关闭鼠标加速
	ip.EnableMouseAccuracy(0)
	ip.SetMouseSpeed(6)

	ss := func() {
		ip.SetWindowState(hwnd, 1)
		ip.MoveTo(10, 10)
		ip.MoveR(1000, 0) //天使插件在全境2中MoveR工作异常
	}

	runtime.LockOSThread()
	globalhotkey.RegHotkey(&globalhotkey.Hotkey{
		Modifiers: nil,
		KeyCode:   globalhotkey.VK_HOME,
		Handle: func() {
			ss()
		},
	})
	globalhotkey.ListenBlock()
	runtime.UnlockOSThread()
}
