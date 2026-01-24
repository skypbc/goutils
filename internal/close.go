package internal

import "io"

// Close закрывает ресурс со скрытием ошибки, если она есть.
func Close(closer io.Closer) {
	if closer == nil {
		return
	}
	_ = closer.Close()
}
