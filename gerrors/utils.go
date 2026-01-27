package gerrors

import (
	"fmt"
	"runtime"

	"github.com/google/uuid"
)

func GetStackTrace(skip int, limit ...int) (res []StackTraceItem) {
	lim := 16
	if len(limit) > 0 {
		switch x := limit[0]; {
		case x == 0:
			return []StackTraceItem{}
		case x > 0:
			lim = x
		case x < 0:
			// pass
		}
	}
	pc := make([]uintptr, lim)
	n := runtime.Callers(2+skip, pc)
	frames := runtime.CallersFrames(pc[:n])

	res = make([]StackTraceItem, 0, n)
	for {
		frame, more := frames.Next()
		funcName := frame.Function
		file := frame.File
		line := frame.Line

		res = append(res, StackTraceItem{
			FuncName: funcName,
			File:     file,
			Line:     line,
		})

		if !more {
			break
		}
	}

	return res
}

// CollectErrors возвращает все ошибки типа T из дерева err (включая корневую).
func CollectErrors[T error](err error) []T {
	if err == nil {
		return nil
	}

	var (
		result  []T
		visited = make(map[error]struct{}, 16) // защита от циклов
	)

	var visit func(error)
	visit = func(e error) {
		if e == nil {
			return
		}
		// Защита от повторного обхода одного и того же объекта ошибки
		if _, ok := visited[e]; ok {
			return
		}
		visited[e] = struct{}{}

		// Проверяем только текущую ошибку
		if t, ok := any(e).(T); ok {
			result = append(result, t)
		}

		// Разворачиваем join-подобные ошибки
		if u, ok := e.(interface{ Unwrap() []error }); ok {
			for _, ue := range u.Unwrap() {
				visit(ue)
			}
			return
		}

		// Поддержка обычных обёрток
		if u, ok := e.(interface{ Unwrap() error }); ok {
			visit(u.Unwrap())
			return
		}
	}

	visit(err)
	return result
}

func CollectErrorsWithFunc[T error](err error, fn func(T) bool) []T {
	if err == nil {
		return nil
	}

	var (
		result  []T
		visited = make(map[error]struct{}, 16) // защита от циклов
	)

	var visit func(error)
	visit = func(e error) {
		if e == nil {
			return
		}
		// Защита от повторного обхода одного и того же объекта ошибки
		if _, ok := visited[e]; ok {
			return
		}
		visited[e] = struct{}{}

		// Проверяем только текущую ошибку
		if t, ok := any(e).(T); ok {
			if fn(t) {
				result = append(result, t)
			}
		}

		// Разворачиваем join-подобные ошибки
		if u, ok := e.(interface{ Unwrap() []error }); ok {
			for _, ue := range u.Unwrap() {
				visit(ue)
			}
			return
		}

		// Поддержка обычных обёрток
		if u, ok := e.(interface{ Unwrap() error }); ok {
			visit(u.Unwrap())
			return
		}
	}

	visit(err)
	return result
}

// FirstError возвращает первую ошибку типа T, найденную в дереве ошибок.
func FirstError[T error](err error) (T, bool) {
	var zero T
	if err == nil {
		return zero, false
	}

	visited := make(map[error]struct{}, 16)

	var visit func(error) (T, bool)
	visit = func(e error) (T, bool) {
		if e == nil {
			return zero, false
		}
		if _, ok := visited[e]; ok {
			return zero, false
		}
		visited[e] = struct{}{}

		// Проверяем ТОЛЬКО текущий узел
		if t, ok := any(e).(T); ok {
			return t, true
		}

		// Join-подобные ошибки
		if u, ok := e.(interface{ Unwrap() []error }); ok {
			for _, ue := range u.Unwrap() {
				if r, ok := visit(ue); ok {
					return r, true
				}
			}
			return zero, false
		}

		// Обычные wrapper-ошибки
		if u, ok := e.(interface{ Unwrap() error }); ok {
			return visit(u.Unwrap())
		}

		return zero, false
	}

	return visit(err)
}

