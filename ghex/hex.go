package ghex

import (
	"encoding/hex"
	"fmt"
	"github.com/skypbc/goutils/gerrors"
	"strings"
)

type FromOpts struct {
	Upper bool
	Lower bool
	Sep   string
}

func FromBytes(data []byte, opts ...FromOpts) (res string) {
	n := len(data)
	if n == 0 {
		return ""
	}

	upper := false
	lower := false
	sep := ""
	if len(opts) > 0 {
		o := opts[0]
		upper = o.Upper
		lower = o.Lower
		if len(o.Sep) > 0 {
			sep = o.Sep
		}
	}

	sepLen := len(sep)
	raw := make([]byte, n*2+(n-1)*sepLen)

	var hex string
	for i, b := range data {
		if i > 0 {
			copy(raw[i*2+(i-1)*sepLen:], sep)
		}
		if upper {
			hex = fmt.Sprintf("%02X", b)
		} else if lower {
			hex = fmt.Sprintf("%02x", b)
		} else {
			hex = fmt.Sprintf("%02X", b)
		}
		copy(raw[i*2+i*sepLen:], hex)
	}

	return string(raw)
}

func ToBytes(data string, sep ...string) (res []byte) {
	res, err := TryToBytes(data, sep...)
	if err != nil {
		panic(err)
	}
	return res
}

func TryToBytes(data string, sep ...string) (res []byte, err error) {
	if len(sep) > 0 {
		data = strings.ReplaceAll(data, sep[0], "")
	}
	if len(data)%2 != 0 {
		return nil, gerrors.NewIncorrectParamsError().
			SetTemplate(`hex string must be even, got length "{length}"`).
			AddInt("length", len(data))
	}
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return nil, gerrors.Wrap(err)
	}
	return bytes, nil
}
