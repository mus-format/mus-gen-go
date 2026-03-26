package musgen_test

import (
	"os"
	"reflect"
	"testing"

	musgen "github.com/mus-format/mus-gen-go/mus"
	genops "github.com/mus-format/mus-gen-go/options/gen"
	types "github.com/mus-format/mus-gen-go/test/types"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestGenerateTyped(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types"),
	)
	assertfatal.EqualError(t, err, nil)
	typedIntType := reflect.TypeFor[types.TypedInt]()
	err = g.AddDefinedType(typedIntType)
	assertfatal.EqualError(t, err, nil)
	err = g.AddTyped(typedIntType)
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/typed_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
