//go:build windows

package globalhotkey

import (
	"errors"
	"github.com/gonutz/w32/v2"
	"sync/atomic"
	"syscall"
)

// 修饰键，可以多个任意组合
const (
	ModAlt   = 1 << iota //Alt键码 1
	ModCtrl              //Ctrl键码 2
	ModShift             //Shift键码 4
	ModWin               //Win键码 8
	ModNil   = 0         //不使用修饰键 0
)

// Virtual-Key Codes.
// 参考 https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes.
const (
	VK_LBUTTON             = 0x01
	VK_RBUTTON             = 0x02
	VK_CANCEL              = 0x03
	VK_MBUTTON             = 0x04
	VK_XBUTTON1            = 0x05
	VK_XBUTTON2            = 0x06
	VK_BACK                = 0x08
	VK_TAB                 = 0x09
	VK_CLEAR               = 0x0C
	VK_RETURN              = 0x0D
	VK_SHIFT               = 0x10
	VK_CONTROL             = 0x11
	VK_MENU                = 0x12
	VK_PAUSE               = 0x13
	VK_CAPITAL             = 0x14
	VK_KANA                = 0x15
	VK_HANGEUL             = 0x15
	VK_HANGUL              = 0x15
	VK_IME_ON              = 0x16
	VK_JUNJA               = 0x17
	VK_FINAL               = 0x18
	VK_HANJA               = 0x19
	VK_KANJI               = 0x19
	VK_IME_OFF             = 0x1A
	VK_ESCAPE              = 0x1B
	VK_CONVERT             = 0x1C
	VK_NONCONVERT          = 0x1D
	VK_ACCEPT              = 0x1E
	VK_MODECHANGE          = 0x1F
	VK_SPACE               = 0x20
	VK_PRIOR               = 0x21 //PAGE UP key
	VK_NEXT                = 0x22 //PAGE DOWN key
	VK_END                 = 0x23
	VK_HOME                = 0x24
	VK_LEFT                = 0x25
	VK_UP                  = 0x26
	VK_RIGHT               = 0x27
	VK_DOWN                = 0x28
	VK_SELECT              = 0x29
	VK_PRINT               = 0x2A
	VK_EXECUTE             = 0x2B
	VK_SNAPSHOT            = 0x2C
	VK_INSERT              = 0x2D
	VK_DELETE              = 0x2E
	VK_HELP                = 0x2F
	VK_KEY_0               = 0x30
	VK_KEY_1               = 0x31
	VK_KEY_2               = 0x32
	VK_KEY_3               = 0x33
	VK_KEY_4               = 0x34
	VK_KEY_5               = 0x35
	VK_KEY_6               = 0x36
	VK_KEY_7               = 0x37
	VK_KEY_8               = 0x38
	VK_KEY_9               = 0x39
	VK_KEY_A               = 0x41
	VK_KEY_B               = 0x42
	VK_KEY_C               = 0x43
	VK_KEY_D               = 0x44
	VK_KEY_E               = 0x45
	VK_KEY_F               = 0x46
	VK_KEY_G               = 0x47
	VK_KEY_H               = 0x48
	VK_KEY_I               = 0x49
	VK_KEY_J               = 0x4A
	VK_KEY_K               = 0x4B
	VK_KEY_L               = 0x4C
	VK_KEY_M               = 0x4D
	VK_KEY_N               = 0x4E
	VK_KEY_O               = 0x4F
	VK_KEY_P               = 0x50
	VK_KEY_Q               = 0x51
	VK_KEY_R               = 0x52
	VK_KEY_S               = 0x53
	VK_KEY_T               = 0x54
	VK_KEY_U               = 0x55
	VK_KEY_V               = 0x56
	VK_KEY_W               = 0x57
	VK_KEY_X               = 0x58
	VK_KEY_Y               = 0x59
	VK_KEY_Z               = 0x5A
	VK_LWIN                = 0x5B
	VK_RWIN                = 0x5C
	VK_APPS                = 0x5D
	VK_SLEEP               = 0x5F
	VK_NUMPAD0             = 0x60
	VK_NUMPAD1             = 0x61
	VK_NUMPAD2             = 0x62
	VK_NUMPAD3             = 0x63
	VK_NUMPAD4             = 0x64
	VK_NUMPAD5             = 0x65
	VK_NUMPAD6             = 0x66
	VK_NUMPAD7             = 0x67
	VK_NUMPAD8             = 0x68
	VK_NUMPAD9             = 0x69
	VK_MULTIPLY            = 0x6A
	VK_ADD                 = 0x6B
	VK_SEPARATOR           = 0x6C
	VK_SUBTRACT            = 0x6D
	VK_DECIMAL             = 0x6E
	VK_DIVIDE              = 0x6F
	VK_F1                  = 0x70
	VK_F2                  = 0x71
	VK_F3                  = 0x72
	VK_F4                  = 0x73
	VK_F5                  = 0x74
	VK_F6                  = 0x75
	VK_F7                  = 0x76
	VK_F8                  = 0x77
	VK_F9                  = 0x78
	VK_F10                 = 0x79
	VK_F11                 = 0x7A
	VK_F12                 = 0x7B
	VK_F13                 = 0x7C
	VK_F14                 = 0x7D
	VK_F15                 = 0x7E
	VK_F16                 = 0x7F
	VK_F17                 = 0x80
	VK_F18                 = 0x81
	VK_F19                 = 0x82
	VK_F20                 = 0x83
	VK_F21                 = 0x84
	VK_F22                 = 0x85
	VK_F23                 = 0x86
	VK_F24                 = 0x87
	VK_NUMLOCK             = 0x90
	VK_SCROLL              = 0x91
	VK_OEM_NEC_EQUAL       = 0x92
	VK_OEM_FJ_JISHO        = 0x92
	VK_OEM_FJ_MASSHOU      = 0x93
	VK_OEM_FJ_TOUROKU      = 0x94
	VK_OEM_FJ_LOYA         = 0x95
	VK_OEM_FJ_ROYA         = 0x96
	VK_LSHIFT              = 0xA0
	VK_RSHIFT              = 0xA1
	VK_LCONTROL            = 0xA2
	VK_RCONTROL            = 0xA3
	VK_LMENU               = 0xA4
	VK_RMENU               = 0xA5
	VK_BROWSER_BACK        = 0xA6
	VK_BROWSER_FORWARD     = 0xA7
	VK_BROWSER_REFRESH     = 0xA8
	VK_BROWSER_STOP        = 0xA9
	VK_BROWSER_SEARCH      = 0xAA
	VK_BROWSER_FAVORITES   = 0xAB
	VK_BROWSER_HOME        = 0xAC
	VK_VOLUME_MUTE         = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
	VK_MEDIA_NEXT_TRACK    = 0xB0
	VK_MEDIA_PREV_TRACK    = 0xB1
	VK_MEDIA_STOP          = 0xB2
	VK_MEDIA_PLAY_PAUSE    = 0xB3
	VK_LAUNCH_MAIL         = 0xB4
	VK_LAUNCH_MEDIA_SELECT = 0xB5
	VK_LAUNCH_APP1         = 0xB6
	VK_LAUNCH_APP2         = 0xB7
	VK_OEM_1               = 0xBA
	VK_OEM_PLUS            = 0xBB
	VK_OEM_COMMA           = 0xBC
	VK_OEM_MINUS           = 0xBD
	VK_OEM_PERIOD          = 0xBE
	VK_OEM_2               = 0xBF
	VK_OEM_3               = 0xC0
	VK_OEM_4               = 0xDB
	VK_OEM_5               = 0xDC
	VK_OEM_6               = 0xDD
	VK_OEM_7               = 0xDE
	VK_OEM_8               = 0xDF
	VK_OEM_AX              = 0xE1
	VK_OEM_102             = 0xE2
	VK_ICO_HELP            = 0xE3
	VK_ICO_00              = 0xE4
	VK_PROCESSKEY          = 0xE5
	VK_ICO_CLEAR           = 0xE6
	VK_PACKET              = 0xE7
	VK_OEM_RESET           = 0xE9
	VK_OEM_JUMP            = 0xEA
	VK_OEM_PA1             = 0xEB
	VK_OEM_PA2             = 0xEC
	VK_OEM_PA3             = 0xED
	VK_OEM_WSCTRL          = 0xEE
	VK_OEM_CUSEL           = 0xEF
	VK_OEM_ATTN            = 0xF0
	VK_OEM_FINISH          = 0xF1
	VK_OEM_COPY            = 0xF2
	VK_OEM_AUTO            = 0xF3
	VK_OEM_ENLW            = 0xF4
	VK_OEM_BACKTAB         = 0xF5
	VK_ATTN                = 0xF6
	VK_CRSEL               = 0xF7
	VK_EXSEL               = 0xF8
	VK_EREOF               = 0xF9
	VK_PLAY                = 0xFA
	VK_ZOOM                = 0xFB
	VK_NONAME              = 0xFC
	VK_PA1                 = 0xFD
	VK_OEM_CLEAR           = 0xFE
)

