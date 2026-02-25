package typeops

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestNumEncoding(t *testing.T) {

	t.Run("Package", func(t *testing.T) {
		testCases := []struct {
			e       NumEncoding
			wantPkg string
		}{
			{
				e:       Raw,
				wantPkg: "raw",
			},
			{
				e:       Varint,
				wantPkg: "varint",
			},
			{
				e:       VarintPositive,
				wantPkg: "varint",
			},
		}

		for _, c := range testCases {
			asserterror.Equal(t, c.e.Package(), c.wantPkg)
		}
	})

	t.Run("LenSer", func(t *testing.T) {
		testCases := []struct {
			e          NumEncoding
			wantLenSer string
		}{
			{
				e:          Raw,
				wantLenSer: "raw.Int",
			},
			{
				e:          Varint,
				wantLenSer: "varint.Int",
			},
			{
				e:          UndefinedNumEncoding,
				wantLenSer: "varint.PositiveInt",
			},
			{
				e:          VarintPositive,
				wantLenSer: "varint.PositiveInt",
			},
		}

		for _, c := range testCases {
			asserterror.Equal(t, c.e.LenSer(), c.wantLenSer)
		}
	})

}
