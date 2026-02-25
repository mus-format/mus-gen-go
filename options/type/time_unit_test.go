package typeops

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestTimeUnit(t *testing.T) {

	t.Run("Ser", func(t *testing.T) {
		testCases := []struct {
			u       TimeUnit
			wantSer string
		}{
			{
				u:       UndefinedTimeUnit,
				wantSer: "TimeUnix",
			},
			{
				u:       Milli,
				wantSer: "TimeUnixMilli",
			},
			{
				u:       Micro,
				wantSer: "TimeUnixMicro",
			},
			{
				u:       Nano,
				wantSer: "TimeUnixNano",
			},
			{
				u:       SecUTC,
				wantSer: "TimeUnixUTC",
			},
			{
				u:       MilliUTC,
				wantSer: "TimeUnixMilliUTC",
			},
			{
				u:       MicroUTC,
				wantSer: "TimeUnixMicroUTC",
			},
			{
				u:       NanoUTC,
				wantSer: "TimeUnixNanoUTC",
			},
		}
		for _, c := range testCases {
			asserterror.Equal(t, c.u.Ser(), c.wantSer)
		}
	})

}
