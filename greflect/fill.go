package greflect

import (
	"fmt"
	"reflect"
)

// TryFill пытается присвоить значение src переменной dest с использованием рефлексии.
// Функция поддерживает:
//   - Прямое присваивание при совпадении типов.
//   - Конвертацию типов, если они совместимы (например, int → int64, float32 → float64, float64 -> int16).
//   - Преобразование срезов ([]T или []interface{}) в срезы другого типа,
//     включая автоматическое приведение чисел к целевому типу или преобразование в строки.
//
// Возвращает true, если присвоение прошло успешно, и false в случае несовместимости типов.
func TryFill(src any, dest any) (ok bool) {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dest)

	if dstVal.Kind() != reflect.Ptr {
		return false
	}

	dstElem := dstVal.Elem()
	if !dstElem.IsValid() {
		return false
	}

	// Если типы совпадают – просто присваиваем
	if srcVal.Type().AssignableTo(dstElem.Type()) {
		dstElem.Set(srcVal)
		return true
	}

	// Если конвертируемые типы – используем Convert
	if srcVal.Type().ConvertibleTo(dstElem.Type()) {
		dstElem.Set(srcVal.Convert(dstElem.Type()))
		return true
	}

	// Обработка случая slice → slice
	if srcVal.Kind() == reflect.Slice && dstElem.Kind() == reflect.Slice {
		elemType := dstElem.Type().Elem()
		newSlice := reflect.MakeSlice(dstElem.Type(), srcVal.Len(), srcVal.Len())

		for i := 0; i < srcVal.Len(); i++ {
			item := srcVal.Index(i)

			// Разыменовываем interface{}
			if item.Kind() == reflect.Interface && !item.IsNil() {
				item = item.Elem()
			}

			if item.Type().AssignableTo(elemType) {
				newSlice.Index(i).Set(item)
				continue
			}

			if item.Type().ConvertibleTo(elemType) {
				newSlice.Index(i).Set(item.Convert(elemType))
				continue
			}

			if isNumber(item.Kind()) && isNumber(elemType.Kind()) {
				converted := convertNumber(item, elemType.Kind())
				newSlice.Index(i).Set(converted)
				continue
			}

			if elemType.Kind() == reflect.String {
				newSlice.Index(i).Set(reflect.ValueOf(fmt.Sprintf("%v", item.Interface())))
				continue
			}

			return false
		}

		dstElem.Set(newSlice)
		return true
	}

	return false
}

func isNumber(k reflect.Kind) bool {
	return k >= reflect.Int && k <= reflect.Float64
}

func convertNumber(v reflect.Value, target reflect.Kind) reflect.Value {
	// Всегда сначала приводим к float64
	var f float64
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		f = v.Convert(reflect.TypeOf(float64(0))).Float()
	default:
		f = float64(v.Convert(reflect.TypeOf(int64(0))).Int())
	}

	switch target {
	case reflect.Int:
		return reflect.ValueOf(int(f))
	case reflect.Int8:
		return reflect.ValueOf(int8(f))
	case reflect.Int16:
		return reflect.ValueOf(int16(f))
	case reflect.Int32:
		return reflect.ValueOf(int32(f))
	case reflect.Int64:
		return reflect.ValueOf(int64(f))
	case reflect.Uint:
		return reflect.ValueOf(uint(f))
	case reflect.Uint8:
		return reflect.ValueOf(uint8(f))
	case reflect.Uint16:
		return reflect.ValueOf(uint16(f))
	case reflect.Uint32:
		return reflect.ValueOf(uint32(f))
	case reflect.Uint64:
		return reflect.ValueOf(uint64(f))
	case reflect.Uintptr:
		return reflect.ValueOf(uintptr(f))
	case reflect.Float32:
		return reflect.ValueOf(float32(f))
	case reflect.Float64:
		return reflect.ValueOf(f)
	default:
		return v
	}
}
