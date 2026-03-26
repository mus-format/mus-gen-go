package builder

import (
	"reflect"

	genops "github.com/mus-format/mus-gen-go/options/gen"
	"github.com/mus-format/mus-gen-go/spec"
	bldr "github.com/mus-format/mus-gen-go/spec/builder"
	"github.com/mus-format/mus-gen-go/test/mock"
	"github.com/mus-format/mus-gen-go/typename"
	"github.com/ymz-ncnk/mok"
)

func BuildDefinedPrimitiveTestCase() BuildDefinedTypeTestCase {
	name := "Should succeed for primitive type"

	var (
		converter = mock.NewTypeNameConvertor()
		gops      = genops.New()
		mocks     = []*mok.Mock{converter.Mock}
	)

	converter.RegisterConvertToFullName(
		func(cname typename.CompleteName) typename.FullName {
			if cname == "github.com/mus-format/mus-gen-go/test/builder/builder.Int" {
				return "builder.Int"
			}
			panic("unexpected cname: " + cname)
		},
	).RegisterConvertToFullName(
		func(cname typename.CompleteName) typename.FullName {
			if cname == "int" {
				return "int"
			}
			panic("unexpected cname: " + cname)
		},
	)

	return BuildDefinedTypeTestCase{
		Name: name,
		Setup: BuildDefinedTypeSetup{
			Converter: converter,
			Gops:      gops,
		},
		Params: BuildDefinedTypeParams{
			Type: reflect.TypeFor[Int](),
		},
		Want: BuildDefinedTypeWant{
			DefinedType: spec.DefinedType{
				FullName:           "builder.Int",
				UnderlyingTypeName: "int",
				Gops:               gops,
			},
			Mocks: mocks,
		},
	}
}

func BuildDefinedSliceTestCase() BuildDefinedTypeTestCase {
	name := "Should succeed for slice type"

	var (
		converter = mock.NewTypeNameConvertor()
		gops      = genops.New()
		mocks     = []*mok.Mock{converter.Mock}
	)

	converter.RegisterConvertToFullName(
		func(cname typename.CompleteName) typename.FullName {
			if cname == "github.com/mus-format/mus-gen-go/test/builder/builder.Slice" {
				return "builder.Slice"
			}
			panic("unexpected cname: " + cname)
		},
	).RegisterConvertToFullName(
		func(cname typename.CompleteName) typename.FullName {
			if cname == "[]int" {
				return "[]int"
			}
			panic("unexpected cname: " + cname)
		},
	)

	return BuildDefinedTypeTestCase{
		Name: name,
		Setup: BuildDefinedTypeSetup{
			Converter: converter,
			Gops:      gops,
		},
		Params: BuildDefinedTypeParams{
			Type: reflect.TypeFor[Slice](),
		},
		Want: BuildDefinedTypeWant{
			DefinedType: spec.DefinedType{
				FullName:           "builder.Slice",
				UnderlyingTypeName: "[]int",
				Gops:               gops,
			},
			Mocks: mocks,
		},
	}
}

func BuildDefinedNotTestCase() BuildDefinedTypeTestCase {
	name := "Should fail if receives not defined type"

	var gops = genops.New()

	return BuildDefinedTypeTestCase{
		Name: name,
		Setup: BuildDefinedTypeSetup{
			Gops: gops,
		},
		Params: BuildDefinedTypeParams{
			Type: reflect.TypeFor[int](),
		},
		Want: BuildDefinedTypeWant{
			Err: bldr.NewUnsupportedTypeError(reflect.TypeFor[int]()),
		},
	}
}

func BuildDefinedForStructTestCase() BuildDefinedTypeTestCase {
	name := "Should fail if receives struct type"

	var gops = genops.New()

	return BuildDefinedTypeTestCase{
		Name: name,
		Setup: BuildDefinedTypeSetup{
			Gops: gops,
		},
		Params: BuildDefinedTypeParams{
			Type: reflect.TypeFor[Struct](),
		},
		Want: BuildDefinedTypeWant{
			Err: bldr.NewUnexpectedStructTypeError(reflect.TypeFor[Struct]()),
		},
	}
}

func BuildDefinedForInterfaceTestCase() BuildDefinedTypeTestCase {
	name := "Should fail if receives interface type"

	var gops = genops.New()

	return BuildDefinedTypeTestCase{
		Name: name,
		Setup: BuildDefinedTypeSetup{
			Gops: gops,
		},
		Params: BuildDefinedTypeParams{
			Type: reflect.TypeFor[Interface](),
		},
		Want: BuildDefinedTypeWant{
			Err: bldr.NewUnexpectedInterfaceTypeError(reflect.TypeFor[Interface]()),
		},
	}
}

func BuildDefinedForUnsupportedTypeTestCase() BuildDefinedTypeTestCase {
	name := "Should fail if receives unsupported type"

	var gops = genops.New()

	return BuildDefinedTypeTestCase{
		Name: name,
		Setup: BuildDefinedTypeSetup{
			Gops: gops,
		},
		Params: BuildDefinedTypeParams{
			Type: reflect.TypeFor[int](),
		},
		Want: BuildDefinedTypeWant{
			Err: bldr.NewUnsupportedTypeError(reflect.TypeFor[int]()),
		},
	}
}
