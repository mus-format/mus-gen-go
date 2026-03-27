package tpopts

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestWithOptions(t *testing.T) {
	o := Options{}
	WithNumEncoding(NumEncodingVarint)(&o)
	asserterror.Equal(t, o.NumEncoding, NumEncodingVarint)

	WithTimeUnit(TimeUnitSec)(&o)
	asserterror.Equal(t, o.TimeUnit, TimeUnitSec)

	WithValidator("Validator")(&o)
	asserterror.Equal(t, o.Validator, "Validator")

	WithLenEncoding(NumEncodingRaw)(&o)
	asserterror.Equal(t, o.LenEnc, NumEncodingRaw)

	WithLenValidator("LenValidator")(&o)
	asserterror.Equal(t, o.LenValidator, "LenValidator")

	WithElemValidator("ElemValidator")(&o)
	asserterror.Equal(t, o.ElemValidator, "ElemValidator")

	WithKeyValidator("KeyValidator")(&o)
	asserterror.Equal(t, o.KeyValidator, "KeyValidator")
}

func TestApply(t *testing.T) {
	o := Options{}
	Apply(&o,
		WithNumEncoding(NumEncodingVarint),
		WithTimeUnit(TimeUnitSec),
		nil,
	)
	asserterror.Equal(t, o.NumEncoding, NumEncodingVarint)
	asserterror.Equal(t, o.TimeUnit, TimeUnitSec)
}
