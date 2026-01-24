package internal

import (
	"github.com/skypbc/goutils/gerrors"
	"github.com/skypbc/goutils/gnum"
	"github.com/skypbc/goutils/greflect"
	"reflect"
	"strings"
)

// Если после нормализации item, его значение изменится и item передавался по ссылке, то это измение будет
// отражено в оригинальном item (при условии, что его значение соотвествует any).
func Set(item any, key string, value any, sep string) (newItem any, err error) {
	root := greflect.LastElem(reflect.ValueOf(item))
	keys := strings.Split(key, sep)

	if root.Kind() == reflect.Map || root.Kind() == reflect.Slice {
		if new, ok := normalize(root); ok {
			if root.CanSet() {
				root.Set(new)
			} else {
				root = new
			}
		}
	}

	parent, last, lastGroupKey, ok := root, root, "", false

	if len(keys) > 1 {
		for _, key := range keys[:len(keys)-1] {
			parent = last
			lastGroupKey = key

			switch parent.Kind() {
			case reflect.Map:
				// Пробуем достать элемент используя заданную часть ключа
				last, err = getFromMap(parent, key)
				if err != nil || !last.IsValid() {
					// Если ключ не существует, создаем новый map-элемент
					last = reflect.ValueOf(map[string]any{})
					parent.SetMapIndex(
						reflect.ValueOf(key),
						last,
					)
				}
			case reflect.Array, reflect.Slice:
				last, err = getFromSlice(parent, key)
				if err != nil || !last.IsValid() {
					return nil, gerrors.NewParseError().
						SetTemplate(`key: "{key}"`).
						AddStr("key", key)
				}

			default:
				return nil, gerrors.NewParseError().
					SetTemplate(`key: "{key}"`).
					AddStr("key", key)
			}

			if last.Kind() == reflect.Interface {
				last = last.Elem()
			}

			if last, ok = normalize(last); ok {
				switch parent.Kind() {
				case reflect.Map:
					parent.SetMapIndex(reflect.ValueOf(key), last)
				case reflect.Array, reflect.Slice:
					i64, err := gnum.ParseInt(key)
					if err != nil {
						return nil, err
					}
					index := int(i64)
					if index < parent.Len() {
						val := parent.Index(int(index))
						val.Set(last)
					} else {
						return nil, gerrors.NewParseError().
							SetTemplate(`key: "{key}"`).
							AddStr("key", key)
					}
				default:
					return nil, gerrors.NewParseError().
						SetTemplate(`key: "{key}"`).
						AddStr("key", key)
				}
			}
		}
		key = keys[len(keys)-1]
	}

	rValue := greflect.LastElem(reflect.ValueOf(value))
	inccorectType := last.Type().Elem() != rValue.Type()

	if inccorectType {
		// Считаем типы совпдающими, если принимающий тип является any
		if last.Type().Elem() == reflect.TypeOf((*any)(nil)).Elem() {
			inccorectType = false
		}
	}

	switch k := last.Kind(); k {
	case reflect.Array, reflect.Slice:
		i64, err := gnum.ParseInt(key)
		if err != nil {
			return nil, err
		}
		index := int(i64)
		if index >= last.Len() {
			return nil, gerrors.NewParseError().
				SetTemplate(`key: "{key}"`).
				AddStr("key", key)
		}
		if inccorectType {
			slice := make([]any, last.Len())
			var i int
			for i = 0; i < last.Len(); i++ {
				slice[i] = last.Index(i).Interface()
			}
			slice[index] = rValue.Interface()
			last = reflect.ValueOf(slice)
		} else {
			last.Index(index).Set(rValue)
		}

	case reflect.Map:
		if inccorectType {
			mapType := reflect.MapOf(last.Type().Key(), reflect.TypeOf((*any)(nil)).Elem())
			newMap := reflect.MakeMapWithSize(mapType, 0)
			for _, rk := range last.MapKeys() {
				rv := last.MapIndex(rk)
				newMap.SetMapIndex(rk, rv)
			}
			last = newMap
		}
		last.SetMapIndex(reflect.ValueOf(key), rValue)

	default:
		return nil, gerrors.NewUnknownError()
	}

	if len(keys) > 1 {
		switch parent.Kind() {
		case reflect.Map:
			parent.SetMapIndex(reflect.ValueOf(lastGroupKey), last)
		case reflect.Array, reflect.Slice:
			i64, err := gnum.ParseInt(lastGroupKey)
			if err != nil {
				return nil, err
			}
			parent.Index(int(i64)).Set(last)
		default:
			return nil, gerrors.NewUnknownError()
		}
	} else {
		root = last
	}

	return root.Interface(), nil
}

// Конвертирует:
// - map[key]value{} в map[key]any{}
// - []value{} в []any{}
// - value в map[string]any{}
func normalize(m reflect.Value) (res reflect.Value, normalized bool) {
	k := m.Kind()

	if k == reflect.Array || k == reflect.Slice {
		if m.Type().Elem() != reflect.TypeOf((*any)(nil)).Elem() {
			slice := make([]any, 0, m.Len())
			for i := 0; i < m.Len(); i++ {
				slice = append(slice, m.Index(i).Interface())
			}
			return reflect.ValueOf(slice), true
		}
		return m, false

	} else if k != reflect.Map {
		return reflect.ValueOf(map[string]any{}), true
	}

	// Проверяем, что map[key]value соответствует шаблону map[key]any
	if m.Type().Elem() == reflect.TypeOf((*any)(nil)).Elem() {
		return m, false
	}

	// Если это не таксоздаем новый map и копируем в него значения из старого map'а
	mapType := reflect.MapOf(m.Type().Key(), reflect.TypeOf((*any)(nil)).Elem())
	res = reflect.MakeMapWithSize(mapType, 0)
	//res := reflect.ValueOf(map[string]any{})

	for _, rk := range m.MapKeys() {
		rv := m.MapIndex(rk)
		res.SetMapIndex(rk, rv)
	}
	return res, true
}
