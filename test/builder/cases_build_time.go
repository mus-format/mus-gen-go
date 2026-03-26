package builder

import (
	"reflect"
	"time"

	genops "github.com/mus-format/mus-gen-go/options/gen"
	stopts "github.com/mus-format/mus-gen-go/options/struct"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/spec"
	bldr "github.com/mus-format/mus-gen-go/spec/builder"
	"github.com/mus-format/mus-gen-go/test/mock"
	"github.com/mus-format/mus-gen-go/typename"
	"github.com/ymz-ncnk/mok"
)

func BuildTimeTestCase() BuildTimeTypeTestCase {
	name := "Should succeed for time.Time"

	var (
		converter = mock.NewTypeNameConvertor()
		gops      = genops.New()
		uops      = stopts.UnderlyingTimeOptions{
			TimeUnit: tpopts.TimeUnitSec,
		}
		mocks = []*mok.Mock{converter.Mock}
	)

	converter.RegisterConvertToFullName(
		func(cname typename.CompleteName) typename.FullName {
			if cname == "time/time.Time" {
				return "time.Time"
			}
			panic("unexpected cname: " + cname)
		},
	)

	var tops tpopts.Options
	tpopts.Apply(&tops, tpopts.WithTimeUnit(tpopts.TimeUnitSec))

	return BuildTimeTypeTestCase{
		Name: name,
		Setup: BuildTimeTypeSetup{
			Converter: converter,
			Gops:      gops,
		},
		Params: BuildTimeTypeParams{
			Type: reflect.TypeFor[time.Time](),
			Uops: uops,
		},
		Want: BuildTimeTypeWant{
			DefinedType: spec.DefinedType{
				FullName:           "time.Time",
				UnderlyingTypeName: "time.Time",
				Tops:               tops,
				Gops:               gops,
			},
			Mocks: mocks,
		},
	}
}

func BuildTimeForNotStructTestCase() BuildTimeTypeTestCase {
	name := "Should fail if receives not struct type"

	var gops = genops.New()

	return BuildTimeTypeTestCase{
		Name: name,
		Setup: BuildTimeTypeSetup{
			Gops: gops,
		},
		Params: BuildTimeTypeParams{
			Type: reflect.TypeFor[int](),
		},
		Want: BuildTimeTypeWant{
			Err: bldr.NewNotStructError(reflect.TypeFor[int]()),
		},
	}
}
