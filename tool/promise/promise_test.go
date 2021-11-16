package promise

import (
	"errors"
	"runtime"
	"testing"
	"time"
)

func TestPromise_All(t *testing.T) {
	println(runtime.NumGoroutine())
	p := NewPromise(func() interface{} {
		time.Sleep(time.Second * 2)
		return "hello"
	}, func() interface{} {
		time.Sleep(time.Second * 3)
		return "hi"
	})
	m, err := p.All(3000)
	if err != nil && errors.Is(err, TimeoutError) {
		println("超时了")
	} else {
		m.Range(func(key, value interface{}) bool {
			println(key.(int), value.(string))
			return true
		})
	}
	// 无状态，可重复执行
	m, err = p.All(2000)
	if err != nil && errors.Is(err, TimeoutError) {
		println("超时了")
	} else {
		m.Range(func(key, value interface{}) bool {
			println(key.(int), value.(string))
			return true
		})
	}
	println(runtime.NumGoroutine())
	time.Sleep(time.Second * 5)
	println(runtime.NumGoroutine())
}

func TestPromose_One(t *testing.T) {
	println(runtime.NumGoroutine())
	p := NewPromise(func() interface{} {
		time.Sleep(time.Second * 1)
		return "hello"
	}, func() interface{} {
		time.Sleep(time.Second * 1)
		return "hi"
	})
	m, err := p.One(2000)
	if err != nil && errors.Is(err, TimeoutError) {
		println("超时了")
	} else {
		m.Range(func(key, value interface{}) bool {
			println(key.(int), value.(string))
			return true
		})
	}
	println(runtime.NumGoroutine())
	time.Sleep(time.Second * 5)
	println(runtime.NumGoroutine())
}
