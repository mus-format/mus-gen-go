package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	typeops "github.com/mus-format/musgen-go/options/type"

	testutil "github.com/mus-format/musgen-go/testutil/register_interface"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestInterfaceTypeRegistration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/register_interface"),
		genops.WithPackage("testutil"),
	)
	assertfatal.EqualError(t, err, nil)

	err = g.RegisterInterface(reflect.TypeFor[testutil.MyInterface](),
		introps.WithStructImpl(reflect.TypeFor[testutil.Impl1]()),
		introps.WithDefinedTypeImpl(reflect.TypeFor[testutil.Impl2](),
			typeops.WithNumEncoding(typeops.Raw),
		),
	)
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/register_interface/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
