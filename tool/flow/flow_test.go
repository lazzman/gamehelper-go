package flow

import (
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestControlMachine(t *testing.T) {
	fcm := NewControlMachine()
	fcm.
		Next("第一步", FuncConditionTrue, func() {
			println("第一步")
		}, FuncFlowDoNothing).
		NextDelay(5200).
		Next("第二步", FuncConditionTrue, func() {
			println("第二步")
		}, FuncFlowDoNothing).
		NextDelay(1000).
		Next("第三步", FuncConditionFalse, func() {
			println("第三步")
		}, func() {
			println("第三步出错了，但是继续执行")
		}).
		NextDelay(1000).
		Next("第四步", FuncConditionTrue, func() {
			println("第四步")
		}, FuncFlowDoNothing).
		NextDelay(1000).
		Next("第五步", FuncConditionFalse, func() {
			println("第五步")
		}, func() {
			//fcm.Stop()
			panic("意外出现了")
		}).
		OnStart(func() {
			println("流程控制器准备完毕")
		}).
		OnEnd(func() {
			println("流程控制器执行完毕")
		}).
		OnStop(func() {
			println("流程控制器停止执行")
		}).
		OnCycleStart(func() {
			println("一个周期执行前")
		}).
		OnCycleEnd(func() {
			println("一个周期执行后")
		}).
		OnFlowStart(func() {
			println("流程执行前:" + fcm.GetFlowName())
		}).
		OnFlowEnd(func() {
			println("流程执行后:" + fcm.GetFlowName())
		}).
		OnError(func() {
			println("发生错误了")
		}).
		Forever().
		Start()

	println("bye")
}

func TestGoroutineStop(t *testing.T) {
	fcm := NewControlMachine()
	fcm.
		Next("第一步", FuncConditionTrue, func() {
			println("第一步")
		}, FuncFlowDoNothing).
		NextDelay(500).
		Next("第二步", FuncConditionTrue, func() {
			println("第二步")
		}, FuncFlowDoNothing).
		NextDelay(1000).
		Next("第三步", FuncConditionFalse, func() {
			println("第三步")
		}, func() {
			println("第三步出错了，但是继续执行")
		}).
		NextDelay(1000).
		Next("第四步", FuncConditionTrue, func() {
			println("第四步")
		}, FuncFlowDoNothing).
		NextDelay(1000).
		Next("第五步", FuncConditionTrue, func() {
			println("第五步")
		}, FuncFlowDoNothing).
		OnStart(func() {
			println("流程控制器准备完毕")
		}).
		OnEnd(func() {
			println("流程控制器执行完毕")
		}).
		OnStop(func() {
			println("流程控制器停止执行")
		}).
		OnCycleStart(func() {
			println("一个周期执行前")
		}).
		OnCycleEnd(func() {
			println("一个周期执行后")
		}).
		OnFlowStart(func() {
			println("流程执行前:" + fcm.GetFlowName())
		}).
		OnFlowEnd(func() {
			println("流程执行后:" + fcm.GetFlowName())
		}).
		OnError(func() {
			println("发生错误了")
		}).
		Forever()

	go func() {
		fcm.Start()
	}()
	println("协程数:" + strconv.Itoa(runtime.NumGoroutine()))
	// 测试从另一个协程停止流程控制器
	time.AfterFunc(time.Second*1, func() {
		fcm.Stop()
	})
	println("协程数:" + strconv.Itoa(runtime.NumGoroutine()))
	time.Sleep(time.Second * 5)
	println("协程数:" + strconv.Itoa(runtime.NumGoroutine()))
}
