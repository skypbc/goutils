package gvars

import (
	"reflect"
)

// Contains проверяет, содержится ли source в себе элемент типа T
func Contains[T any](source any) bool {
	targetType := reflect.TypeOf((*T)(nil)).Elem()

	for targetType.Kind() == reflect.Pointer {
		targetType = targetType.Elem()
	}

	return contains(reflect.ValueOf(source), targetType)
}

// contains проверяет, содержится ли тип targetType в source (рекурсивно)
func contains(source reflect.Value, targetType reflect.Type) bool {
	if !source.IsValid() || targetType == nil {
		return false
	}

	for source.Kind() == reflect.Pointer {
		source = source.Elem()
	}

	sourceType := source.Type()
	if sourceType == targetType {
		return true
	}

	if sourceType.Kind() == reflect.Struct {
		for i := 0; i < source.NumField(); i++ {
			field := source.Field(i)
			if field.Type() == targetType {
				return true
			}
			if contains(field, targetType) {
				return true
			}
		}
	}

	return false
}
