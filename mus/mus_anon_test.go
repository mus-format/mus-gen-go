package musgen_test

import (
	"os"
	"reflect"
	"testing"

	musgen "github.com/mus-format/mus-gen-go/mus"
	genops "github.com/mus-format/mus-gen-go/options/gen"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	types "github.com/mus-format/mus-gen-go/test/types"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestGenerate_Anon(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types"),
	)
	assertfatal.EqualError(t, err, nil)

	// LenString

	err = g.AddDefinedType(reflect.TypeFor[types.LenString](),
		tpopts.WithLenEncoding(tpopts.NumEncodingRaw),
	)
	assertfatal.EqualError(t, err, nil)

	// LenSlice

	err = g.AddDefinedType(reflect.TypeFor[types.LenSlice](),
		tpopts.WithLenEncoding(tpopts.NumEncodingRaw),
	)
	assertfatal.EqualError(t, err, nil)

	// LenArray

	err = g.AddDefinedType(reflect.TypeFor[types.LenArray](),
		tpopts.WithLenEncoding(tpopts.NumEncodingRaw),
	)
	assertfatal.EqualError(t, err, nil)

	// LenMap

	err = g.AddDefinedType(reflect.TypeFor[types.LenMap](),
		tpopts.WithLenEncoding(tpopts.NumEncodingRaw),
	)
	assertfatal.EqualError(t, err, nil)

	// Slice

	err = g.AddDefinedType(reflect.TypeFor[types.Slice]())
	assertfatal.EqualError(t, err, nil)

	// Array

	err = g.AddDefinedType(reflect.TypeFor[types.Array]())
	assertfatal.EqualError(t, err, nil)

	// Map

	err = g.AddDefinedType(reflect.TypeFor[types.Map]())
	assertfatal.EqualError(t, err, nil)

	// Ptr

	err = g.AddDefinedType(reflect.TypeFor[types.Ptr]())
	assertfatal.EqualError(t, err, nil)

	// Valid String

	err = g.AddDefinedType(reflect.TypeFor[types.ValidString](),
		tpopts.WithLenValidator("ValidateLen"),
	)
	assertfatal.EqualError(t, err, nil)

	// Valid Slice

	err = g.AddDefinedType(reflect.TypeFor[types.ValidSlice](),
		tpopts.WithLenValidator("ValidateLen"),
		tpopts.WithElemValidator("ValidateNum"),
	)
	assertfatal.EqualError(t, err, nil)

	// Valid Array

	err = g.AddDefinedType(reflect.TypeFor[types.ValidArray](),
		tpopts.WithElemValidator("ValidateNum"),
	)
	assertfatal.EqualError(t, err, nil)

	// Valid Map

	err = g.AddDefinedType(reflect.TypeFor[types.ValidMap](),
		tpopts.WithLenValidator("ValidateLen"),
		tpopts.WithKeyValidator("ValidateStr"),
		tpopts.WithElemValidator("ValidateNum"),
	)
	assertfatal.EqualError(t, err, nil)

	// ComplexMap

	err = g.AddDefinedType(reflect.TypeFor[types.ComplexMap](),
		tpopts.WithLenValidator("ValidateLen"),
		tpopts.WithElemValidator("ValidateComplexMapValue"),
	)
	assertfatal.EqualError(t, err, nil)

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/anon_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
