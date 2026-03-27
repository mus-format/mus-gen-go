package tpopts

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestNumEncodingPackage(t *testing.T) {
	asserterror.Equal(t, NumEncodingVarint.Package(), "varint")
	asserterror.Equal(t, NumEncodingVarintPositive.Package(), "varint")
	asserterror.Equal(t, NumEncodingRaw.Package(), "raw")
	asserterror.Equal(t, NumEncodingUndefined.Package(), "varint")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Package() should panic for invalid NumEncoding")
		}
	}()
	NumEncoding(99).Package()
}

func TestNumEncodingLenSer(t *testing.T) {
	asserterror.Equal(t, NumEncodingVarint.LenSer(), "varint.Int")
	asserterror.Equal(t, NumEncodingVarintPositive.LenSer(), "varint.PositiveInt")
	asserterror.Equal(t, NumEncodingRaw.LenSer(), "raw.Int")
	asserterror.Equal(t, NumEncodingUndefined.LenSer(), "varint.PositiveInt")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("LenSer() should panic for invalid NumEncoding")
		}
	}()
	NumEncoding(99).LenSer()
}
