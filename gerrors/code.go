package gerrors

import (
	"errors"
)

var (
	CodeWrapped         = NewCode("FFFFFFFF-FFFF-0000-0000-000000000000")
	CodeInternal        = NewCode("FFFFFFFF-FFFF-0000-0000-000000000001")
	CodeIncorrectParams = NewCode("FFFFFFFF-FFFF-0000-0000-000000000002")
	CodeUnknown         = NewCode("FFFFFFFF-FFFF-0000-0000-000000000003")
	CodeNotFound        = NewCode("FFFFFFFF-FFFF-0000-0000-000000000004")
	CodeKeyNotFound     = NewCode("FFFFFFFF-FFFF-0000-0000-000000000005")
	CodeType            = NewCode("FFFFFFFF-FFFF-0000-0000-000000000006")
	CodeParse           = NewCode("FFFFFFFF-FFFF-0000-0000-000000000007")
	CodeShortBuffer     = NewCode("FFFFFFFF-FFFF-0000-0000-000000000008")
	CodeReflect         = NewCode("FFFFFFFF-FFFF-0000-0000-000000000009")
	CodeSlice           = NewCode("FFFFFFFF-FFFF-0000-0000-00000000000A")
	CodeIndexOutOfRange = NewCode("FFFFFFFF-FFFF-0000-0000-00000000000B")

	CodeFile         = NewCode("FFFFFFFF-FFFF-0000-0001-000000000000")
	CodeFileWrite    = NewCode("FFFFFFFF-FFFF-0000-0001-000000000001")
	CodeFileRead     = NewCode("FFFFFFFF-FFFF-0000-0001-000000000002")
	CodeFileNotFound = NewCode("FFFFFFFF-FFFF-0000-0001-000000000003")
	CodeFileExists   = NewCode("FFFFFFFF-FFFF-0000-0001-000000000004")
)

func Join(errs ...error) error {
	errs = filter(errs, func(i int) (include bool) {
		return errs[i] != nil
	})
	if len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return errors.Join(errs...)
}

func Wrap(err ...error) IError {
	err = filter(err, func(i int) (include bool) {
		return err[i] != nil
	})
	if len(err) == 0 {
		return nil
	}
	return NewError(NewErrorArgs{
		Code:       CodeWrapped,
		Categories: []Code{CodeWrapped},
		Name:       "Wrapped",
		Parents:    err,
	})
}

func WrapWithSkip(skip int, err ...error) IError {
	return NewError(NewErrorArgs{
		Code:           CodeWrapped,
		Categories:     []Code{CodeWrapped},
		Name:           "Wrapped",
		Parents:        err,
		StackTraceSkip: &skip,
	})
}

func NewInternalError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeInternal,
		Name:     "Internal",
		Template: "An internal error has occurred",
		Parents:  err,
	})
}

func NewIncorrectParamsError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeIncorrectParams,
		Name:     "IncorrectParams",
		Template: "Incorrect parameters provided",
		Parents:  err,
	})
}

func NewUnknownError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeUnknown,
		Name:     "Unknown",
		Template: "An unknown error has occurred",
		Parents:  err,
	})
}

func NewNotFoundError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeNotFound,
		Name:     "NotFound",
		Template: "Not found",
		Parents:  err,
	})
}

func NewKeyNotFoundError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeKeyNotFound,
		Name:     "KeyNotFound",
		Template: "Key not found",
		Parents:  err,
	})
}

func NewTypeError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeType,
		Name:     "Type",
		Template: "Type error",
		Parents:  err,
	})
}

func NewParseError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeParse,
		Name:     "Parse",
		Template: "Parse error",
		Parents:  err,
	})
}

func NewShortBufferError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeShortBuffer,
		Name:     "ShortBuffer",
		Template: "Short buffer error",
		Parents:  err,
	})
}

func NewReflectError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeReflect,
		Name:     "Reflect",
		Template: "Reflect error",
		Parents:  err,
	})
}

func NewFileError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:       CodeFile,
		Categories: []Code{CodeFile},
		Name:       "FileError",
		Template:   "File error",
		Parents:    err,
	})
}

func NewFileWriteError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:       CodeFileWrite,
		Categories: []Code{CodeFile},
		Name:       "FileWrite",
		Template:   "File write error",
		Parents:    err,
	})
}

func NewFileReadError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:       CodeFileRead,
		Categories: []Code{CodeFile},
		Name:       "FileRead",
		Template:   "File read error",
		Parents:    err,
	})
}

func NewFileNotFoundError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:       CodeFileNotFound,
		Categories: []Code{CodeFile},
		Name:       "FileNotFound",
		Template:   "File not found",
		Parents:    err,
	})
}

func NewFileExistsError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:       CodeFileExists,
		Categories: []Code{CodeFile},
		Name:       "FileExists",
		Template:   "File already exists",
		Parents:    err,
	})
}

func NewSliceError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeSlice,
		Name:     "Slice",
		Template: "Slice error",
		Parents:  err,
	})
}

func NewIndexOutOfRangeError(err ...error) IError {
	return NewError(NewErrorArgs{
		Code:     CodeIndexOutOfRange,
		Name:     "IndexOutOfRange",
		Template: "Index out of range",
		Parents:  err,
	})
}
