package greader

import (
	"io"
)

// ReadFull читает данные из reader в buff, пока не заполнит его полностью или не получит любую ошибку.
func ReadFull(reader io.Reader, buff []byte) (readed int, err error) {
	for len(buff) > 0 {
		n, err := reader.Read(buff[readed:])
		buff = buff[n:]
		readed += n
		if err != nil {
			return readed, err
		}
	}
	return readed, io.ErrShortBuffer
}
