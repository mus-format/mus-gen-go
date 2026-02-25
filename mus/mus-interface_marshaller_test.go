package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	testutil "github.com/mus-format/musgen-go/testutil/interface_marshaller"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestInterfaceTypeWithMarshallerGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/interface_marshaller"),
		genops.WithPackage("testutil"),
	)
	assertfatal.EqualError(t, err, nil)

	tp1 := reflect.TypeFor[testutil.Impl1]()
	err = g.AddStruct(tp1)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDTS(tp1)
	assertfatal.EqualError(t, err, nil)

	tp2 := reflect.TypeFor[testutil.Impl2]()
	err = g.AddStruct(tp2)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDTS(tp2)
	assertfatal.EqualError(t, err, nil)

	err = g.AddInterface(reflect.TypeFor[testutil.MyInterface](),
		introps.WithImpl(tp1),
		introps.WithImpl(tp2),
		introps.WithMarshaller())
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/interface_marshaller/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
