package builder

import (
	"reflect"

	genops "github.com/mus-format/mus-gen-go/options/gen"
	intropts "github.com/mus-format/mus-gen-go/options/interface"
	"github.com/mus-format/mus-gen-go/spec"
	bldr "github.com/mus-format/mus-gen-go/spec/builder"
	"github.com/mus-format/mus-gen-go/test/mock"
	"github.com/mus-format/mus-gen-go/typename"
	"github.com/ymz-ncnk/mok"
)

func BuildInterfaceTestCase() BuildInterfaceTypeTestCase {
	name := "Should succeed for interface"

	var (
		converter = mock.NewTypeNameConvertor()
		gops      = genops.New()
		iops      = intropts.Options{}
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

	return BuildInterfaceTypeTestCase{
		Name: name,
		Setup: BuildInterfaceTypeSetup{
			Converter: converter,
			Gops:      gops,
		},
		Params: BuildInterfaceTypeParams{
			Type: reflect.TypeFor[Interface](),
			Iops: iops,
		},
		Want: BuildInterfaceTypeWant{
			InterfaceType: spec.InterfaceType{
				FullName: "builder.Interface",
				Impls:    []typename.FullName{},
				Iops:     iops,
				Gops:     gops,
			},
			Mocks: mocks,
		},
	}
}

func BuildInterfaceForDefinedTypeTestCase() BuildInterfaceTypeTestCase {
	name := "Should fail if receives defined type"

	var gops = genops.New()

	return BuildInterfaceTypeTestCase{
		Name: name,
		Setup: BuildInterfaceTypeSetup{
			Gops: gops,
		},
		Params: BuildInterfaceTypeParams{
			Type: reflect.TypeFor[Int](),
			Iops: intropts.Options{},
		},
		Want: BuildInterfaceTypeWant{
			Err: bldr.NewUnexpectedDefinedTypeError(reflect.TypeFor[Int]()),
		},
	}
}

func BuildInterfaceForStructTestCase() BuildInterfaceTypeTestCase {
	name := "Should fail if receives struct type"

	var gops = genops.New()

	return BuildInterfaceTypeTestCase{
		Name: name,
		Setup: BuildInterfaceTypeSetup{
			Gops: gops,
		},
		Params: BuildInterfaceTypeParams{
			Type: reflect.TypeFor[Struct](),
			Iops: intropts.Options{},
		},
		Want: BuildInterfaceTypeWant{
			Err: bldr.NewUnexpectedStructTypeError(reflect.TypeFor[Struct]()),
		},
	}
}

func BuildInterfaceForUnsupportedTypeTestCase() BuildInterfaceTypeTestCase {
	name := "Should fail if receives unsupported type"

	var gops = genops.New()

	return BuildInterfaceTypeTestCase{
		Name: name,
		Setup: BuildInterfaceTypeSetup{
			Gops: gops,
		},
		Params: BuildInterfaceTypeParams{
			Type: reflect.TypeFor[int](),
			Iops: intropts.Options{},
		},
		Want: BuildInterfaceTypeWant{
			Err: bldr.NewUnsupportedTypeError(reflect.TypeFor[int]()),
		},
	}
}
