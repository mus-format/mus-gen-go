package intropts

import (
	"reflect"

	stopts "github.com/mus-format/mus-gen-go/options/struct"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
)

// NewRegisterOptions creates a new RegisterOptions.
func NewRegisterOptions() RegisterOptions {
	return RegisterOptions{
		StructImpls:      []StructImpl{},
		DefinedTypeImpls: []DefinedTypeImpl{},
	}
}

// RegisterOptions specifies configuration options for registering interface
// implementation types.
type RegisterOptions struct {
	StructImpls      []StructImpl
	DefinedTypeImpls []DefinedTypeImpl
	Marshaller       bool
}

// StructImpl represents a concrete struct implementation and its options.
type StructImpl struct {
	Type reflect.Type
	Opts []stopts.SetOption
}

// DefinedTypeImpl represents a concrete defined type implementation and its
// options.
type DefinedTypeImpl struct {
	Type reflect.Type
	Opts []tpopts.SetOption
}

// SetRegisterOption defines a function type that applies a configuration
// option to a RegisterOptions struct.
type SetRegisterOption func(o *RegisterOptions)

// WithStructImpl returns a SetRegisterOption that registers a concrete
// implementation type for the interface being generated. The provided type
// must implement the target interface.
func WithStructImpl(t reflect.Type, opts ...stopts.SetOption) SetRegisterOption {
	impl := StructImpl{Type: t, Opts: opts}
	return func(o *RegisterOptions) {
		o.StructImpls = append(o.StructImpls, impl)
	}
}

// WithDefinedTypeImpl returns a SetRegisterOption that registers a concrete
// implementation type for the interface being generated. The provided type
// must implement the target interface.
func WithDefinedTypeImpl(t reflect.Type, opts ...tpopts.SetOption) SetRegisterOption {
	impl := DefinedTypeImpl{Type: t, Opts: opts}
	return func(o *RegisterOptions) {
		o.DefinedTypeImpls = append(o.DefinedTypeImpls, impl)
	}
}

// WithRegisterMarshaller returns a SetRegisterOption that enables
// serialization using the Marshaller interface (from the mus-go/mus-stream-go
// modules).
func WithRegisterMarshaller() SetRegisterOption {
	return func(o *RegisterOptions) {
		o.Marshaller = true
	}
}

// ApplyRegister applies a set of register options to the provided
// RegisterOptions struct.
func ApplyRegister(o *RegisterOptions, opts ...SetRegisterOption) {
	for i := range opts {
		if opts[i] != nil {
			opts[i](o)
		}
	}
}
