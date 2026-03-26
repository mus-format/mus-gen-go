// Package classifier provides functions for classifying Go types.
package classifier

import (
	"reflect"
)

func DefinedNonEmptyInterface(t reflect.Type) (ok bool) {
	return DefinedType(t) && InterfaceType(t) && t.NumMethod() > 0
}

func DefinedInterface(t reflect.Type) (ok bool) {
	return DefinedType(t) && InterfaceType(t)
}

func DefinedBasicType(t reflect.Type) (ok bool) {
	return DefinedType(t) && BasicType(t)
}

func DefinedStruct(t reflect.Type) (ok bool) {
	return DefinedType(t) && t.Kind() == reflect.Struct
}

func DefinedType(t reflect.Type) (ok bool) {
	return t.PkgPath() != ""
}

func InterfaceType(t reflect.Type) (ok bool) {
	return t.Kind() == reflect.Interface
}

func BasicType(t reflect.Type) (ok bool) {
	return PrimitiveType(t) || ContainerType(t) || PtrType(t)
}

func PtrType(t reflect.Type) bool {
	return t.Kind() == reflect.Pointer
}

func ContainerType(t reflect.Type) (ok bool) {
	return ArrayType(t) || SliceType(t) || MapType(t)
}

func ArrayType(t reflect.Type) (ok bool) {
	return t.Kind() == reflect.Array
}

func SliceType(t reflect.Type) (ok bool) {
	return t.Kind() == reflect.Slice
}

func MapType(t reflect.Type) (ok bool) {
	return t.Kind() == reflect.Map
}

func PrimitiveType(t reflect.Type) (ok bool) {
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String:
		return true
	default:
		return
	}
}
