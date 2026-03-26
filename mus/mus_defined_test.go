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

func TestGenerate_Defined(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types"),
	)
	assertfatal.EqualError(t, err, nil)

	// Simple

	err = g.AddDefinedType(reflect.TypeFor[types.Int]())
	assertfatal.EqualError(t, err, nil)

	// Raw

	err = g.AddDefinedType(reflect.TypeFor[types.RawInt](),
		tpopts.WithNumEncoding(tpopts.NumEncodingRaw),
	)
	assertfatal.EqualError(t, err, nil)

	// Valid

	err = g.AddDefinedType(reflect.TypeFor[types.ValidInt](),
		tpopts.WithValidator("ValidateInt"),
	)
	assertfatal.EqualError(t, err, nil)

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/defined_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
