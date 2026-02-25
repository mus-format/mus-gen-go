package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	sttime_testdata "github.com/mus-format/musgen-go/testutil/struct_time"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestStructTimeGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/struct_time"),
		genops.WithPackage("testutil"),
	)
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(reflect.TypeFor[sttime_testdata.MyMicroTime](),
		structops.WithTops(
			typeops.WithSourceType(typeops.Time),
			typeops.WithTimeUnit(typeops.Micro),
		))
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(reflect.TypeFor[sttime_testdata.MyDefaultTime](),
		structops.WithTops(
			typeops.WithSourceType(typeops.Time),
		),
	)
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/struct_time/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)

}
