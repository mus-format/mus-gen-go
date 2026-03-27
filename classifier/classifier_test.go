package classifier

import (
	"reflect"
	"testing"

	types "github.com/mus-format/mus-gen-go/test/classifier"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestClassifier(t *testing.T) {
	var (
		undefinedStructType     = reflect.TypeFor[struct{}]()
		definedStructType       = reflect.TypeFor[types.Struct]()
		doubleDefinedStructType = reflect.TypeFor[types.DoubleDefinedStruct]()

		undefinedInterfaceType     = reflect.TypeFor[interface{ Print() }]()
		definedInterfaceType       = reflect.TypeFor[types.Interface]()
		doubleDefinedInterfaceType = reflect.TypeFor[types.DoubleDefinedInterface]()
		anyType                    = reflect.TypeFor[any]()

		undefinedPrimitiveType = reflect.TypeFor[int]()
		definedPrimitiveType   = reflect.TypeFor[types.Int]()

		undefinedContainerType           = reflect.TypeFor[[]int]()
		definedContainerType             = reflect.TypeFor[types.Slice]()
		definedParametrizedContainerType = reflect.TypeFor[types.GenericSlice[int]]()

		ptrPrimitiveType = reflect.TypeFor[*int]()
		ptrContainerType = reflect.TypeFor[*[]int]()
		ptrStructType    = reflect.TypeFor[*struct{}]()
		ptrInterfaceType = reflect.TypeFor[*interface{ Print() }]()

		ptrDefinedPrimitiveType = reflect.TypeFor[types.IntPtr]()
		ptrDefinedContainerType = reflect.TypeFor[types.SlicePtr]()
		ptrDefinedStructType    = reflect.TypeFor[types.StructPtr]()
		ptrDefinedInterfaceType = reflect.TypeFor[types.InterfacePtr]()
	)

	t.Run("DefinedBasicType", func(t *testing.T) {
		testCases := []struct {
			tp   reflect.Type
			want bool
		}{
			{tp: undefinedStructType, want: false},
			{tp: definedStructType, want: false},
			{tp: doubleDefinedStructType, want: false},
			{tp: undefinedInterfaceType, want: false},
			{tp: definedInterfaceType, want: false},
			{tp: doubleDefinedInterfaceType, want: false},
			{tp: anyType, want: false},
			{tp: undefinedPrimitiveType, want: false},
			{tp: undefinedContainerType, want: false},
			{tp: ptrPrimitiveType, want: false},
			{tp: ptrContainerType, want: false},
			{tp: ptrStructType, want: false},
			{tp: ptrInterfaceType, want: false},

			{tp: definedPrimitiveType, want: true},
			{tp: definedContainerType, want: true},
			{tp: definedParametrizedContainerType, want: true},
			{tp: ptrDefinedPrimitiveType, want: true},
			{tp: ptrDefinedContainerType, want: true},
			{tp: ptrDefinedStructType, want: true},
			{tp: ptrDefinedInterfaceType, want: true},
		}
		for _, c := range testCases {
			asserterror.Equal(t, DefinedBasicType(c.tp), c.want)
		}
	})

	t.Run("DefinedStruct", func(t *testing.T) {
		testCases := []struct {
			tp   reflect.Type
			want bool
		}{
			{tp: undefinedStructType, want: false},
			{tp: undefinedInterfaceType, want: false},
			{tp: definedInterfaceType, want: false},
			{tp: doubleDefinedInterfaceType, want: false},
			{tp: anyType, want: false},
			{tp: undefinedPrimitiveType, want: false},
			{tp: definedPrimitiveType, want: false},
			{tp: undefinedContainerType, want: false},
			{tp: definedContainerType, want: false},
			{tp: definedParametrizedContainerType, want: false},
			{tp: ptrPrimitiveType, want: false},
			{tp: ptrContainerType, want: false},
			{tp: ptrStructType, want: false},
			{tp: ptrInterfaceType, want: false},
			{tp: ptrDefinedPrimitiveType, want: false},
			{tp: ptrDefinedContainerType, want: false},
			{tp: ptrDefinedStructType, want: false},
			{tp: ptrDefinedInterfaceType, want: false},

			{tp: definedStructType, want: true},
			{tp: doubleDefinedStructType, want: true},
		}
		for _, c := range testCases {
			asserterror.Equal(t, DefinedStruct(c.tp), c.want)
		}
	})

	t.Run("DefinedInterface", func(t *testing.T) {
		testCases := []struct {
			tp   reflect.Type
			want bool
		}{
			{tp: undefinedStructType, want: false},
			{tp: definedStructType, want: false},
			{tp: doubleDefinedStructType, want: false},
			{tp: undefinedInterfaceType, want: false},
			{tp: anyType, want: false},
			{tp: undefinedPrimitiveType, want: false},
			{tp: definedPrimitiveType, want: false},
			{tp: undefinedContainerType, want: false},
			{tp: definedContainerType, want: false},
			{tp: definedParametrizedContainerType, want: false},
			{tp: ptrPrimitiveType, want: false},
			{tp: ptrContainerType, want: false},
			{tp: ptrStructType, want: false},
			{tp: ptrInterfaceType, want: false},
			{tp: ptrDefinedPrimitiveType, want: false},
			{tp: ptrDefinedContainerType, want: false},
			{tp: ptrDefinedStructType, want: false},
			{tp: ptrDefinedInterfaceType, want: false},

			{tp: definedInterfaceType, want: true},
			{tp: doubleDefinedInterfaceType, want: true},
		}
		for _, c := range testCases {
			asserterror.Equal(t, DefinedNonEmptyInterface(c.tp), c.want)
		}
	})
}
