package gerrors

import (
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
	"github.com/valyala/fasttemplate"
)

var InternalError IError
var DefaultErrorMessage string

func init() {
	InternalError = NewInternalError().
		SetPublic(true).
		ClearStackTrace()
	DefaultErrorMessage, _ = InternalError.Message()
}

const (
	TypeNone   ArgType = 0
	TypeNumber ArgType = 1
	TypeFloat  ArgType = 2
	TypeString ArgType = 3
	TypeBool   ArgType = 4
	TypeError  ArgType = 5
	TypeAny    ArgType = 6
	TypeUuid   ArgType = 7
)

type ArgType int

type Arg struct {
	Key   string  `json:"key"`
	Value any     `json:"value"`
	Type  ArgType `json:"type"`
}

type StackTraceItem struct {
	FuncName string `json:"func"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

func (e *StackTraceItem) String() string {
	return fmt.Sprintf("%s:%d, %s", e.File, e.Line, e.FuncName)
}

type Code uuid.UUID

func (c Code) Error() string {
	return "error code"
}

func (c Code) String() string {
	return uuid.UUID(c).String()
}

func (c Code) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, uuid.UUID(c).String())), nil
}

type NewErrorArgs struct {
	Code       Code
	Categories []Code
	Name       string

	Template string
	Args     map[string]any

	Public  bool
	Parents []error

	StackTraceSkip  *int
	StackTraceLimit *int
}

type IEmptyError interface {
	error

	Set(err IError) bool
	Get() IError

	Code() Code
	Unwrap() error
}

type emptyError struct {
	code Code
	err  IError
}

func (e *emptyError) Set(err IError) bool {
	if e.code != err.Code() {
		return false
	}
	e.err = err
	return true
}

func (e *emptyError) Get() IError {
	return e.err
}

func (e *emptyError) Error() string {
	if e.err == nil {
		return "<empty error>"
	}
	return e.err.Error()
}

func (e *emptyError) String() string {
	return e.Error()
}

func (e *emptyError) Code() Code {
	return e.code
}

func (e *emptyError) Unwrap() error {
	return e.err
}

// NewEmptyError используется для извлечения ошибки по коду с помощью errors.As
func NewEmptyError(code Code) IEmptyError {
	return &emptyError{
		code: code,
	}
}

func NewError(args NewErrorArgs) IError {
	if args.StackTraceSkip == nil {
		// По умолчанию предполагаем, что вызов NewError обернут в другую функцию, которая инициализирует ошибку,
		// поэтому пропускаем два уровня стека (внутрений NewError + внешний NewError)
		skip := 2
		args.StackTraceSkip = &skip
	}

	if args.StackTraceLimit == nil {
		// По умолчанию получаем полный стек вызовов
		limit := -1
		args.StackTraceLimit = &limit
	}

	parents := make([]error, 0, len(args.Parents))
	for _, p := range args.Parents {
		if p != nil {
			parents = append(parents, p)
		}
	}

	e := &BaseError{
		code:       args.Code,
		categories: args.Categories,
		name:       args.Name,
		template:   args.Template,
		public:     args.Public,
		parents:    parents,
	}

	if len(args.Args) > 0 {
		e.SetArgs(args.Args)
	}

	if *args.StackTraceLimit != 0 {
		if len(args.Parents) > 0 {
			if IsStackTrace(args.Parents...) {
				// Если среди родителей есть стек вызовов, то ограничиваем стек только первым уровнем
				// чтобы не дублировать информацию
				e.stackTrace = GetStackTrace(*args.StackTraceSkip, 1)
			} else {
				// Иначе получаем полный стек вызовов
				e.stackTrace = GetStackTrace(*args.StackTraceSkip, *args.StackTraceLimit)
			}
		} else {
			// Если нет родительских ошибок, то получаем полный стек вызовов
			e.stackTrace = GetStackTrace(*args.StackTraceSkip, *args.StackTraceLimit)
		}
	}

	return e
}

type IError interface {
	error
	fmt.Stringer

	Code() Code
	Categories() []Code

	Name() string

	HasCode(code Code) bool

	Is(target error) bool
	// Возвращает родительские ошибки
	Unwrap() []error

	Template() string
	SetTemplate(template string, args ...map[string]any) IError

	Args() map[string]Arg
	SetArgs(args map[string]any) IError

	Message() (string, error)
	StackTrace() []StackTraceItem
	ClearStackTrace() IError

	Public() bool
	SetPublic(public ...bool) IError

	ToMap(public bool, includeParents ...bool) map[string]any

	AddKey(key string, val any) IError
	AddUuid(key string, val uuid.UUID) IError
	AddStr(key string, val string) IError

	AddInt(key string, val int) IError
	AddInt64(key string, val int64) IError
	AddInt32(key string, val int32) IError
	AddInt16(key string, val int16) IError
	AddInt8(key string, val int8) IError

	AddUint(key string, val uint) IError
	AddUint64(key string, val uint64) IError
	AddUint32(key string, val uint32) IError
	AddUint16(key string, val uint16) IError
	AddUint8(key string, val uint8) IError

	AddFloat(key string, val float64) IError
	AddFloat32(key string, val float32) IError

	AddBool(key string, val bool) IError
	AddErr(key string, val error) IError
	AddAny(key string, val any) IError
}

type BaseError struct {
	// Уникальный идентификатор ошибки
	code Code
	// Категория к которой относится ошибка
	categories []Code
	// Имя ошибки
	name string

	// Шаблон сообщения об ошибке
	template string
	// Аргументы для шаблона сообщения об ошибке
	args map[string]Arg

	msg string

	// Стек вызовов, в котором была создана ошибка
	stackTrace []StackTraceItem

	// Является ли ошибка публичной (т.е. может передаваться за пределы приложения)
	public bool

	// Родительские ошибки
	parents []error
}

func (e *BaseError) Code() Code {
	return e.code
}

func (e *BaseError) Categories() []Code {
	return e.categories
}

func (e *BaseError) HasCode(code Code) bool {
	if e.code == code {
		return true
	}
	for _, c := range e.categories {
		if c == code {
			return true
		}
	}
	return false
}

func (e *BaseError) Name() string {
	return e.name
}

func (e *BaseError) Parents() []error {
	return e.parents
}

func (e *BaseError) Args() map[string]Arg {
	return e.args
}

func (e *BaseError) SetArgs(args map[string]any) IError {
	e.msg = ""
	for k, v := range args {
		e.setArgValue(k, v)
	}
	return e
}

func (e *BaseError) Template() string {
	return e.template
}

func (e *BaseError) SetTemplate(template string, args ...map[string]any) IError {
	e.template = template
	e.msg = ""
	if len(args) > 0 {
		e.SetArgs(args[0])
	}
	return e
}

func (e *BaseError) Message() (string, error) {
	if e.template == "" {
		return "", nil
	}
	if e.msg == "" {
		tmpl, err := fasttemplate.NewTemplate(e.template, "{", "}")
		if err != nil {
			return DefaultErrorMessage, err
		}
		if e.msg, err = tmpl.ExecuteFuncStringWithErr(func(w io.Writer, tag string) (int, error) {
			if strings.HasPrefix(tag, "{") {
				return w.Write([]byte(fmt.Sprintf("{%s}", tag)))
			}
			if item, ok := e.args[tag]; ok {
				switch item.Type {
				case TypeNumber:
					return w.Write([]byte(fmt.Sprintf("%d", item.Value)))
				case TypeFloat:
					return w.Write([]byte(fmt.Sprintf("%f", item.Value)))
				case TypeString:
					return w.Write([]byte(fmt.Sprintf("%s", item.Value)))
				case TypeBool:
					return w.Write([]byte(fmt.Sprintf("%t", item.Value)))
				case TypeError:
					if item.Value == nil {
						return w.Write([]byte("<nil>"))
					}
					if err, ok := item.Value.(error); ok {
						return w.Write([]byte(err.Error()))
					}
					return w.Write([]byte(fmt.Sprintf("%v", item.Value)))
				case TypeUuid:
					if item.Value == nil {
						return w.Write([]byte("<nil>"))
					}
					if u, ok := item.Value.(uuid.UUID); ok {
						return w.Write([]byte(u.String()))
					}
					return w.Write([]byte(fmt.Sprintf("%v", item.Value)))

				default:
					if item.Value == nil {
						return w.Write([]byte("<nil>"))
					}
					if r, ok := item.Value.(fmt.Stringer); ok {
						return w.Write([]byte(r.String()))
					}
					if r, ok := item.Value.(error); ok {
						return w.Write([]byte(r.Error()))
					}
					return w.Write([]byte(fmt.Sprintf("%v", item.Value)))
				}
			}
			return w.Write([]byte(fmt.Sprintf("{%s}", tag)))
		}); err != nil {
			return DefaultErrorMessage, err
		}
	}
	return e.msg, nil
}

func (e *BaseError) StackTrace() []StackTraceItem {
	return e.stackTrace
}

func (e *BaseError) ClearStackTrace() IError {
	e.stackTrace = nil
	return e
}

func (e *BaseError) Public() bool {
	return e.public
}

func (e *BaseError) SetPublic(public ...bool) IError {
	if len(public) > 0 {
		e.public = public[0]
	} else {
		e.public = true
	}
	return e
}

func (e *BaseError) AddKey(key string, val any) IError {
	return e.setArg(Arg{Type: TypeAny, Value: val, Key: key})
}

func (e *BaseError) AddStr(key string, val string) IError {
	return e.setArg(Arg{Type: TypeString, Value: val, Key: key})
}

func (e *BaseError) AddInt(key string, val int) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: int64(val), Key: key})
}

func (e *BaseError) AddInt64(key string, val int64) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: val, Key: key})
}

func (e *BaseError) AddInt32(key string, val int32) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: int64(val), Key: key})
}

func (e *BaseError) AddInt16(key string, val int16) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: int64(val), Key: key})
}

func (e *BaseError) AddInt8(key string, val int8) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: int64(val), Key: key})
}

func (e *BaseError) AddUint(key string, val uint) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: int64(val), Key: key})
}

func (e *BaseError) AddUint64(key string, val uint64) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: int64(val), Key: key})
}

func (e *BaseError) AddUint32(key string, val uint32) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: int64(val), Key: key})
}

func (e *BaseError) AddUint16(key string, val uint16) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: int64(val), Key: key})
}

func (e *BaseError) AddUint8(key string, val uint8) IError {
	return e.setArg(Arg{Type: TypeNumber, Value: int64(val), Key: key})
}

func (e *BaseError) AddFloat(key string, val float64) IError {
	return e.setArg(Arg{Type: TypeFloat, Value: val, Key: key})
}

func (e *BaseError) AddFloat32(key string, val float32) IError {
	return e.setArg(Arg{Type: TypeFloat, Value: float64(val), Key: key})
}

func (e *BaseError) AddBool(key string, val bool) IError {
	return e.setArg(Arg{Type: TypeBool, Value: val, Key: key})
}

func (e *BaseError) AddErr(key string, val error) IError {
	return e.setArg(Arg{Type: TypeError, Value: val, Key: key})
}

func (e *BaseError) AddAny(key string, val any) IError {
	return e.setArg(Arg{Type: TypeAny, Value: val, Key: key})
}

func (e *BaseError) AddUuid(key string, val uuid.UUID) IError {
	return e.setArg(Arg{Type: TypeUuid, Value: val, Key: key})
}

func (e *BaseError) setArgValue(key string, value any) IError {
	var typ ArgType
	switch value.(type) {
	case string:
		typ = TypeString
	case int, int8, int16, int32, int64:
		typ = TypeNumber
	case uint, uint8, uint16, uint32, uint64:
		typ = TypeNumber
	case float32, float64:
		typ = TypeFloat
	case bool:
		typ = TypeBool
	case error:
		typ = TypeError
	case uuid.UUID:
		typ = TypeUuid
	default:
		typ = TypeAny
	}
	return e.setArg(Arg{Key: key, Value: value, Type: typ})
}

func (e *BaseError) setArg(item Arg) IError {
	if e.args == nil {
		e.args = make(map[string]Arg)
	}
	e.args[item.Key] = item
	e.msg = ""
	return e
}

func (e *BaseError) Error() string {
	var sb strings.Builder

	if len(e.stackTrace) > 0 {
		for i, item := range e.stackTrace[:len(e.stackTrace)] {
			if i == 0 {
				sb.WriteString(item.String())
				sb.WriteString(" => ")
				e.writeMsg(&sb)
				continue
			}
			sb.Write([]byte(strings.Repeat(" ", 4*i)))
			sb.WriteString(item.String())
			sb.WriteString("\n")
		}
	} else {
		e.writeMsg(&sb)
	}

	if len(e.parents) > 0 {
		for i := len(e.parents) - 1; i >= 0; i-- {
			sb.WriteString(e.parents[i].Error())
		}
	}

	return sb.String()
}

func (e *BaseError) writeMsg(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%sError", e.name))
	sb.WriteString(": ")

	msg, _ := e.Message()
	if len(msg) == 0 {
		sb.WriteString("<no message>")
	} else {
		sb.WriteString(msg)
	}

	sb.WriteString("\n")
}

func (e *BaseError) String() string {
	return e.Error()
}

func (e *BaseError) Is(target error) bool {
	if target == nil {
		return false
	}
	switch x := any(target).(type) {
	case IError:
		return e.HasCode(x.Code())
	case *Code:
		return e.HasCode(*x)
	case Code:
		return e.HasCode(x)
	default:
		return false
	}
}

func (e *BaseError) As(target any) bool {
	if target == nil {
		return false
	}

	switch x := target.(type) {
	case *IEmptyError:
		return (*x).Set(e)

	case *IError:
		if e.code != (*x).Code() {
			return false
		}
		*x = e
		return true

	case **BaseError:
		if e.code != (*x).code {
			return false
		}
		*x = e
		return true

	default:
		return false
	}
}

func (e *BaseError) Unwrap() []error {
	return e.parents
}

func (e *BaseError) ToMap(public bool, includeParents ...bool) (res map[string]any) {
	res = make(map[string]any)
	if e == nil {
		return res
	}

	var parents []IError
	if len(includeParents) > 0 && includeParents[0] {
		for _, pe := range e.parents {
			parents = append(parents, CollectErrors[IError](pe)...)
		}
	}

	lastErr := IError(e)
	if public {
		for !lastErr.Public() {
			lastErr = nil
			if len(parents) == 0 {
				break
			}
			lastErr = parents[0]
			parents = parents[1:]
		}
		if lastErr == nil {
			lastErr = InternalError
		}
		resParents := []any{}
		for _, pe := range parents {
			if !pe.Public() {
				continue
			}
			parentJson := pe.ToMap(public, false)
			resParents = append(resParents, parentJson)
		}
		res["parents"] = resParents

	} else {
		resParents := []any{}
		for _, pe := range parents {
			parentJson := pe.ToMap(public, false)
			resParents = append(resParents, parentJson)
		}
		res["parents"] = resParents
	}

	res["code"] = lastErr.Code().String()
	res["name"] = lastErr.Name()
	res["text"], _ = lastErr.Message()
	res["template"] = lastErr.Template()

	if args := lastErr.Args(); args == nil {
		res["args"] = map[string]any{}
	} else {
		res["args"] = args
	}

	res["public"] = lastErr.Public()

	if !lastErr.Public() {
		res["stack_trace"] = lastErr.StackTrace()
	}

	return res
}
