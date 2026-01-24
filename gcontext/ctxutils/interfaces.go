package ctxutils

import (
	"context"
	"sync"
)

type IParamsContext interface {
	context.Context

	Param(string) (any, bool)
	SetParam(string, any)
	Locker() *sync.RWMutex
}
