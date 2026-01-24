package gparams

import (
	"github.com/skypbc/goutils/gerrors"
)

// Параметры разделяются пробелами. Например:
// - send_file=true filename='hello.txt'
// Значение с пробелами должны оборачиваться в кавычки " или '. Если в значении нужно применить кавычку,
// которой оно было обернуто, тогда эту кавычку нужно удвоить. Например:
// - filename='helo ”with” quote.txt'
// - filename="hello ""with"" qoute.txt"
// Пример, аналогов, без удвоения кавычки:
// - filename="hello 'with' qutoe.txt"
// - filename='hello "with" qutoe.txt'
func ParseString(data string) (res map[string]string, err error) {
	if len(data) == 0 {
		return nil, nil
	}

	res = map[string]string{}
	keyFound, valFound := false, false
	startQuote := rune(0)
	key := []rune{}
	val := []rune{}
	text := []rune(data)
	size := len(text)

	i := 0
	for i < size {
		r := text[i]
		if valFound && keyFound {
			if r == ' ' {
				i++
				continue
			}
			startQuote, keyFound, valFound = rune(0), false, false
		}
		if !keyFound {
			if r != '=' {
				if r == ' ' || r == '\t' || r == '\n' {
					return nil, gerrors.NewParseError().
						SetTemplate(`Incorrect white space, last position "{pos}"`).
						AddInt("pos", i)
				}
				key = append(key, r)
				i++
				continue
			} else {
				if len(key) == 0 {
					return nil, gerrors.NewParseError().
						SetTemplate(`Key not found, last position "{pos}"`).
						AddInt("pos", i)
				}
				keyFound = true
				i++
				continue
			}
		}

		if len(val) == 0 {
			switch r {
			case '\'':
				startQuote = r
			case '"':
				startQuote = r
			}
		}

		if r == startQuote {
			if len(val) == 0 {
				i++
				continue
			}
			if i+1 == size {
				break
			}
			if text[i+1] == startQuote {
				val = append(val, r)
				i += 2
				continue
			} else {
				res[string(key)] = string(val)
				key = key[:0]
				val = val[:0]
				valFound = true
				i++
				continue
			}
		}
		if r == ' ' {
			if startQuote > 0 {
				res[string(key)] = string(val)
				i++
				continue
			}
			res[string(key)] = string(val)
			key = key[:0]
			val = val[:0]
			valFound = true
			i++
			continue
		}
		val = append(val, r)
		i++
	}

	if len(key) > 0 && len(val) > 0 {
		res[string(key)] = string(val)
	}

	return res, nil
}
