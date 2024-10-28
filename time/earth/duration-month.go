/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/math/integer"
	"memar/time/duration"
)

type Month integer.S64

// Common durations.
const (
	NanoSecondInMonth duration.NanoSecond = 2629743 * duration.NanoSecondInSecond
	SecondInMonth     duration.Second     = 2629743 // 30.44 days
)

func (m *Month) FromNanoSecond(d duration.NanoSecond) {
	// TODO::: any bad situation?
	*m = Month(d / NanoSecondInMonth)
}

func (m *Month) FromSecond(d duration.Second) {
	// TODO::: any bad situation?
	*m = Month(d / SecondInMonth)
}
