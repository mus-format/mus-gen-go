package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	struct_testdata "github.com/mus-format/musgen-go/testutil/struct"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestWithNotUnsafeGeneration(t *testing.T) {
	var (
		myIntType         = reflect.TypeFor[struct_testdata.MyInt]()
		mySliceType       = reflect.TypeFor[struct_testdata.MySlice]()
		myStructType      = reflect.TypeFor[struct_testdata.MyStruct]()
		complexStructType = reflect.TypeFor[struct_testdata.ComplexStruct]()
	)

	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/notunsafe"),
		genops.WithPackage("testutil"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testutil/struct",
			"struct_testdata"),
		genops.WithNotUnsafe(),
	)
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(myIntType)
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(mySliceType)
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(myStructType)
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(complexStructType)
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/notunsafe/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)

}
