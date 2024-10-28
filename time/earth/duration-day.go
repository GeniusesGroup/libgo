/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/math/integer"
	"memar/time/duration"
)

// Day show number of days pass from specific epoch.
// e.g. `unix.Now().DayElapsed()` return number of days from unix epoch.
type Day integer.S64

// Common day durations
const (
	NanoSecondInDay duration.NanoSecond = 24 * NanoSecondInHour
	SecondInDay     duration.Second     = 24 * SecondInHour
	MinuteInDay     Minute              = 24 * MinuteInHour
	HourInDay       Hour                = 24
)

func (day *Day) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*day = Day(d / NanoSecondInDay)
}

func (day *Day) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*day = Day(d / SecondInDay)
}
