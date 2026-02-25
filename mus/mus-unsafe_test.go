package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	struct_testdata "github.com/mus-format/musgen-go/testutil/struct"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestWithUnsafeGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/unsafe"),
		genops.WithPackage("testutil"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testutil/struct",
			"struct_testdata"),
		genops.WithUnsafe(),
	)
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[struct_testdata.MyInt]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[struct_testdata.MySlice]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(reflect.TypeFor[struct_testdata.MyStruct]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDTS(reflect.TypeFor[struct_testdata.MyInt]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddInterface(reflect.TypeFor[struct_testdata.MyInterface](),
		introps.WithImpl(reflect.TypeFor[struct_testdata.MyInt]()))
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(reflect.TypeFor[struct_testdata.ComplexStruct]())
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/unsafe/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
