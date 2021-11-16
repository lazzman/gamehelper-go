// 未实现功能

package op

import "time"

func (com *Opsoft) SetKeypadDelay(types string, delay int) int {
	return 1
}

func (com *Opsoft) SetMouseDelay(types string, delay int) int {
	return 1
}

func (com *Opsoft) EnableIme(enable int) int {
	return 1
}

func (com *Opsoft) EnableRealKeypad(enable int) int {
	return 1
}

func (com *Opsoft) EnableRealMouse(enable, mousedelay, mousestep int) int {
	return 1
}

func (com *Opsoft) KeyPressStr(keyStrs string, delay int) int {
	var ret int
	for i := range keyStrs {
		keyStr := keyStrs[i]
		ret = com.KeyPressChar(string([]byte{keyStr}))
		if ret == 0 {
			return ret
		}
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
	return ret
}

func (com *Opsoft) SetMouseSpeed(speed int) int {
	return 1
}

func (com *Opsoft) EnableMouseAccuracy(enable int) int {
	return 1
}
