package intropts

import (
	"reflect"
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestWithImpl(t *testing.T) {
	opts := Options{}
	typ := reflect.TypeFor[int]()
	WithImpl(typ)(&opts)
	asserterror.Equal(t, len(opts.Impls), 1)
	asserterror.Equal(t, opts.Impls[0], typ)
}

func TestWithImpls(t *testing.T) {
	opts := Options{}
	typs := []reflect.Type{reflect.TypeFor[int](), reflect.TypeFor[string]()}
	WithImpls(typs...)(&opts)
	asserterror.Equal(t, len(opts.Impls), 2)
	asserterror.Equal(t, opts.Impls[0], typs[0])
	asserterror.Equal(t, opts.Impls[1], typs[1])
}

func TestWithMarshaller(t *testing.T) {
	opts := Options{}
	WithMarshaller()(&opts)
	asserterror.Equal(t, opts.Marshaller, true)
}

func TestApply(t *testing.T) {
	opts := Options{}
	Apply(&opts,
		WithImpl(reflect.TypeFor[int]()),
		WithMarshaller(),
		nil,
	)
	asserterror.Equal(t, len(opts.Impls), 1)
	asserterror.Equal(t, opts.Marshaller, true)
}
