package musgen_test

import (
	"os"
	"reflect"
	"testing"

	musgen "github.com/mus-format/mus-gen-go/mus"
	genops "github.com/mus-format/mus-gen-go/options/gen"
	veropts "github.com/mus-format/mus-gen-go/options/versioned"
	"github.com/mus-format/mus-gen-go/test/types"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestGenerateVersioned(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types"),
	)
	assertfatal.EqualError(t, err, nil)
	var (
		ver1Type = reflect.TypeFor[types.Ver1]()
		ver2Type = reflect.TypeFor[types.Ver2]()
	)
	err = g.AddDefinedType(ver1Type)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(ver2Type)
	assertfatal.EqualError(t, err, nil)
	err = g.AddTyped(ver1Type)
	assertfatal.EqualError(t, err, nil)
	err = g.AddTyped(ver2Type)
	assertfatal.EqualError(t, err, nil)
	err = g.AddVersioned(reflect.TypeFor[types.Versioned](),
		veropts.WithVersion(ver1Type, "MigrateVer1"),
		veropts.WithCurrentVersion(ver2Type),
	)
	assertfatal.EqualError(t, err, nil)

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/versioned_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

func TestGenerateVersioned_Register(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types"),
	)
	assertfatal.EqualError(t, err, nil)
	var (
		ver3Type = reflect.TypeFor[types.Ver3]()
		ver4Type = reflect.TypeFor[types.Ver4]()
	)
	err = g.RegisterVersioned(reflect.TypeFor[types.VersionedRegister](),
		veropts.WithVersion(ver3Type, "MigrateVer3"),
		veropts.WithCurrentVersion(ver4Type),
	)
	assertfatal.EqualError(t, err, nil)

	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/versioned_register_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
