package flow

import (
	"errors"
	"fmt"
	"gamehelper/utils"
	"sync"
)

// 生命周期Hook方法
const (
	ONSTART      = "onStart"
	ONEND        = "onEnd"
	ONSTOP       = "onStop"
	ONERROR      = "onError"
	ONCYCLESTART = "onCycleStart"
	ONCYCLEEND   = "onCycleEnd"
	ONFLOWSTART  = "onFlowStart"
	ONFLOWEND    = "onFlowEnd"
)

var (
	FuncConditionTrue = func() bool {
		return true
	}
	FuncConditionFalse = func() bool {
		return false
	}
	FuncFlowDoNothing = func() {}
)

type FuncFlow func()

type FuncCondition func() bool

var StopError = fmt.Errorf("流程控制停止")

// FullFlow 一个完整的流程定义
type FullFlow struct {
	flowName      string        //流程名称
	conditionFunc FuncCondition //流程条件
	trueHandle    FuncFlow      //满足条件则调用当前函数
	falseHandle   FuncFlow      //不满足条件则调用当前哈数
}

// ControlMachine 流程控制器.
// 非线程安全对象.
type ControlMachine struct {
	flowName string              //当前流程名称
	i        int                 //当前索引
	iRefer   int                 //索引参照，正常情况下与i值相等，当发生流程跳转等动作后两值不等
	flowList []FullFlow          //流程定义集合
	forever  bool                //是否循环执行流程
	running  bool                //是否运行中
	hookDict map[string]FuncFlow //流程控制器生命周期hook
	lock     sync.Mutex          //防止并发启动
}

func NewControlMachine() *ControlMachine {
	c := &ControlMachine{
		flowName: "",
		i:        0,
		iRefer:   0,
		flowList: nil,
		forever:  false,
		running:  false,
	}

	defaultHookDict := make(map[string]FuncFlow, 8)
	defaultHookDict[ONSTART] = func() {}
	defaultHookDict[ONEND] = func() {}
	defaultHookDict[ONSTOP] = func() {}
	defaultHookDict[ONERROR] = func() {}
	defaultHookDict[ONCYCLESTART] = func() {}
	defaultHookDict[ONCYCLEEND] = func() {}
	defaultHookDict[ONFLOWSTART] = func() {
		c.PanicIfStop()
	}
	defaultHookDict[ONFLOWEND] = func() {
		c.PanicIfStop()
	}
	c.hookDict = defaultHookDict

	return c
}

// GetFlowName 获取当前流程名称
func (m *ControlMachine) GetFlowName() string {
	return m.flowName
}

// Next 流程定义.
// flowName 流程名称.不能为空且不能重复.
// conditionFunc 进入流程的条件判断函数.不能为nil.
// trueHandle 条件函数返回true则进入流程处理函数.不能为nil.
// falseHandle 条件函数返回false则进入到该处理函数中.不能为nil.
func (m *ControlMachine) Next(flowName string, conditionFunc FuncCondition, trueHandle FuncFlow, falseHandle FuncFlow) *ControlMachine {
	if flowName == "" {
		panic("流程定义 Next() flowName为空")
	}
	if conditionFunc == nil {
		panic("流程定义 Next() conditionFunc为nil")
	}
	if trueHandle == nil {
		panic("流程定义 Next() trueHandle为nil")
	}
	if falseHandle == nil {
		panic("流程定义 Next() falseHandle为nil")
	}
	m.flowList = append(m.flowList, FullFlow{
		flowName:      flowName,
		conditionFunc: conditionFunc,
		trueHandle:    trueHandle,
		falseHandle:   falseHandle,
	})
	return m
}

// NextDelay 定义阻塞流程.
// mills 阻塞毫秒.
func (m *ControlMachine) NextDelay(mills int) *ControlMachine {
	return m.Next("delay", FuncConditionTrue, func() {
		m.DelayMills(mills)
	}, FuncFlowDoNothing)
}

// Forever 流程不停循环执行
func (m *ControlMachine) Forever() *ControlMachine {
	m.forever = true
	return m
}

// OnStart 流程控制器运行前的处理.
func (m *ControlMachine) OnStart(handle FuncFlow) *ControlMachine {
	if handle != nil {
		m.hookDict[ONSTART] = handle
	}
	return m
}

// OnEnd 流程控制器运行完毕后的处理. Forever 的流程不会执行 OnEnd.
func (m *ControlMachine) OnEnd(handle FuncFlow) *ControlMachine {
	if handle != nil {
		m.hookDict[ONEND] = handle
	}
	return m
}

// OnStop 流程控制器终止运行后的处理.与 OnEnd 只会生效一个.
// 当调用 Stop 方法时会触发handle.
// Forever 的流程会执行 OnStop,当调用 Stop 后将触发 OnStop.
func (m *ControlMachine) OnStop(handle FuncFlow) *ControlMachine {
	if handle != nil {
		m.hookDict[ONSTOP] = handle
	}
	return m
}

