package tpopts

import "fmt"

const (
	// TimeUnitUndefined represents an undefined time unit.
	TimeUnitUndefined TimeUnit = iota
	// TimeUnitSec represents seconds.
	TimeUnitSec
	// TimeUnitMilli represents milliseconds.
	TimeUnitMilli
	// TimeUnitMicro represents microseconds.
	TimeUnitMicro
	// TimeUnitNano represents nanoseconds.
	TimeUnitNano
	// TimeUnitSecUTC represents seconds in UTC.
	TimeUnitSecUTC
	// TimeUnitMilliUTC represents milliseconds in UTC.
	TimeUnitMilliUTC
	// TimeUnitMicroUTC represents microseconds in UTC.
	TimeUnitMicroUTC
	// TimeUnitNanoUTC represents nanoseconds in UTC.
	TimeUnitNanoUTC
)

// TimeUnit represents the time unit used for time.Time values.
type TimeUnit int

// Ser returns the name of the serialization function for the time unit.
func (u TimeUnit) Ser() string {
	switch u {
	case TimeUnitSec:
		return "TimeUnix"
	case TimeUnitMilli:
		return "TimeUnixMilli"
	case TimeUnitMicro:
		return "TimeUnixMicro"
	case TimeUnitNano:
		return "TimeUnixNano"
	case TimeUnitSecUTC:
		return "TimeUnixUTC"
	case TimeUnitMilliUTC:
		return "TimeUnixMilliUTC"
	case TimeUnitMicroUTC:
		return "TimeUnixMicroUTC"
	case TimeUnitNanoUTC:
		return "TimeUnixNanoUTC"
	default:
		panic(fmt.Sprintf("unexpected TimeUnit %v", u))
	}
}
