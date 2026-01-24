package gmap

import (
	"github.com/skypbc/goutils/gmap/internal"
)

func Set(m any, key string, value any) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetWithSep(m any, key string, value any, sep string) (nm any, err error) {
	return internal.Set(m, key, value, sep)
}

func SetInt(m any, key string, value int) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetInt8(m any, key string, value int8) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetInt16(m any, key string, value int16) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetInt32(m any, key string, value int32) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetInt64(m any, key string, value int64) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetUint(m any, key string, value uint) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetUint8(m any, key string, value uint8) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetUint16(m any, key string, value uint16) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetUint32(m any, key string, value uint32) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetUint64(m any, key string, value uint64) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetFloat32(m any, key string, value float32) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetFloat64(m any, key string, value float64) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}

func SetString(m any, key string, value string) (nm any, err error) {
	return internal.Set(m, key, value, ".")
}
