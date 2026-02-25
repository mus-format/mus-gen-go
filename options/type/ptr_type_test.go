package typeops

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestPtrType(t *testing.T) {

	t.Run("PackageName", func(t *testing.T) {
		testCases := []struct {
			e               PtrType
			wantPackageName string
		}{
			{
				e:               Undefined,
				wantPackageName: "ord",
			},
			{
				e:               Ordinary,
				wantPackageName: "ord",
			},
			{
				e:               PM,
				wantPackageName: "pm",
			},
		}

		for _, c := range testCases {
			asserterror.Equal(t, c.e.PackageName(), c.wantPackageName)
		}
	})

}
