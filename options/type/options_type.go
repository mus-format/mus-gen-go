// Package tpopts provides options for types.
package tpopts

// Options specifies configuration options for a type.
type Options struct {
	TimeUnit TimeUnit

	NumEncoding   NumEncoding
	Validator     string
	LenEnc        NumEncoding
	LenValidator  string
	ElemValidator string
	KeyValidator  string
}

// SetOption defines a function type that applies a configuration option to an
// Options struct.
type SetOption func(o *Options)

// WithNumEncoding returns a SetOption that sets the numeric encoding for the
// type.
func WithNumEncoding(numEnc NumEncoding) SetOption {
	return func(o *Options) {
		o.NumEncoding = numEnc
	}
}

// WithTimeUnit returns a SetOption that sets the time unit for the type.
func WithTimeUnit(timeUnit TimeUnit) SetOption {
	return func(o *Options) {
		o.TimeUnit = timeUnit
	}
}

// WithValidator returns a SetOption that sets the validator for the type.
func WithValidator(vl string) SetOption {
	return func(o *Options) {
		o.Validator = vl
	}
}

// WithLenEncoding returns a SetOption that sets the length encoding for a type.
func WithLenEncoding(lenEnc NumEncoding) SetOption {
	return func(o *Options) {
		o.LenEnc = lenEnc
	}
}

// WithLenValidator returns a SetOption that sets the length validator for a type.
func WithLenValidator(vl string) SetOption {
	return func(o *Options) {
		o.LenValidator = vl
	}
}

// WithElemValidator returns a SetOption that sets the element validator for a
// container type.
func WithElemValidator(vl string) SetOption {
	return func(o *Options) {
		o.ElemValidator = vl
	}
}

// WithKeyValidator returns a SetOption that sets the key validator for a map
// type.
func WithKeyValidator(v string) SetOption {
	return func(o *Options) {
		o.KeyValidator = v
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
