package musgen_test

import (
	"os"
	"reflect"
	"testing"

	musgen "github.com/mus-format/mus-gen-go/mus"
	genops "github.com/mus-format/mus-gen-go/options/gen"
	intropts "github.com/mus-format/mus-gen-go/options/interface"
	veropts "github.com/mus-format/mus-gen-go/options/versioned"
	types "github.com/mus-format/mus-gen-go/test/types/mode"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

// Safe ------------------------------------------------------------------------

func TestGenerateMode_Safe(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types/modes/safe"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[types.FullDefined]())
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterInterface(reflect.TypeFor[types.FullInterface](),
		intropts.WithDefinedTypeImpl(reflect.TypeFor[types.FullInterfaceImpl]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterVersioned(reflect.TypeFor[types.Versioned](),
		veropts.WithVersion(reflect.TypeFor[types.FooV1](), "mode.MigrateFooV1"),
		veropts.WithCurrentVersion(reflect.TypeFor[types.FooV2]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(reflect.TypeFor[types.FullStruct]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/mode/safe/safe_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

// Unsafe ----------------------------------------------------------------------

func TestGenerateMode_Unsafe(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types/mode/unsafe"),
		genops.WithUnsafe(),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[types.FullDefined]())
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterInterface(reflect.TypeFor[types.FullInterface](),
		intropts.WithDefinedTypeImpl(reflect.TypeFor[types.FullInterfaceImpl]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterVersioned(reflect.TypeFor[types.Versioned](),
		veropts.WithVersion(reflect.TypeFor[types.FooV1](), "mode.MigrateFooV1"),
		veropts.WithCurrentVersion(reflect.TypeFor[types.FooV2]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(reflect.TypeFor[types.FullStruct]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/mode/unsafe/unsafe_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

// Not Unsafe ------------------------------------------------------------------

func TestGenerateMode_NotUnsafe(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types/mode/notunsafe"),
		genops.WithNotUnsafe(),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[types.FullDefined]())
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterInterface(reflect.TypeFor[types.FullInterface](),
		intropts.WithDefinedTypeImpl(reflect.TypeFor[types.FullInterfaceImpl]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterVersioned(reflect.TypeFor[types.Versioned](),
		veropts.WithVersion(reflect.TypeFor[types.FooV1](), "mode.MigrateFooV1"),
		veropts.WithCurrentVersion(reflect.TypeFor[types.FooV2]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(reflect.TypeFor[types.FullStruct]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/mode/notunsafe/notunsafe_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

// Stream Safe -----------------------------------------------------------------

func TestGenerateMode_StreamSafe(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types/mode/stream_safe"),
		genops.WithStream(),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[types.FullDefined]())
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterInterface(reflect.TypeFor[types.FullInterface](),
		intropts.WithDefinedTypeImpl(reflect.TypeFor[types.FullInterfaceImpl]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterVersioned(reflect.TypeFor[types.Versioned](),
		veropts.WithVersion(reflect.TypeFor[types.FooV1](), "mode.MigrateFooV1"),
		veropts.WithCurrentVersion(reflect.TypeFor[types.FooV2]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(reflect.TypeFor[types.FullStruct]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/mode/stream_safe/stream_safe_mus.gen.go", bs,
		0644)
	assertfatal.EqualError(t, err, nil)
}

// Stream Unsafe -----------------------------------------------------------------

func TestGenerateMode_StreamUnsafe(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types/mode/stream_unsafe"),
		genops.WithStream(),
		genops.WithUnsafe(),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[types.FullDefined]())
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterInterface(reflect.TypeFor[types.FullInterface](),
		intropts.WithDefinedTypeImpl(reflect.TypeFor[types.FullInterfaceImpl]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterVersioned(reflect.TypeFor[types.Versioned](),
		veropts.WithVersion(reflect.TypeFor[types.FooV1](), "mode.MigrateFooV1"),
		veropts.WithCurrentVersion(reflect.TypeFor[types.FooV2]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(reflect.TypeFor[types.FullStruct]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/mode/stream_unsafe/stream_unsafe_mus.gen.go",
		bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

// Stream Not Unsafe -----------------------------------------------------------------

func TestGenerateMode_StreamNotUnsafe(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types/mode/stream_notunsafe"),
		genops.WithStream(),
		genops.WithNotUnsafe(),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[types.FullDefined]())
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterInterface(reflect.TypeFor[types.FullInterface](),
		intropts.WithDefinedTypeImpl(reflect.TypeFor[types.FullInterfaceImpl]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterVersioned(reflect.TypeFor[types.Versioned](),
		veropts.WithVersion(reflect.TypeFor[types.FooV1](), "mode.MigrateFooV1"),
		veropts.WithCurrentVersion(reflect.TypeFor[types.FooV2]()),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(reflect.TypeFor[types.FullStruct]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/mode/stream_notunsafe/stream_notunsafe_mus.gen.go",
		bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
