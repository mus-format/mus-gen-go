package tpopts

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestTimeUnitSer(t *testing.T) {
	asserterror.Equal(t, TimeUnitSec.Ser(), "TimeUnix")
	asserterror.Equal(t, TimeUnitMilli.Ser(), "TimeUnixMilli")
	asserterror.Equal(t, TimeUnitMicro.Ser(), "TimeUnixMicro")
	asserterror.Equal(t, TimeUnitNano.Ser(), "TimeUnixNano")
	asserterror.Equal(t, TimeUnitSecUTC.Ser(), "TimeUnixUTC")
	asserterror.Equal(t, TimeUnitMilliUTC.Ser(), "TimeUnixMilliUTC")
	asserterror.Equal(t, TimeUnitMicroUTC.Ser(), "TimeUnixMicroUTC")
	asserterror.Equal(t, TimeUnitNanoUTC.Ser(), "TimeUnixNanoUTC")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Ser() should panic for invalid TimeUnit")
		}
	}()
	TimeUnit(99).Ser()
}
