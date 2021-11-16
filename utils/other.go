package utils

import (
	"errors"
	"time"
)

func DelayMills(mills int) {
	time.Sleep(time.Millisecond * time.Duration(mills))
}

type FuncTry func() bool

var OutMaxTryTimesError = errors.New("超过最大重试次数")

// RetryRunByTimes 重复执行 FuncTry ，最大重试 times 次，若 FuncTry 返回true 则停止重试，error 为 nil.
// 若超过最大重试次数则将返回 OutMaxTryTimesError.
func RetryRunByTimes(try FuncTry, maxTryTimes int) (err error) {
	for i := 0; i < maxTryTimes; i++ {
		if try() {
			return nil
		}
	}
	err = OutMaxTryTimesError
	return err
}

// RetryRunByTimesAndDelay 重复执行 FuncTry ，最大重试 times 次，若 FuncTry 返回true 则停止重试，error 为 nil.
// 若超过最大重试次数则将返回 OutMaxTryTimesError.
func RetryRunByTimesAndDelay(try FuncTry, maxTryTimes int, delayMills int) (err error) {
	for i := 0; i < maxTryTimes; i++ {
		if try() {
			return nil
		}
		time.Sleep(time.Millisecond * time.Duration(delayMills))
	}
	err = OutMaxTryTimesError
	return err
}

var TimeOutRetryError = errors.New("重试已超时")

// RetryRunByTimeout 尝试运行 FuncTry，直到 FuncTry 返回true 或者 超时为止.
func RetryRunByTimeout(try FuncTry, timeOutMills int, delayMills int) (err error) {
	timeOutTime := time.Now().UnixMilli() + int64(timeOutMills)
	for true {
		if try() {
			return nil
		} else {
			time.Sleep(time.Millisecond * time.Duration(delayMills))
		}
		if time.Now().UnixMilli() > timeOutTime {
			err = TimeOutRetryError
			break
		}
	}
	return err
}
