package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	strm_testdata "github.com/mus-format/musgen-go/testutil/stream"
	struct_testdata "github.com/mus-format/musgen-go/testutil/struct"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestWithStreamGeneration(t *testing.T) {

	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/stream"),
		genops.WithPackage("testutil"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testutil/struct",
			"struct_testdata"),
		genops.WithSerName(reflect.TypeFor[struct_testdata.MyInterface](),
			"MyAnotherInterface"),
		genops.WithStream(),
	)
	assertfatal.EqualError(t, err, nil)

	// struct

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

	// interface

	tp1 := reflect.TypeFor[strm_testdata.Impl1]()
	err = g.AddStruct(tp1)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDTS(tp1)
	assertfatal.EqualError(t, err, nil)

	tp2 := reflect.TypeFor[strm_testdata.Impl2]()
	err = g.AddDefinedType(tp2)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDTS(tp2)
	assertfatal.EqualError(t, err, nil)

	err = g.AddInterface(reflect.TypeFor[strm_testdata.MyInterface](),
		introps.WithImpl(tp1),
		introps.WithImpl(tp2),
		introps.WithMarshaller(),
	)
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/stream/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
