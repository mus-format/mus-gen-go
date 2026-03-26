package builder

import (
	"fmt"
	"reflect"

	"github.com/mus-format/mus-gen-go/typename"
)

// NewUnexpectedDefinedTypeError returns an error indicating that the specifiedVG
// type t should be added using the Generator.AddDefinedType method.
func NewUnexpectedDefinedTypeError(t reflect.Type) error {
	return fmt.Errorf("use the Generator.AddDefinedType method for the %v type ",
		t)
}

// NewUnexpectedStructTypeError returns an error indicating that the specified
// struct type t should be added using the Generator.AddStruct method.
func NewUnexpectedStructTypeError(t reflect.Type) error {
	return fmt.Errorf("use the Generator.AddStruct method for the %v type", t)
}

// NewUnexpectedInterfaceTypeError returns an error indicating that the
// specified interface type t should be added using the
// Generator.AddInterface method.
func NewUnexpectedInterfaceTypeError(t reflect.Type) error {
	return fmt.Errorf("use the Generator.AddInterface method for the %v type",
		t)
}

// NewUnsupportedTypeError returns an error indicating that the provided type is
// not supported.
func NewUnsupportedTypeError(t reflect.Type) error {
	return fmt.Errorf("unsupported %v type", t)
}

func NewNotDefinedTypeError(t reflect.Type) error {
	return fmt.Errorf("requires a defined type, but %v is not", t)
}

func NewNotStructError(t reflect.Type) error {
	return fmt.Errorf("requires a defined struct, but %v is not", t)
}

func NewNotInterfaceError(t reflect.Type) error {
	return fmt.Errorf("requires a defined interface, but %v is not", t)
}

func NewTwoPathsSameAliasError(pkgPath1, pkgPath2 typename.PkgPath,
	alias typename.Package) error {
	return fmt.Errorf("two pkgPath '%v' and '%v' have the same alias '%v'",
		pkgPath1, pkgPath2, alias)
}

func NewWrongFieldsCountError(want int) error {
	return fmt.Errorf("not enough structops.WithField() calls, want %v", want)
}
