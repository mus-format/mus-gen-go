package musgen_test

import (
	"os"
	"reflect"
	"testing"

	musgen "github.com/mus-format/mus-gen-go/mus"
	fldopts "github.com/mus-format/mus-gen-go/options/field"
	genops "github.com/mus-format/mus-gen-go/options/gen"
	stopts "github.com/mus-format/mus-gen-go/options/struct"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	types "github.com/mus-format/mus-gen-go/test/types"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestGenerateStruct(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types"),
	)
	assertfatal.EqualError(t, err, nil)

	// Simple

	err = g.AddStruct(reflect.TypeFor[types.SimpleStruct]())
	assertfatal.EqualError(t, err, nil)

	// Time

	err = g.AddStruct(reflect.TypeFor[types.TimeStruct](),
		stopts.WithUnderlyingTime(),
	)
	assertfatal.EqualError(t, err, nil)

	// Embedding

	err = g.AddStruct(reflect.TypeFor[types.InnerStruct]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(reflect.TypeFor[types.EmbeddingStruct]())
	assertfatal.EqualError(t, err, nil)

	// Ignore

	err = g.AddStruct(reflect.TypeFor[types.IgnoreStruct](),
		stopts.WithField(),
		stopts.WithField(fldopts.WithIgnore()),
	)
	assertfatal.EqualError(t, err, nil)

	// Valid

	err = g.AddStruct(reflect.TypeFor[types.ValidStruct](),
		stopts.WithValidator("ValidateStruct"),

		stopts.WithField(
			fldopts.WithType(
				tpopts.WithValidator("ValidateNum"),
			),
		),
		stopts.WithField(
			fldopts.WithType(
				tpopts.WithValidator("ValidateStr"),
			),
		),
	)
	assertfatal.EqualError(t, err, nil)

	// Parametric
	err = g.AddStruct(reflect.TypeFor[types.ParametricStruct[int]]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/struct_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
