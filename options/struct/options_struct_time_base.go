package stopts

import tpopts "github.com/mus-format/mus-gen-go/options/type"

// UnderlyingTimeOptions specifies configuration options for a time.Time
// underlying type.
type UnderlyingTimeOptions struct {
	TimeUnit tpopts.TimeUnit
}

// SetUnderlyingTimeOption defines a function type that applies a configuration
// option to an UnderlyingTimeOptions struct.
type SetUnderlyingTimeOption func(o *UnderlyingTimeOptions)

// WithUnderlyingTimeUnit returns a SetUnderlyingTimeOption that sets the
// time unit for the underlying time type.
func WithUnderlyingTimeUnit(timeUnit tpopts.TimeUnit) SetUnderlyingTimeOption {
	return func(o *UnderlyingTimeOptions) {
		o.TimeUnit = timeUnit
	}
}

// UnderlyingTimeApply applies a set of options to the provided
// UnderlyingTimeOptions struct.
func UnderlyingTimeApply(opts []SetUnderlyingTimeOption, o *UnderlyingTimeOptions) {
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
}
