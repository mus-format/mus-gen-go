package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	testutil "github.com/mus-format/musgen-go/testutil/primitive"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestPrimitiveTypesGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/primitive"),
		genops.WithPackage("testutil"),
		genops.WithImport("github.com/mus-format/musgen-go/testutil"),
	)
	assertfatal.EqualError(t, err, nil)

	// bool

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyBool]())
	assertfatal.EqualError(t, err, nil)

	// byte

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyByte]())
	assertfatal.EqualError(t, err, nil)

	// float32

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyFloat32]())
	assertfatal.EqualError(t, err, nil)

	// float64

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyFloat64]())
	assertfatal.EqualError(t, err, nil)

	// int

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyInt]())
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[testutil.RawMyInt](),
		typeops.WithNumEncoding(typeops.Raw),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[testutil.VarintPositiveMyInt](),
		typeops.WithNumEncoding(typeops.VarintPositive),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[testutil.ValidMyInt](),
		typeops.WithValidator("testutil.ValidateZeroValue"))
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[testutil.AllMyInt](),
		typeops.WithNumEncoding(typeops.Raw),
		typeops.WithValidator("testutil.ValidateZeroValue"))
	assertfatal.EqualError(t, err, nil)

	// string

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyString]())
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[testutil.LenEncodingMyString](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[testutil.LenValidMyString](),
		typeops.WithLenValidator("testutil.ValidateLength3"))
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[testutil.ValidMyString](),
		typeops.WithValidator("testutil.ValidateZeroValue"))
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[testutil.AllMyString](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testutil.ValidateLength3"),
		typeops.WithValidator("testutil.ValidateZeroValue"))
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/primitive/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
