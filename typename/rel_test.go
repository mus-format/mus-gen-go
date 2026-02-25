package typename

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestRelName(t *testing.T) {

	t.Run("WithoutSquares", func(t *testing.T) {
		testCases := []struct {
			name    RelName
			wantStr string
		}{
			{
				name:    RelName("pkg.Type"),
				wantStr: "pkg.Type",
			},
			{
				name:    RelName("pkg.Type[Param1, Param2]"),
				wantStr: "pkg.Type",
			},
		}
		for _, c := range testCases {
			str := c.name.WithoutSquares()
			asserterror.Equal(t, str, c.wantStr)
		}
	})

}
