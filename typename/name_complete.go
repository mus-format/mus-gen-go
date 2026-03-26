package typename

import (
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/mus-format/mus-gen-go/classifier"
)

// CompleteName examples: "*PkgPath/Pkg.TypeName", "[]PkgPath/Pkg.TypeName",
// "PkgPath/Pkg.Type[PkgPath.Type]".
//
// NOTE: For type parameters CompleteName contains only PkgPath, not PkgPath/Pkg.
type CompleteName string

func MustTypeCompleteName(t reflect.Type) (name CompleteName) {
	name, err := TypeCompleteName(t)
	if err != nil {
		panic(err)
	}
	return
}

// TypeCompleteName for types like 'type MyType *int' will return complete name
// of MyType, not *int.
func TypeCompleteName(t reflect.Type) (name CompleteName, err error) {
	var stars string
	if strings.HasPrefix(t.String(), "*") {
		if stars, t, err = parsePtr(t); err != nil {
			return
		}
	}
	return typeCompleteName(stars, t)
}

func BaseTypeCompleteName(t reflect.Type) (name CompleteName, err error) {
	if !classifier.DefinedType(t) {
		err = ErrTypeMismatch
		return
	}
	switch {
	case classifier.PtrType(t):
		var stars string
		if stars, t, err = parsePtr(t); err != nil {
			return
		}
		return typeCompleteName(stars, t)
	case classifier.PrimitiveType(t):
		return primitiveTypeCompleteName("", t)
	case classifier.ArrayType(t):
		return arrayTypeCompleteName("", t)
	case classifier.SliceType(t):
		return sliceTypeCompleteName("", t)
	case classifier.MapType(t):
		return mapTypeCompleteName("", t)
	default:
		err = NewUnsupportedTypeError(t)
		return
	}
}

func ParsePtr(t reflect.Type) (stars string, at reflect.Type) {
	at = t
	for {
		if !classifier.PtrType(at) {
			return
		}
		at = at.Elem()
		stars += "*"
		continue
	}
}

func parsePtr(t reflect.Type) (stars string, at reflect.Type, err error) {
	stars, at = ParsePtr(t)
	if len(stars) > 1 {
		err = NewMultiPointerError(t)
	}
	return
}

func typeCompleteName(stars string, t reflect.Type) (name CompleteName,
	err error) {
	switch {
	case classifier.DefinedType(t):
		return definedTypeCompleteName(stars, t)
	case classifier.PrimitiveType(t):
		return primitiveTypeCompleteName(stars, t)
	case classifier.ArrayType(t):
		return arrayTypeCompleteName(stars, t)
	case classifier.SliceType(t):
		return sliceTypeCompleteName(stars, t)
	case classifier.MapType(t):
		return mapTypeCompleteName(stars, t)
	default:
		err = NewUnsupportedTypeError(t)
		return
	}
}

func definedTypeCompleteName(stars string, t reflect.Type) (name CompleteName,
	err error) {
	name = CompleteName(stars + filepath.Join(t.PkgPath(), t.String()))
	return
}

func primitiveTypeCompleteName(stars string, t reflect.Type) (
	name CompleteName, err error) {
	name = CompleteName(stars + t.Kind().String())
	return
}

func arrayTypeCompleteName(stars string, t reflect.Type) (name CompleteName,
	err error) {
	elemName, err := TypeCompleteName(t.Elem())
	if err != nil {
		return
	}
	name = CompleteName(stars + "[" + strconv.Itoa(t.Len()) + "]" +
		string(elemName))
	return
}

func sliceTypeCompleteName(stars string, t reflect.Type) (name CompleteName,
	err error) {
	elemName, err := TypeCompleteName(t.Elem())
	if err != nil {
		return
	}
	name = CompleteName(stars + "[]" + string(elemName))
	return
}

func mapTypeCompleteName(stars string, t reflect.Type) (name CompleteName,
	err error) {
	keyName, err := TypeCompleteName(t.Key())
	if err != nil {
		return
	}
	elemName, err := TypeCompleteName(t.Elem())
	if err != nil {
		return
	}
	name = CompleteName(stars + "map[" + string(keyName) + "]" +
		string(elemName))
	return
}
