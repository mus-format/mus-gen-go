package typename

import (
	"errors"
	"fmt"
	"reflect"
)

// ErrTypeMismatch is returned when there is a mismatch between the expected and
// actual types.
var ErrTypeMismatch = errors.New("type mismatch")

// NewNotDefinedTypeError returns an error indicating that a type is not a
// defined type of the specified kind.
func NewNotDefinedTypeError(t reflect.Type, kind string) error {
	return fmt.Errorf("%v is not a defined %v type", t, kind)
}

// NewInvalidPkgPathError returns an error indicating that a package path string
// has an invalid format.
func NewInvalidPkgPathError(str string) error {
	return fmt.Errorf("invalid '%v' pkg path format", str)
}

// NewInvalidPackageError returns an error indicating that a package name
// string has an invalid format.
func NewInvalidPackageError(str string) error {
	return fmt.Errorf("invalid '%v' package format", str)
}

// NewMultiPointerError returns an error indicating that multi-pointer types
// (e.g., **int) are not supported.
func NewMultiPointerError(t reflect.Type) error {
	return fmt.Errorf("do not support multi-pointer types, like %v", t)
}

// NewUnsupportedTypeError returns an error indicating that a name cannot be
// obtained for the given type.
func NewUnsupportedTypeError(t reflect.Type) error {
	return fmt.Errorf("can't get a name for the %v type", t)
}