func FirstErrorWithFunc[T error](err error, fn func(T) bool) (T, bool) {
	var zero T
	if err == nil {
		return zero, false
	}

	visited := make(map[error]struct{}, 16)

	var visit func(error) (T, bool)
	visit = func(e error) (T, bool) {
		if e == nil {
			return zero, false
		}
		if _, ok := visited[e]; ok {
			return zero, false
		}
		visited[e] = struct{}{}

		// Проверяем ТОЛЬКО текущий узел
		if t, ok := any(e).(T); ok {
			if fn(t) {
				return t, true
			}
		}

		// Join-подобные ошибки
		if u, ok := e.(interface{ Unwrap() []error }); ok {
			for _, ue := range u.Unwrap() {
				if r, ok := visit(ue); ok {
					return r, true
				}
			}
			return zero, false
		}

		// Обычные wrapper-ошибки
		if u, ok := e.(interface{ Unwrap() error }); ok {
			return visit(u.Unwrap())
		}

		return zero, false
	}

	return visit(err)
}

func LastError[T error](err error) (T, bool) {
	items := CollectErrors[T](err)
	if len(items) == 0 {
		var zero T
		return zero, false
	}
	return items[len(items)-1], true
}

func LastErrorWithFunc[T error](err error, fn func(T) bool) (T, bool) {
	items := CollectErrorsWithFunc(err, fn)
	if len(items) == 0 {
		var zero T
		return zero, false
	}
	return items[len(items)-1], true
}

func Contains(err error, code Code) bool {
	_, ok := FirstErrorWithFunc(err, func(i IError) bool {
		if i.Code() == code {
			return true
		}
		for _, c := range i.Categories() {
			if c == code {
				return true
			}
		}
		return false
	})
	return ok
}

func IsStackTrace(err ...error) bool {
	for _, e := range err {
		if e == nil {
			continue
		}
		if _, ok := FirstError[IError](e); ok {
			return true
		}
	}
	return false
}

func NewCode(v any) Code {
	c, err := TryNewCode(v)
	if err != nil {
		panic(fmt.Sprintf("gerrors3: NewCode: %v", err))
	}
	return c
}

// TryNewCode приводит входное значение к Code.
// Поддерживаемые типы:
//   - Code, *Code
//   - uuid.UUID, *uuid.UUID
//   - string, *string (UUID в текстовом виде)
//   - [16]byte, *[]byte (len=16), []byte (len=16)
//
// Если тип не поддержан или строка/байты некорректны — вернёт ошибку.
func TryNewCode(v any) (Code, error) {
	switch x := v.(type) {
	case nil:
		return Code(uuid.Nil), fmt.Errorf("code is nil")

	case Code:
		return x, nil
	case *Code:
		if x == nil {
			return Code(uuid.Nil), fmt.Errorf("code is nil")
		}
		return *x, nil

	case uuid.UUID:
		return Code(x), nil
	case *uuid.UUID:
		if x == nil {
			return Code(uuid.Nil), fmt.Errorf("uuid is nil")
		}
		return Code(*x), nil

	case string:
		u, err := uuid.Parse(x)
		if err != nil {
			return Code(uuid.Nil), fmt.Errorf("invalid uuid string: %w", err)
		}
		return Code(u), nil
	case *string:
		if x == nil {
			return Code(uuid.Nil), fmt.Errorf("string is nil")
		}
		u, err := uuid.Parse(*x)
		if err != nil {
			return Code(uuid.Nil), fmt.Errorf("invalid uuid string: %w", err)
		}
		return Code(u), nil

	case [16]byte:
		return Code(uuid.UUID(x)), nil

	case []byte:
		if len(x) != 16 {
			return Code(uuid.Nil), fmt.Errorf("invalid uuid bytes length: %d", len(x))
		}
		var b [16]byte
		copy(b[:], x)
		return Code(uuid.UUID(b)), nil

	case *[]byte:
		if x == nil {
			return Code(uuid.Nil), fmt.Errorf("bytes is nil")
		}
		if len(*x) != 16 {
			return Code(uuid.Nil), fmt.Errorf("invalid uuid bytes length: %d", len(*x))
		}
		var b [16]byte
		copy(b[:], *x)
		return Code(uuid.UUID(b)), nil

	default:
		return Code(uuid.Nil), fmt.Errorf("unsupported code type %T", v)
	}
}

func filter[S ~[]T, T any](s S, f func(i int) (include bool)) (res S) {
	ids := make([]int, 0, len(s))
	for i := range s {
		if f(i) {
			ids = append(ids, i)
		}
	}
	res = make([]T, 0, len(ids))
	for _, index := range ids {
		res = append(res, s[index])
	}
	return res
}
