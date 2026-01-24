package gstring

import "unicode/utf8"

func TruncateUTF8(s string, maxBytes int) string {
	if len(s) <= maxBytes {
		return s
	}
	i := 0
	for i < len(s) && i < maxBytes {
		_, size := utf8.DecodeRuneInString(s[i:])
		if i+size > maxBytes {
			break
		}
		i += size
	}
	return s[:i]
}
