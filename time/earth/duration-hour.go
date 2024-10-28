/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/math/integer"
	"memar/time/duration"
)

type Hour integer.S64

// Common durations.
const (
	NanoSecondInHour duration.NanoSecond = 60 * NanoSecondInMinute
	SecondInHour     duration.Second     = 60 * SecondInMinute
	MinuteInHour     Minute              = 60
)

func (h *Hour) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*h = Hour(d / NanoSecondInHour)
}

func (h *Hour) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*h = Hour(d / SecondInHour)
}
