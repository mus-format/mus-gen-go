package testutil

import (
	"testing"

	"github.com/mus-format/mus-go"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestUnmarshalError[T any](v T, wantErr error, ser mus.Serializer[T],
	t *testing.T) {
	t.Helper()
	bs := make([]byte, ser.Size(v))
	ser.Marshal(v, bs)
	_, _, err := ser.Unmarshal(bs)
	asserterror.EqualError(t, err, wantErr)
}
