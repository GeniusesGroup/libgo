/* For license and copyright information please see the LEGAL file in the code repository */

package earth

import (
	"memar/math/integer"
	"memar/time/duration"
)

// Elapsed types specified time elapsed of January 1 of the absolute year.
// January 1 of the absolute year(0001), like January 1 of 2001, was a Monday.
type (
	Year         integer.S64
	CalendarYear integer.S64
	TropicalYear integer.S64
)

// Common durations.
const (
	NanoSecondInYear duration.NanoSecond = 31556926 * duration.NanoSecondInSecond
	SecondInYear     duration.Second     = 31556926 // 365.24 days

	// TropicalYear also known as a solar year - https://en.wikipedia.org/wiki/Tropical_year
	SecondInTropicalYear duration.Second = (365 * 24 * 60 * 60) + (5 * 60 * 60) + (48 * 60) + 46 // 365.24219 * 24 * 60 * 60 = 31,556,925.216

	DayIn4Years   Day = 365*4 + 1
	DayIn100Years Day = 365*100 + 24
	DayIn400Years Day = 365*400 + 97
)
