package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	testutil "github.com/mus-format/musgen-go/testutil/container"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestContainerTypeGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/container"),
		genops.WithPackage("testutil"),
		genops.WithImport("github.com/mus-format/musgen-go/testutil"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testutil/generic",
			"generic_testdata"),
	)
	assertfatal.EqualError(t, err, nil)

	// array

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyArray]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.LenEncodingMyArray](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ElemEncodingMyArray](),
		typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ElemValidMyArray](),
		typeops.WithElem(typeops.WithValidator("testutil.ValidateZeroValue[int]")))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ValidMyArray](),
		typeops.WithValidator("testutil.ValidateZeroValue[ValidMyArray]"))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.AllMyArray](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testutil.ValidateLength"),
		typeops.WithElem(
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testutil.ValidateZeroValue[int]")),
		typeops.WithValidator("testutil.ValidateZeroValue[AllMyArray]"),
	)
	assertfatal.EqualError(t, err, nil)
	// byte_slice

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyByteSlice]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.LenEncodingMyByteSlice](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.LenValidMyByteSlice](),
		typeops.WithLenValidator("testutil.ValidateLength"))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ValidMyByteSlice](),
		typeops.WithValidator("ValidateByteSlice1"))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.AllMyByteSlice](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testutil.ValidateLength"),
		typeops.WithValidator("ValidateByteSlice2"))
	assertfatal.EqualError(t, err, nil)

	// slice

	err = g.AddDefinedType(reflect.TypeFor[testutil.MySlice]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.LenEncodingMySlice](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.LenValidMySlice](),
		typeops.WithLenValidator("testutil.ValidateLength"))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ElemEncodingMySlice](),
		typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ElemValidMySlice](),
		typeops.WithElem(typeops.WithValidator("testutil.ValidateZeroValue[int]")))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ValidMySlice](),
		typeops.WithValidator("ValidateMySlice1"))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.AllMySlice](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testutil.ValidateLength"),
		typeops.WithElem(
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testutil.ValidateZeroValue[int]")),
		typeops.WithValidator("ValidateMySlice2"))
	assertfatal.EqualError(t, err, nil)

	// map

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyMap]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.LenEncodingMyMap](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.LenValidMyMap](),
		typeops.WithLenValidator("testutil.ValidateLength"))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.KeyEncodingMyMap](),
		typeops.WithKey(typeops.WithNumEncoding(typeops.Raw)))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.KeyValidMyMap](),
		typeops.WithKey(typeops.WithValidator("testutil.ValidateZeroValue[int]")))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ElemEncodingMyMap](),
		typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ElemValidMyMap](),
		typeops.WithElem(typeops.WithValidator("testutil.ValidateZeroValue[string]")))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ValidMyMap](),
		typeops.WithValidator("ValidateMyMap1"))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.AllMyMap](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testutil.ValidateLength"),
		typeops.WithValidator("ValidateMyMap2"),
		typeops.WithKey(
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testutil.ValidateZeroValue[int]")),
		typeops.WithElem(
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testutil.ValidateZeroValue[string]")))
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[testutil.ComplexMap](),
		typeops.WithKey(
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithValidator("testutil.ValidateZeroValue[[3]int]"),
		),
		typeops.WithElem(
			typeops.WithElem(
				typeops.WithLenEncoding(typeops.Raw),
			),
		),
	)
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../testutil/container/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