// OnError 当发生 panic 时，执行当前handle.当前handle中无法执行 Back Jump Forward 方法，可以执行 ReStart 方法。
func (m *ControlMachine) OnError(handle FuncFlow) *ControlMachine {
	if handle != nil {
		m.hookDict[ONERROR] = handle
	}
	return m
}

// OnCycleStart 循环一个周期的开始前执行.
func (m *ControlMachine) OnCycleStart(handle FuncFlow) *ControlMachine {
	if handle != nil {
		m.hookDict[ONCYCLESTART] = handle
	}
	return m
}

// OnCycleEnd 循环一个周期的结束后执行.
func (m *ControlMachine) OnCycleEnd(handle FuncFlow) *ControlMachine {
	if handle != nil {
		m.hookDict[ONCYCLEEND] = handle
	}
	return m
}

// OnFlowStart 进入每个流程前执行.
func (m *ControlMachine) OnFlowStart(handle FuncFlow) *ControlMachine {
	if handle != nil {
		m.hookDict[ONFLOWSTART] = handle
	}
	return m
}

// OnFlowEnd 每个流程处理完毕后执行.
func (m *ControlMachine) OnFlowEnd(handle FuncFlow) *ControlMachine {
	if handle != nil {
		m.hookDict[ONFLOWEND] = handle
	}
	return m
}

// Back 流程后退.
// n 后退几步.
func (m *ControlMachine) Back(n int) {
	m.i = m.i - n
	m.iRefer = -1
}

// Forward 流程前进.
// n 前进几步.
func (m *ControlMachine) Forward(n int) {
	m.i = m.i + n
	m.iRefer = -1
}

// Jump 流程跳转.
// flowName 根据名称跳转到指定流程.
func (m *ControlMachine) Jump(flowName string) {
	for i, fullFlow := range m.flowList {
		if fullFlow.flowName == flowName {
			m.i = i
			m.iRefer = -1
			break
		}
	}
}

// Start 开始运行流程控制器
func (m *ControlMachine) Start() *ControlMachine {
	m.lock.Lock()
	defer m.lock.Unlock()
	defer func() {
		if p := recover(); p != nil {
			if err, ok := p.(error); ok {
				if errors.Is(err, StopError) {
					m.hookDict[ONSTOP]()
				} else {
					m.hookDict[ONERROR]()
				}
			} else {
				m.hookDict[ONERROR]()
			}
		}
	}()
	// 运行中则不会重复运行，避免执行错乱
	if m.running {
		return m
	}
	m.running = true
	m.hookDict[ONSTART]()
	if m.forever {
		for true {
			m.run()
		}
	} else {
		m.run()
	}
	m.hookDict[ONEND]()
	return m
}

func (m *ControlMachine) run() {
	m.hookDict[ONCYCLESTART]()
	m.i = 0
	for m.i < len(m.flowList) {
		m.iRefer = m.i
		flow := m.flowList[m.i]
		m.flowName = flow.flowName
		if !m.running {
			panic(StopError)
		}
		m.hookDict[ONFLOWSTART]()
		if flow.conditionFunc() {
			flow.trueHandle()
		} else {
			flow.falseHandle()
		}
		m.hookDict[ONFLOWEND]()
		if m.iRefer == m.i {
			// 在flow的handle中没有执行过back、next或jump等方法，则i递增
			m.i++
		}
	}
	m.hookDict[ONCYCLEEND]()
}

// ReStart 重新运行流程控制器.
func (m *ControlMachine) ReStart() {
	m.Start()
}

// Stop 控制器在进入下一个流程后停止运行.
// 协程无法被外部直接停止.
// 停止的判断条件设置在进入每个流程前，也就是在当前流程执行完毕后，进入下一个流程前判断 ControlMachine.running 字段状态决定是否停止.
func (m *ControlMachine) Stop() {
	m.running = false
}

// PanicIfStop 控制器运行状态检查，若为停止状态则 panic StopError.
func (m *ControlMachine) PanicIfStop() {
	if !m.running {
		panic(StopError)
	}
}

// DelayMills 流程控制器提供的延迟方法，内置了控制器运行状态检查
func (m *ControlMachine) DelayMills(mills int) {
	if mills <= 500 {
		m.PanicIfStop()
		utils.DelayMills(mills)
		m.PanicIfStop()
	} else {
		// 超过500ms的延迟方法分段延迟，更快响应控制器运行状态检查
		x := mills / 500
		y := mills % 500
		for i := 0; i < x; i++ {
			m.PanicIfStop()
			utils.DelayMills(500)
			m.PanicIfStop()
		}
		m.PanicIfStop()
		utils.DelayMills(y)
		m.PanicIfStop()
	}
}
