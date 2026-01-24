package errors

import (
	"github.com/skypbc/goutils/gerrors"
)

var (
	CodeFmtError         = gerrors.NewCode("FFFFFFFF-FFFF-2000-0000-000000000000")
	CodeFmtPrintError    = gerrors.NewCode("FFFFFFFF-FFFF-2000-0000-000000000001")
	CodeFmtPrintfError   = gerrors.NewCode("FFFFFFFF-FFFF-2000-0000-000000000002")
	CodeFmtPrintflnError = gerrors.NewCode("FFFFFFFF-FFFF-2000-0000-000000000003")
	CodeFmtPrintlnError  = gerrors.NewCode("FFFFFFFF-FFFF-2000-0000-000000000004")
)

func NewError(err ...error) gerrors.IError {
	return gerrors.NewError(gerrors.NewErrorArgs{
		Code:       CodeFmtError,
		Categories: []gerrors.Code{CodeFmtError},
		Name:       "GFMT",
		Template:   "A gfmt error has occurred",
		Parents:    err,
	})
}

func NewPrintError(err ...error) gerrors.IError {
	return gerrors.NewError(gerrors.NewErrorArgs{
		Code:       CodeFmtPrintError,
		Categories: []gerrors.Code{CodeFmtError},
		Name:       "GFMT.Print",
		Template:   "A gfmt print error has occurred",
		Parents:    err,
	})
}

func NewPrintfError(err ...error) gerrors.IError {
	return gerrors.NewError(gerrors.NewErrorArgs{
		Code:       CodeFmtPrintfError,
		Categories: []gerrors.Code{CodeFmtError},
		Name:       "GFMT.Printf",
		Template:   "A gfmt printf error has occurred",
		Parents:    err,
	})
}

func NewPrintflnError(err ...error) gerrors.IError {
	return gerrors.NewError(gerrors.NewErrorArgs{
		Code:       CodeFmtPrintflnError,
		Categories: []gerrors.Code{CodeFmtError},
		Name:       "GFMT.Printfln",
		Template:   "A gfmt printfln error has occurred",
		Parents:    err,
	})
}

func NewPrintlnError(err ...error) gerrors.IError {
	return gerrors.NewError(gerrors.NewErrorArgs{
		Code:       CodeFmtPrintlnError,
		Categories: []gerrors.Code{CodeFmtError},
		Name:       "GFMT.Println",
		Template:   "A gfmt println error has occurred",
		Parents:    err,
	})
}
