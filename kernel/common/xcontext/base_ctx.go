// 定义公共上下文结构，明确定义上下文结构，方便代码阅读
package xcontext

import (
	"context"
	"time"

	"github.com/xuperchain/xupercore/lib/logs"
	"github.com/xuperchain/xupercore/lib/timer"
)

type XContext interface {
	context.Context
	GetLog() logs.Logger
	GetTimer() *timer.XTimer
}

// 考虑到后续对操作生命周期做控制，先实现空context.Context接口，方便扩展
// 同时定义扩展全局需要的公共成员，方便为各对象统一注入和管理
type BaseCtx struct {
	context.Context
	XLog  logs.Logger
	Timer *timer.XTimer
}

func WithNewContext(parent XContext, ctx context.Context) XContext {
	return &BaseCtx{
		Context: ctx,
		XLog:    parent.GetLog(),
		Timer:   parent.GetTimer(),
	}
}

func (t *BaseCtx) GetLog() logs.Logger {
	return t.XLog
}

func (t *BaseCtx) GetTimer() *timer.XTimer {
	return t.Timer
}

func (t *BaseCtx) Deadline() (deadline time.Time, ok bool) {
	if t.Context != nil {
		return t.Context.Deadline()
	}
	return
}

func (t *BaseCtx) Done() <-chan struct{} {
	if t.Context != nil {
		return t.Context.Done()
	}
	return nil
}

func (t *BaseCtx) Err() error {
	if t.Context != nil {
		t.Context.Err()
	}
	return nil
}

func (t *BaseCtx) Value(key interface{}) interface{} {
	if t.Context != nil {
		return t.Context.Value(key)
	}
	return nil
}
