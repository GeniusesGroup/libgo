/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/math/integer"
	"memar/time/duration"
)

type Minute integer.S64

// Common durations.
const (
	NanoSecondInMinute duration.NanoSecond = 60 * duration.NanoSecondInSecond
	SecondInMinute     duration.Second     = 60
)

func (m *Minute) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*m = Minute(d / NanoSecondInMinute)
}

func (m *Minute) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*m = Minute(d / SecondInMinute)
}
