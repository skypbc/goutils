package gpause

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// PressAnyKeyToContinue печатает приглашение и ждёт нажатия ЛЮБОЙ клавиши (без Enter).
func PressAnyKeyToContinue(msg ...string) error {
	// Сообщение по умолчанию
	text := "Press any key to continue..."
	if len(msg) > 0 {
		text = msg[0]
	}
	fmt.Print(text)

	// Если stdin не терминал (перенаправление/pipe) — читаем 1 байт как есть.
	fd := int(os.Stdin.Fd())
	if !term.IsTerminal(fd) {
		var b [1]byte
		_, err := os.Stdin.Read(b[:])
		fmt.Println()
		return err
	}

	// Переводим терминал в raw-режим и обязательно восстанавливаем.
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer func() {
		_ = term.Restore(fd, oldState)
		fmt.Println()
	}()

	// Читаем один байт — этого достаточно для "любой клавиши".
	var b [1]byte
	_, err = os.Stdin.Read(b[:])
	return err
}
