package stopts

import (
	"testing"

	tpopts "github.com/mus-format/mus-gen-go/options/type"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestWithUnderlyingTimeTimeUnit(t *testing.T) {
	opts := UnderlyingTimeOptions{}
	WithUnderlyingTimeUnit(tpopts.TimeUnitSec)(&opts)
	asserterror.Equal(t, opts.TimeUnit, tpopts.TimeUnitSec)
}

func TestUnderlyingTimeApply(t *testing.T) {
	opts := UnderlyingTimeOptions{}
	UnderlyingTimeApply([]SetUnderlyingTimeOption{
		WithUnderlyingTimeUnit(tpopts.TimeUnitSec),
		nil,
	}, &opts)
	asserterror.Equal(t, opts.TimeUnit, tpopts.TimeUnitSec)
}
