package gpause

import (
	"bufio"
	"fmt"
	"os"
)

func PressEnterToContinue(msg ...string) {
	// Сообщение по умолчанию
	text := `Press "Enter" to continue...`
	if len(msg) > 0 {
		text = msg[0]
	}
	// Печатаем и гарантируем, что буфер stdout ушёл в консоль
	fmt.Print(text)
	_ = os.Stdout.Sync()

	// Читаем до перевода строки. Работает на Windows/Linux/macOS.
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
