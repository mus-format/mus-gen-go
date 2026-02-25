package musgen

import (
	"os"
	"reflect"
	"testing"

	"github.com/mus-format/musgen-go/data/builders"
	genops "github.com/mus-format/musgen-go/options/generate"
	testutil "github.com/mus-format/musgen-go/testutil/any"
	"github.com/mus-format/musgen-go/typename"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestAnyGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/any"),
		genops.WithPackage("testutil"))
	assertfatal.EqualError(t, err, nil)

	tp := reflect.TypeFor[testutil.MyAny]()
	err = g.AddDefinedType(tp)
	assertfatal.EqualError(t, err, builders.NewUnsupportedTypeError(tp))

	anyType := reflect.TypeFor[any]()

	tp = reflect.TypeFor[testutil.MyAnySlice]()
	err = g.AddDefinedType(tp)
	assertfatal.EqualError(t, err, typename.NewUnsupportedTypeError(
		reflect.TypeFor[any]()))

	err = g.AddStruct(reflect.TypeFor[testutil.MyAnyStruct]())
	assertfatal.EqualError(t, err, typename.NewUnsupportedTypeError(anyType))

	tp = reflect.TypeFor[testutil.MyAnyGenericSlice[any]]()
	err = g.AddDefinedType(tp)
	assertfatal.EqualError(t, err, typename.NewUnsupportedTypeError(
		reflect.TypeFor[any]()))

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/any/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
