package typeops

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestOptions(t *testing.T) {
	var (
		o                = Options{}
		wantIgnore       = true
		wantNumEncoding  = Raw
		wantSourceType   = Time
		wantTimeUint     = Milli
		wantValidator    = "ValidateType"
		wantLenEncoding  = Varint
		wantLenValidator = "ValidateLength"
		wantKey          = &Options{Ignore: true}
		wantElem         = &Options{Ignore: true}
	)
	Apply([]SetOption{
		WithIgnore(),
		WithNumEncoding(Raw),
		WithSourceType(Time),
		WithTimeUnit(Milli),
		WithValidator("ValidateType"),
		WithLenEncoding(Varint),
		WithLenValidator("ValidateLength"),
		WithKey(WithIgnore()),
		WithElem(WithIgnore()),
	}, &o)
	asserterror.EqualDeep(t, o.Ignore, wantIgnore)
	asserterror.Equal(t, o.NumEncoding, wantNumEncoding)
	asserterror.Equal(t, o.SourceType, wantSourceType)
	asserterror.Equal(t, o.TimeUnit, wantTimeUint)
	asserterror.Equal(t, o.Validator, wantValidator)
	asserterror.Equal(t, o.LenEncoding, wantLenEncoding)
	asserterror.Equal(t, o.LenValidator, wantLenValidator)
	asserterror.Equal(t, *o.Key, *wantKey)
	asserterror.Equal(t, *o.Elem, *wantElem)

	t.Run("Hash", func(t *testing.T) {
		o1 := Options{}
		Apply([]SetOption{
			WithIgnore(),
			WithNumEncoding(Raw),
			WithSourceType(Time),
			WithTimeUnit(Milli),
			WithValidator("ValidateType"),
			WithLenEncoding(Varint),
			WithLenValidator("ValidateLength"),
			WithKey(WithIgnore()),
			WithElem(WithIgnore()),
		}, &o1)
		o2 := Options{}
		Apply([]SetOption{
			WithIgnore(),
			WithNumEncoding(Raw),
			WithSourceType(Time),
			WithTimeUnit(Milli),
			WithValidator("ValidateType"),
			WithLenEncoding(Varint),
			WithLenValidator("ValidateLength"),
			WithKey(WithIgnore()),
			WithElem(WithIgnore()),
		}, &o2)
		asserterror.Equal(t, o1.Hash(), o2.Hash())
	})
}
