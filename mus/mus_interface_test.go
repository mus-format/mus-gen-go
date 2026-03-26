package musgen_test

import (
	"os"
	"reflect"
	"testing"

	musgen "github.com/mus-format/mus-gen-go/mus"
	genops "github.com/mus-format/mus-gen-go/options/gen"
	intropts "github.com/mus-format/mus-gen-go/options/interface"
	types "github.com/mus-format/mus-gen-go/test/types"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestGenerateInterface(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types"),
	)
	assertfatal.EqualError(t, err, nil)
	var (
		impl1Type = reflect.TypeFor[types.Impl1]()
		impl2Type = reflect.TypeFor[types.Impl2]()
	)
	err = g.AddStruct(impl1Type)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(impl2Type)
	assertfatal.EqualError(t, err, nil)
	err = g.AddTyped(impl1Type)
	assertfatal.EqualError(t, err, nil)
	err = g.AddTyped(impl2Type)
	assertfatal.EqualError(t, err, nil)
	err = g.AddInterface(reflect.TypeFor[types.Interface](),
		intropts.WithImpls(impl1Type, impl2Type),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddInterface(reflect.TypeFor[types.MarshallerInterface](),
		intropts.WithImpls(impl1Type, impl2Type),
		intropts.WithMarshaller(),
	)
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/interface_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

func TestGenerateInterface_Register(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/types"),
	)
	assertfatal.EqualError(t, err, nil)
	var (
		impl3Type = reflect.TypeFor[types.Impl3]()
		impl4Type = reflect.TypeFor[types.Impl4]()
		impl5Type = reflect.TypeFor[types.Impl5]()
		impl6Type = reflect.TypeFor[types.Impl6]()
	)
	err = g.RegisterInterface(reflect.TypeFor[types.InterfaceRegister](),
		intropts.WithStructImpl(impl3Type),
		intropts.WithDefinedTypeImpl(impl4Type),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.RegisterInterface(reflect.TypeFor[types.MarshallerInterfaceRegister](),
		intropts.WithStructImpl(impl5Type),
		intropts.WithDefinedTypeImpl(impl6Type),
		intropts.WithRegisterMarshaller(),
	)
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/types/interface_register_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
