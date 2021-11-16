package globalhotkey

import (
	"runtime"
	"testing"
)

func TestInit(t *testing.T) {

	// 将当前groutine固定到一个线程上执行，以便能正常调用os的api(注册热键api)
	runtime.LockOSThread()

	RegHotkey(&Hotkey{
		Modifiers: []int{ModNil},
		KeyCode:   VK_HOME,
		Handle: func() {
			println("HOME")
		},
	})
	RegHotkey(&Hotkey{
		Modifiers: nil,
		KeyCode:   VK_END,
		Handle: func() {
			println("END")
			QuitListenBlock()
		},
	})
	RegHotkey(&Hotkey{
		Modifiers: []int{ModAlt},
		KeyCode:   VK_KEY_X,
		Handle: func() {
			println("Alt+X")
			QuitListenBlock()
		},
	})
	RegHotkey(&Hotkey{
		Modifiers: []int{ModAlt, ModCtrl},
		KeyCode:   VK_KEY_Z,
		Handle: func() {
			println("Alt+Ctrl+Z")
			QuitListenBlock()
		},
	})

	err := ListenBlock()
	if err != nil {
		println(err)
	}

	println("bye")

	// 解锁
	runtime.UnlockOSThread()
}
