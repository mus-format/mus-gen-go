package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	testutil "github.com/mus-format/musgen-go/testutil/pointer"
	"github.com/mus-format/musgen-go/typename"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestPointerGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/pointer"),
		genops.WithPackage("testutil"),
	)
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyIntPtr]())
	assertfatal.EqualError(t, err, nil)

	tp := reflect.TypeFor[testutil.MyDoubleIntPtr]()
	err = g.AddDefinedType(tp)
	assertfatal.EqualError(t, err, typename.NewMultiPointerError(tp))

	err = g.AddDefinedType(reflect.TypeFor[testutil.MySlicePtr]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(reflect.TypeFor[testutil.MyStruct]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyStructPtr]())
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/pointer/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
