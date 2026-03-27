package fldopts

import (
	"testing"

	tpopts "github.com/mus-format/mus-gen-go/options/type"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestWithIgnore(t *testing.T) {
	opts := Options{}
	WithIgnore()(&opts)
	asserterror.Equal(t, opts.Ignore, true)
}

func TestWithType(t *testing.T) {
	var (
		opts   = Options{}
		numEnc = tpopts.NumEncodingVarint
	)
	WithType(tpopts.WithNumEncoding(numEnc))(&opts)
	asserterror.Equal(t, opts.Type.NumEncoding, numEnc)
}

func TestApply(t *testing.T) {
	var (
		opts   = Options{}
		numEnc = tpopts.NumEncodingVarint
	)
	Apply(&opts,
		WithIgnore(),
		WithType(tpopts.WithNumEncoding(numEnc)),
		nil,
	)
	asserterror.Equal(t, opts.Ignore, true)
	asserterror.Equal(t, opts.Type.NumEncoding, numEnc)
}
