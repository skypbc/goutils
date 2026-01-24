package gcontext

import (
	"context"
	"github.com/skypbc/goutils/gcontext/ctxutils"
	"sync"
)

func WithParams(parent context.Context) (context.Context, map[string]any, *sync.RWMutex) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	params := map[string]any{}
	ctx := &paramsContext{parent, params, &sync.RWMutex{}}
	return ctx, params, ctx.paramsLock
}

func ExtractParamsContext(ctx context.Context) ctxutils.IParamsContext {
	if cmd_ctx, ok := ctx.Value(ctxutils.PARAMS_CONTEXT).(ctxutils.IParamsContext); ok {
		return cmd_ctx
	}
	return nil
}

type paramsContext struct {
	context.Context

	params     map[string]any
	paramsLock *sync.RWMutex
}

func (c *paramsContext) Value(key any) any {
	if key == ctxutils.PARAMS_CONTEXT {
		return ctxutils.IParamsContext(c)
	}
	if v, ok := key.(string); ok {
		if val, ok := c.Param(v); ok {
			return val
		}
	}
	return c.Context.Value(key)
}

func (c *paramsContext) Param(key string) (val any, ok bool) {
	c.paramsLock.RLock()
	defer c.paramsLock.RUnlock()

	val, ok = c.params[key]
	return val, ok
}

func (c *paramsContext) SetParam(key string, val any) {
	c.paramsLock.Lock()
	defer c.paramsLock.Unlock()
	c.params[key] = val
}

func (c *paramsContext) Locker() *sync.RWMutex {
	return c.paramsLock
}
