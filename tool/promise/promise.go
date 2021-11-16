package promise

import (
	"fmt"
	"sync"
	"time"
)

type FuncPromise func() interface{}

var TimeoutError = fmt.Errorf("执行超时")

// Promise 是一个无状态的对象，可以重复执行其中的 All 和 One 方法
type Promise struct {
	funcs []FuncPromise
}

func NewPromise(funcs ...FuncPromise) *Promise {
	return &Promise{funcs: funcs}
}

// TimeoutRun 执行 FuncPromise ,若超时将返回 TimeoutError ,正常则返回 FuncPromise 的执行结果.
// timeout 阻塞超时时间,单位毫秒.
func TimeoutRun(funcPromise FuncPromise, timeout time.Duration) (interface{}, error) {

	done := make(chan interface{}, 1)
	go func(done chan interface{}) {
		ret := funcPromise()
		done <- ret
	}(done)

	select {
	case r := <-done:
		close(done)
		return r, nil
	case <-time.After(time.Millisecond * timeout):
		return nil, TimeoutError
	}

}

// All 并行执行所有 FuncPromise,阻塞至全部完成后返回.
// timeout 阻塞超时时间,单位毫秒.
// 正常则返回一个 *sync.Map ,其中key为 NewPromise 时每个 FuncPromise 的索引,值为 FuncPromise 执行完毕后的返回结果
// 超时则将返回错误 TimeoutError.
func (p *Promise) All(timeout time.Duration) (*sync.Map, error) {
	r, err := TimeoutRun(func() interface{} {
		var wg sync.WaitGroup
		sm := &sync.Map{}
		for i, funcPromise := range p.funcs {
			wg.Add(1)
			funcPromise := funcPromise
			go func(i int) {
				defer wg.Done()
				ret := funcPromise()
				sm.Store(i, ret)
			}(i)
		}
		wg.Wait()
		return sm
	}, timeout)

	if err != nil {
		return nil, err
	}
	return r.(*sync.Map), err
}

// One 并行执行所有 FuncPromise,阻塞至任意一个完成后立即返回.
// timeout 阻塞超时时间,单位毫秒.
// 正常则返回一个 *sync.Map ,其中key为 NewPromise 时每个 FuncPromise 的索引,值为 FuncPromise 执行完毕后的返回结果
// 超时则将返回错误 TimeoutError.
func (p Promise) One(timeout time.Duration) (*sync.Map, error) {
	r, err := TimeoutRun(func() interface{} {
		// 该channel不能手动关闭，手动关闭将导致未完成的goroutine向channel发送数据时panic。
		// channel不需要通过close释放资源，只要没有goroutine持有channel，相关资源会自动释放。
		inCh := make(chan [2]interface{}, len(p.funcs))
		sm := &sync.Map{}
		for i, funcPromise := range p.funcs {
			funcPromise := funcPromise
			go func(i int) {
				ret := funcPromise()
				inCh <- [2]interface{}{i, ret}
			}(i)
		}
		select {
		case a := <-inCh:
			sm.Store(a[0], a[1])
			return sm
		}
	}, timeout)

	if err != nil {
		return nil, err
	}
	return r.(*sync.Map), err
}
