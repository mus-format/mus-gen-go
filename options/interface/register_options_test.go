package introps

import (
	"reflect"
	"testing"

	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	prim_testdata "github.com/mus-format/musgen-go/testutil/primitive"
	struct_testdata "github.com/mus-format/musgen-go/testutil/struct"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestRegisterOptions(t *testing.T) {
	var (
		o               = NewRegisterOptions()
		wantStructImpls = []StructImpl{
			{
				Type: reflect.TypeFor[struct_testdata.MyStruct](),
				Ops:  []structops.SetOption{},
			},
			{
				Type: reflect.TypeFor[struct_testdata.ComplexStruct](),
				Ops:  []structops.SetOption{},
			},
		}
		wantDefinedTypeImpls = []DefinedTypeImpl{
			{
				Type: reflect.TypeFor[prim_testdata.MyInt](),
				Ops:  []typeops.SetOption{},
			},
			{
				Type: reflect.TypeFor[prim_testdata.MyInt](),
				Ops:  []typeops.SetOption{},
			},
		}
		wantMarshaller = true
	)
	ApplyRegister([]SetRegisterOption{
		WithStructImpl(wantStructImpls[0].Type, wantStructImpls[0].Ops...),
		WithStructImpl(wantStructImpls[1].Type, wantStructImpls[1].Ops...),
		WithDefinedTypeImpl(wantDefinedTypeImpls[0].Type, wantDefinedTypeImpls[0].Ops...),
		WithDefinedTypeImpl(wantDefinedTypeImpls[1].Type, wantDefinedTypeImpls[1].Ops...),
		WithRegisterMarshaller(),
	}, &o)
	asserterror.EqualDeep(t, o.StructImpls, wantStructImpls)
	asserterror.EqualDeep(t, o.DefinedTypeImpls, wantDefinedTypeImpls)
	asserterror.Equal(t, o.Marshaller, wantMarshaller)
}
