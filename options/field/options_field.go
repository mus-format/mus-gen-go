// Package fldopts provides options for struct fields.
package fldopts

import tpopts "github.com/mus-format/mus-gen-go/options/type"

// Options specifies configuration options for a struct field.
type Options struct {
	Type   tpopts.Options
	Ignore bool
}

// SetOption defines a function type to set options on an Options struct.
type SetOption func(o *Options)

// WithIgnore returns a SetOption that marks a field to be ignored during
// serialization.
func WithIgnore() SetOption {
	return func(o *Options) {
		o.Ignore = true
	}
}

// WithType returns a SetOption that applies the given type options to the
// field.
func WithType(opts ...tpopts.SetOption) SetOption {
	return func(o *Options) {
		tpopts.Apply(&o.Type, opts...)
	}
}

// Apply applies the given set of options to the provided Options struct.
func Apply(o *Options, opts ...SetOption) {
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
}
