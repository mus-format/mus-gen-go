package stopts

import (
	"testing"

	fldopts "github.com/mus-format/mus-gen-go/options/field"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestWithValidator(t *testing.T) {
	opts := Options{}
	validator := "Validator"
	WithValidator(validator)(&opts)
	asserterror.Equal(t, opts.Validator, validator)
}

func TestWithField(t *testing.T) {
	opts := Options{}
	WithField(fldopts.WithIgnore())(&opts)
	asserterror.Equal(t, len(opts.Fields), 1)
	asserterror.Equal(t, opts.Fields[0].Ignore, true)
}

func TestWithUnderlyingTime(t *testing.T) {
	opts := Options{}
	WithUnderlyingTime(WithUnderlyingTimeUnit(tpopts.TimeUnitSec))(&opts)
	asserterror.Equal(t, len(opts.UnderlyingTime), 1)
}

func TestApply(t *testing.T) {
	opts := Options{}
	Apply([]SetOption{
		WithValidator("Validator"),
		WithField(fldopts.WithIgnore()),
		nil,
	}, &opts)
	asserterror.Equal(t, opts.Validator, "Validator")
	asserterror.Equal(t, len(opts.Fields), 1)
}
