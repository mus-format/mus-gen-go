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

func BuildTypedDefinedTestCase() BuildTypedTestCase {
	name := "Should succeed for defined basic type"

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
	)

	return BuildTypedTestCase{
		Name: name,
		Setup: BuildTypedSetup{
			Converter: converter,
			Gops:      gops,
		},
		Params: BuildTypedParams{
			Type: reflect.TypeFor[Int](),
		},
		Want: BuildTypedWant{
			Typed: spec.Type{
				FullName: "builder.Int",
				Gops:     gops,
			},
			Mocks: mocks,
		},
	}
}

func BuildTypedStructTestCase() BuildTypedTestCase {
	name := "Should succeed for struct type"

	var (
		converter = mock.NewTypeNameConvertor()
		gops      = genops.New()
		mocks     = []*mok.Mock{converter.Mock}
	)

	converter.RegisterConvertToFullName(
		func(cname typename.CompleteName) typename.FullName {
			if cname == "github.com/mus-format/mus-gen-go/test/builder/builder.Struct" {
				return "builder.Struct"
			}
			panic("unexpected cname: " + cname)
		},
	)

	return BuildTypedTestCase{
		Name: name,
		Setup: BuildTypedSetup{
			Converter: converter,
			Gops:      gops,
		},
		Params: BuildTypedParams{
			Type: reflect.TypeFor[Struct](),
		},
		Want: BuildTypedWant{
			Typed: spec.Type{
				FullName: "builder.Struct",
				Gops:     gops,
			},
			Mocks: mocks,
		},
	}
}

func BuildTypedInterfaceTestCase() BuildTypedTestCase {
	name := "Should succeed for interface type"

	var (
		converter = mock.NewTypeNameConvertor()
		gops      = genops.New()
		mocks     = []*mok.Mock{converter.Mock}
	)

	converter.RegisterConvertToFullName(
		func(cname typename.CompleteName) typename.FullName {
			if cname == "github.com/mus-format/mus-gen-go/test/builder/builder.Interface" {
				return "builder.Interface"
			}
			panic("unexpected cname: " + cname)
		},
	)

	return BuildTypedTestCase{
		Name: name,
		Setup: BuildTypedSetup{
			Converter: converter,
			Gops:      gops,
		},
		Params: BuildTypedParams{
			Type: reflect.TypeFor[Interface](),
		},
		Want: BuildTypedWant{
			Typed: spec.Type{
				FullName: "builder.Interface",
				Gops:     gops,
			},
			Mocks: mocks,
		},
	}
}

func BuildTypedNotSupportedTestCase() BuildTypedTestCase {
	name := "Should fail for not supported type"

	var gops = genops.New()

	return BuildTypedTestCase{
		Name: name,
		Setup: BuildTypedSetup{
			Gops: gops,
		},
		Params: BuildTypedParams{
			Type: reflect.TypeFor[int](),
		},
		Want: BuildTypedWant{
			Err: bldr.NewUnsupportedTypeError(reflect.TypeFor[int]()),
		},
	}
}
