// Package intropts provides options for interfaces.
package intropts

import (
	"reflect"
)

// Options specifies configuration options for an interface.
type Options struct {
	Impls      []reflect.Type
	Marshaller bool
}

// SetOption defines a function type that applies a configuration option to an
// Options struct.
type SetOption func(o *Options)

// WithImpl returns a SetOption that registers a single implementation type for
// the interface.
func WithImpl(impl reflect.Type) SetOption {
	return WithImpls(impl)
}

// WithMarshaller returns a SetOption that enables marshaller code generation
// for the interface.
func WithMarshaller() SetOption {
	return func(o *Options) {
		o.Marshaller = true
	}
}

// WithImpls returns a SetOption that registers multiple implementation types
// for the interface.
func WithImpls(impls ...reflect.Type) SetOption {
	return func(o *Options) {
		o.Impls = append(o.Impls, impls...)
	}
}

// Apply applies a set of options to the provided Options struct.
func Apply(o *Options, opts ...SetOption) {
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
}
