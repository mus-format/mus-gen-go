package intropts

import (
	"reflect"
	"testing"

	stopts "github.com/mus-format/mus-gen-go/options/struct"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestNewRegisterOptions(t *testing.T) {
	o := NewRegisterOptions()
	asserterror.Equal(t, o.StructImpls != nil, true)
	asserterror.Equal(t, o.DefinedTypeImpls != nil, true)
	asserterror.Equal(t, len(o.StructImpls), 0)
	asserterror.Equal(t, len(o.DefinedTypeImpls), 0)
}

func TestWithStructImpl(t *testing.T) {
	o := NewRegisterOptions()
	typ := reflect.TypeFor[struct{}]()
	opts := []stopts.SetOption{nil}
	WithStructImpl(typ, opts...)(&o)

	asserterror.Equal(t, len(o.StructImpls), 1)
	asserterror.Equal(t, o.StructImpls[0].Type, typ)
	asserterror.Equal(t, len(o.StructImpls[0].Opts), 1)
}

func TestWithDefinedTypeImpl(t *testing.T) {
	o := NewRegisterOptions()
	typ := reflect.TypeFor[int]()
	opts := []tpopts.SetOption{tpopts.WithNumEncoding(tpopts.NumEncodingVarint)}
	WithDefinedTypeImpl(typ, opts...)(&o)

	asserterror.Equal(t, len(o.DefinedTypeImpls), 1)
	asserterror.Equal(t, o.DefinedTypeImpls[0].Type, typ)
	asserterror.Equal(t, len(o.DefinedTypeImpls[0].Opts), 1)
}

func TestWithRegisterMarshaller(t *testing.T) {
	o := NewRegisterOptions()
	WithRegisterMarshaller()(&o)
	asserterror.Equal(t, o.Marshaller, true)
}

func TestApplyRegister(t *testing.T) {
	o := NewRegisterOptions()
	ApplyRegister(&o,
		WithStructImpl(reflect.TypeFor[struct{}]()),
		WithRegisterMarshaller(),
		nil,
	)
	asserterror.Equal(t, len(o.StructImpls), 1)
	asserterror.Equal(t, o.Marshaller, true)
}
