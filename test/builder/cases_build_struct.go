package builder

import (
	"reflect"

	fldopts "github.com/mus-format/mus-gen-go/options/field"
	genops "github.com/mus-format/mus-gen-go/options/gen"
	stopts "github.com/mus-format/mus-gen-go/options/struct"
	"github.com/mus-format/mus-gen-go/spec"
	bldr "github.com/mus-format/mus-gen-go/spec/builder"
	"github.com/mus-format/mus-gen-go/test/mock"
	"github.com/mus-format/mus-gen-go/typename"
	"github.com/ymz-ncnk/mok"
)

func BuildWrongFieldsCountTestCase() BuildStructTypeTestCase {
	name := "Should fail if the number of field options is not equal to the number of struct fields"

	var (
		gops = genops.New()
		sops = stopts.Options{}
	)
	stopts.Apply([]stopts.SetOption{
		stopts.WithField(fldopts.WithIgnore()),
	}, &sops)

	return BuildStructTypeTestCase{
		Name: name,
		Setup: BuildStructTypeSetup{
			Gops: gops,
		},
		Params: BuildStructTypeParams{
			Type: reflect.TypeFor[StructWithTwoFields](),
			Sops: sops,
		},
		Want: BuildStructTypeWant{
			Err: bldr.NewWrongFieldsCountError(2),
		},
	}
}

func BuildEmptyStructTestCase() BuildStructTypeTestCase {
	name := "Should succeed for empty struct"

	var (
		converter = mock.NewTypeNameConvertor()
		gops      = genops.New()
		sops      = stopts.Options{}
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

	return BuildStructTypeTestCase{
		Name: name,
		Setup: BuildStructTypeSetup{
			Converter: converter,
			Gops:      gops,
		},
		Params: BuildStructTypeParams{
			Type: reflect.TypeFor[Struct](),
			Sops: sops,
		},
		Want: BuildStructTypeWant{
			StructType: spec.StructType{
				FullName: "builder.Struct",
				Fields:   []spec.FieldType{},
				Sops:     sops,
				Gops:     gops,
			},
			Mocks: mocks,
		},
	}
}

func BuildStructWithFieldsTestCase() BuildStructTypeTestCase {
	name := "Should succeed for struct with fields"

	var (
		converter = mock.NewTypeNameConvertor()
		gops      = genops.New()
		sops      = stopts.Options{}
		mocks     = []*mok.Mock{converter.Mock}
	)

	converter.RegisterConvertToFullName(
		func(cname typename.CompleteName) typename.FullName {
			if cname == "github.com/mus-format/mus-gen-go/test/builder/builder.StructWithFields" {
				return "builder.StructWithFields"
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

	return BuildStructTypeTestCase{
		Name: name,
		Setup: BuildStructTypeSetup{
			Converter: converter,
			Gops:      gops,
		},
		Params: BuildStructTypeParams{
			Type: reflect.TypeFor[StructWithFields](),
			Sops: sops,
		},
		Want: BuildStructTypeWant{
			StructType: spec.StructType{
				FullName: "builder.StructWithFields",
				Fields: []spec.FieldType{
					{
						FullName:  "int",
						FieldName: "Int",
						Gops:      gops,
					},
				},
				Sops: sops,
				Gops: gops,
			},
			Mocks: mocks,
		},
	}
}

func BuildStructForDefinedTypeTestCase() BuildStructTypeTestCase {
	name := "Should fail if receives defined type"

	var gops = genops.New()

	return BuildStructTypeTestCase{
		Name: name,
		Setup: BuildStructTypeSetup{
			Gops: gops,
		},
		Params: BuildStructTypeParams{
			Type: reflect.TypeFor[Int](),
			Sops: stopts.Options{},
		},
		Want: BuildStructTypeWant{
			Err: bldr.NewUnexpectedDefinedTypeError(reflect.TypeFor[Int]()),
		},
	}
}

func BuildStructForInterfaceTypeTestCase() BuildStructTypeTestCase {
	name := "Should fail if receives interface type"

	var gops = genops.New()

	return BuildStructTypeTestCase{
		Name: name,
		Setup: BuildStructTypeSetup{
			Gops: gops,
		},
		Params: BuildStructTypeParams{
			Type: reflect.TypeFor[Interface](),
			Sops: stopts.Options{},
		},
		Want: BuildStructTypeWant{
			Err: bldr.NewUnexpectedInterfaceTypeError(reflect.TypeFor[Interface]()),
		},
	}
}

func BuildStructForUnsupportedTypeTestCase() BuildStructTypeTestCase {
	name := "Should fail if receives unsupported type"

	var gops = genops.New()

	return BuildStructTypeTestCase{
		Name: name,
		Setup: BuildStructTypeSetup{
			Gops: gops,
		},
		Params: BuildStructTypeParams{
			Type: reflect.TypeFor[int](),
			Sops: stopts.Options{},
		},
		Want: BuildStructTypeWant{
			Err: bldr.NewUnsupportedTypeError(reflect.TypeFor[int]()),
		},
	}
}
