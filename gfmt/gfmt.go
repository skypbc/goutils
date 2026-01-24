package gfmt

import (
	"fmt"
	"github.com/skypbc/goutils/gfiles"
	"github.com/skypbc/goutils/gfmt/errors"
	"github.com/skypbc/goutils/gfmt/internal"
	"github.com/skypbc/goutils/gfmt/settings"
	"os"
	"path/filepath"
	"strings"
)

func Print(a ...any) (n int, err error) {
	if settings.Print.Active {
		s := fmt.Sprint(a...)
		if settings.Print.Debug.Active {
			s = addDebugInfo("Print", s)
		}
		res, err := fmt.Fprint(os.Stdout, s)
		if err != nil {
			return 0, errors.NewPrintError(err)
		}
		return res, nil
	}
	return 0, nil
}

func Printf(format string, a ...any) (n int, err error) {
	if settings.Print.Active {
		s := fmt.Sprintf(format, a...)
		if settings.Print.Debug.Active {
			s = addDebugInfo("Printf", s)
		}
		res, err := fmt.Fprint(os.Stdout, s)
		if err != nil {
			return 0, errors.NewPrintfError(err)
		}
		return res, nil
	}
	return 0, nil
}

func Printfln(format string, a ...any) (n int, err error) {
	if settings.Print.Active {
		s := fmt.Sprintf(format+"\n", a...)
		if settings.Print.Debug.Active {
			s = addDebugInfo("Printfln", s)
		}
		res, err := fmt.Fprint(os.Stdout, s)
		if err != nil {
			return 0, errors.NewPrintflnError(err)
		}
		return res, nil
	}
	return 0, nil
}

func Println(a ...any) (n int, err error) {
	if settings.Print.Active {
		s := fmt.Sprintln(a...)
		if settings.Print.Debug.Active {
			s = addDebugInfo("Println", s)
		}
		res, err := fmt.Fprint(os.Stdout, s)
		if err != nil {
			return 0, errors.NewPrintlnError(err)
		}
		return res, nil
	}
	return 0, nil
}

func Sprint(a ...any) string {
	return fmt.Sprint(a...)
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func Sprintfln(format string, a ...any) string {
	return fmt.Sprintf(format+"\n", a...)
}

func Sprintln(a ...any) string {
	return fmt.Sprintln(a...)
}

func addDebugInfo(name string, s string) string {
	if settings.Print.Debug.Active {
		ci := internal.GetCallerInfo(2)
		file := settings.Print.Debug.File
		pathDepth := settings.Print.Debug.PathDepth
		line := settings.Print.Debug.Line
		funcName := settings.Print.Debug.FuncName

		if !file && !line && !funcName && settings.Print.Debug.Name {
			s = fmt.Sprintf("File: %s, Func: %s, Line: %d, %s: %s", ci.File, ci.FuncName, ci.Line, name, s)

		} else {
			var parts []string
			if file {
				switch pathDepth {
				case -1:
					parts = append(parts, fmt.Sprintf("File: %s", ci.File))
				case 0:
					parts = append(parts, fmt.Sprintf("File: %s", gfiles.FullName(ci.File)))
				default:
					list := strings.Split(ci.File, settings.Print.Debug.PathSep)
					if len(list) > pathDepth+1 {
						list = list[len(list)-(pathDepth+1):]
					}
					parts = append(parts, fmt.Sprintf("File: %s", filepath.Join(list...)))
				}
			}
			if line {
				parts = append(parts, fmt.Sprintf("Line: %d", ci.Line))
			}
			if funcName {
				parts = append(parts, fmt.Sprintf("Func: %s", ci.FuncName))
			}
			if settings.Print.Debug.Name {
				parts = append(parts, fmt.Sprintf("%s%s %s", name, settings.Print.Debug.Sep, s))
				if settings.Print.Debug.Format {
					s = strings.Join(parts, "\n")
				} else {
					s = strings.Join(parts, ", ")
				}
			} else {
				if settings.Print.Debug.Format {
					s = fmt.Sprintf("%s\n%s", strings.Join(parts, "\n"), s)
				} else {
					s = fmt.Sprintf("%s%s %s", strings.Join(parts, ", "), settings.Print.Debug.Sep, s)
				}
			}
		}

		if settings.Print.Debug.Format {
			preNewLine := false
			postNewLine := false

			if settings.Print.Debug.PreNewLine == nil && settings.Print.Debug.PostNewLine == nil {
				preNewLine = false
				postNewLine = true
			} else {
				if settings.Print.Debug.PreNewLine != nil {
					preNewLine = *settings.Print.Debug.PreNewLine
				}
				if settings.Print.Debug.PostNewLine != nil {
					postNewLine = *settings.Print.Debug.PostNewLine
				}
			}

			if preNewLine && postNewLine {
				s = "\n" + s + "\n"
			} else if preNewLine {
				s = "\n" + s
			} else if postNewLine {
				s = s + "\n"
			}
		}
	}

	return s
}

func Scanln(a ...any) (n int, err error) {
	return fmt.Scanln(a...)
}
