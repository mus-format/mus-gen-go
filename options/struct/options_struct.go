// Package stopts provides options for structs.
package stopts

import (
	fldopts "github.com/mus-format/mus-gen-go/options/field"
)

// Options specifies configuration options for a struct.
type Options struct {
	Validator      string
	Fields         []fldopts.Options
	UnderlyingTime []SetUnderlyingTimeOption
}

// SetOption defines a function type that applies a configuration option to an
// Options struct.
type SetOption func(o *Options)

// WithValidator returns a SetOption that sets the validator for the struct.
func WithValidator(validator string) SetOption {
	return func(o *Options) {
		o.Validator = validator
	}
}

// WithField returns a SetOption that adds field options to the struct.
func WithField(opts ...fldopts.SetOption) SetOption {
	return func(o *Options) {
		fopts := fldopts.Options{}
		fldopts.Apply(&fopts, opts...)
		o.Fields = append(o.Fields, fopts)
	}
}

// WithUnderlyingTime returns a SetOption that adds underlying time options to
// the struct.
func WithUnderlyingTime(opts ...SetUnderlyingTimeOption) SetOption {
	return func(o *Options) {
		if o.UnderlyingTime == nil {
			o.UnderlyingTime = make([]SetUnderlyingTimeOption, 0)
		}
		o.UnderlyingTime = append(o.UnderlyingTime, opts...)
	}
}

// Apply applies a set of options to the provided Options struct.
func Apply(opts []SetOption, o *Options) {
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
}
