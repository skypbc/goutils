package gnet

import (
	"bufio"
	"errors"
	"io"
	"net"
	"time"
)

// SafeConn обертка над net.Conn, позволяет использовать Peek для неблокирующего чтения данных
type SafeConn struct {
	conn   net.Conn
	reader *bufio.Reader
}

// NewSafeConn создает обертку над соединением.
func NewSafeConn(c net.Conn) *SafeConn {
	return &SafeConn{
		conn:   c,
		reader: bufio.NewReader(c),
	}
}

// Read читает из буфера (reader), сохраняя корректный порядок данных.
func (s *SafeConn) Read(p []byte) (int, error) {
	return s.reader.Read(p)
}

// Peek читает данные из буфера без блокировки и без извлечения данных из потока.
func (s *SafeConn) Peek(n int) ([]byte, error) {
	return s.reader.Peek(n)
}

// Write просто прокидывает запись напрямую в net.Conn.
func (s *SafeConn) Write(p []byte) (int, error) {
	return s.conn.Write(p)
}

// Close закрывает соединение.
func (s *SafeConn) Close() error {
	return s.conn.Close()
}

// LocalAddr возвращает локальный адрес.
func (s *SafeConn) LocalAddr() net.Addr {
	return s.conn.LocalAddr()
}

// RemoteAddr возвращает удалённый адрес.
func (s *SafeConn) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}

// SetDeadline прокидывает установку дедлайна.
func (s *SafeConn) SetDeadline(t time.Time) error {
	return s.conn.SetDeadline(t)
}

// SetReadDeadline прокидывает установку дедлайна на чтение.
func (s *SafeConn) SetReadDeadline(t time.Time) error {
	return s.conn.SetReadDeadline(t)
}

// SetWriteDeadline прокидывает установку дедлайна на запись.
func (s *SafeConn) SetWriteDeadline(t time.Time) error {
	return s.conn.SetWriteDeadline(t)
}

// IsClosed выполняет неблокирующую проверку, закрыто ли соединение.
// Использует Peek(1), не нарушая поток данных.
// Функция меняет делайн на чтение, поэтому после её работы необходимо либо его сбросить, либо установить преждний.
func (s *SafeConn) IsClosed(timeout time.Duration) (bool, error) {
	_ = s.conn.SetReadDeadline(time.Now().Add(timeout))
	n, err := s.reader.Peek(1)

	if err != nil {
		if errors.Is(err, io.EOF) {
			if len(n) > 0 {
				return false, nil // Открыто, так как есть данные
			}
			return true, nil // Закрыто
		}
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return false, nil // Открыто, нет данных
		}
		return true, err // Другая ошибка
	}
	return false, nil // Открыто
}