var ops int32 = 0

//genId 线程安全的生成自增ID
func genId() int {
	atomic.AddInt32(&ops, 1)
	return int(atomic.LoadInt32(&ops))
}

type HandleHotKey func()

// Hotkey 热键结构体
type Hotkey struct {
	id        int          // 热键标识符
	Modifiers []int        // 修饰符键值 可以由多个修饰键组合相加
	modifiers int          // 修饰符键值累加值
	KeyCode   int          // 虚拟键码
	Handle    HandleHotKey // 热键触发时异步调用该函数，即在新的groutine中执行该函数
}

// 存放所有注册的热键
var hotkeyMap = make(map[int]*Hotkey, 10)

// RegHotkey 注册全局热键
func RegHotkey(hotkey *Hotkey) error {
	hotkey.id = genId()
	if hotkey.Modifiers == nil {
		hotkey.Modifiers = []int{ModNil}
	}
	for _, modifier := range hotkey.Modifiers {
		hotkey.modifiers += modifier
	}
	r1, _, err := registerHotkey.Call(
		0, uintptr(hotkey.id), uintptr(hotkey.modifiers), uintptr(hotkey.KeyCode))
	if r1 == 1 {
		hotkeyMap[hotkey.id] = hotkey
		return nil
	} else {
		return err
	}
}

