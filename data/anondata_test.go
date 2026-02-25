package data

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestAnonSerKind(t *testing.T) {
	testCases := []struct {
		kind    AnonSerKind
		wantStr string
	}{
		{
			kind:    AnonString,
			wantStr: AnonStringPrefix,
		},
		{
			kind:    AnonArray,
			wantStr: AnonArrayPrefix,
		},
		{
			kind:    AnonByteSlice,
			wantStr: AnonByteSlicePrefix,
		},
		{
			kind:    AnonSlice,
			wantStr: AnonSlicePrefix,
		},
		{
			kind:    AnonMap,
			wantStr: AnonMapPrefix,
		},
	}
	for _, c := range testCases {
		asserterror.Equal(t, c.kind.String(), c.wantStr)
	}
}