var blocking int32 = 0  // 是否正在阻塞监听
var listening int32 = 0 // 是否正在监听热键

// ListenBlock 监听热键事件，调用该方法将阻塞当前groutine.
// 由于GPM模型与windows的消息机制冲突（windows的消息机制是向注册监听事件的线程的事件队列中传递消息，但是GPM中groutine底层的线程会变动，导致无法获取到线程队列中的消息）
func ListenBlock() error {
	if atomic.LoadInt32(&listening) > 0 {
		// 如果已经被其他线程监听了则不允许监听
		return errors.New("其他groutine正在监听")
	}
	atomic.StoreInt32(&listening, 1)
	atomic.StoreInt32(&blocking, 1)
	for {
		var msg = &w32.MSG{}
		// 选择PeekMessage而不是GetMessage的原因是PeekMessage不会阻塞线程，这样可以判断是否要退出当前方法了
		// 最后还是选择GetMessage,PeekMessage性能太差
		// 如果使用 GetMessage，QuitListenBlock 只会在下一次触发热键时生效，因为监听热键的groutine被阻塞在系统调用中，只有再次触发热键让系统调用返回后才可以退出监听。
		w32.GetMessage(msg, 0, 0, 0)
		if atomic.LoadInt32(&blocking) == 0 {
			break
		}
		// 注册的热键 id 对应通知消息中的 WParam 字段
		if id := msg.WParam; id != 0 {
			if hk, ok := hotkeyMap[int(id)]; ok {
				go hk.Handle()
			}
		}
	}
	atomic.StoreInt32(&blocking, 0)
	atomic.StoreInt32(&listening, 0)
	return nil
}

// QuitListenBlock 退出阻塞监听热键.
func QuitListenBlock() {
	atomic.StoreInt32(&blocking, 0)
	for _, hotkey := range hotkeyMap {
		var keys []int
		if len(hotkey.Modifiers) > 0 {
			for _, modifier := range hotkey.Modifiers {
				keys = append(keys, modifier)
			}
		}
		keys = append(keys, hotkey.KeyCode)
		tapKey(keys...)
	}
}

var (
	user32           = syscall.NewLazyDLL("user32")
	registerHotkey   = user32.NewProc("RegisterHotKey")
	unregisterHotkey = user32.NewProc("UnregisterHotKey")
)

type WIN32MSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}

func tapKey(keys ...int) {
	var inputs []w32.INPUT
	for _, key := range keys {
		ip := w32.KeyboardInput(w32.KEYBDINPUT{
			Vk:        uint16(key),
			Scan:      0,
			Flags:     0,
			Time:      0,
			ExtraInfo: 0,
		})
		inputs = append(inputs, ip)
	}
	for _, key := range keys {
		ip := w32.KeyboardInput(w32.KEYBDINPUT{
			Vk:        uint16(key),
			Scan:      0,
			Flags:     w32.KEYEVENTF_KEYUP,
			Time:      0,
			ExtraInfo: 0,
		})
		inputs = append(inputs, ip)
	}
	w32.SendInput(inputs...)
}
