/* For license and copyright information please see the LEGAL file in the code repository */

package monotonic

import (
	error_p "memar/error/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
)

// Now returns runtime monotonic clock in nanoseconds.
func Now() Time {
	return Time(now())
}

// A Time monotonic clock is for measuring time.
// time-measuring operations, specifically comparisons and subtractions, use the monotonic clock.
type Time duration.NanoSecond

//memar:impl memar/time/protocol.Time
func (t *Time) Epoch() time_p.Epoch { return &Epoch }
func (t *Time) SecondElapsed() (sec duration.Second) {
	sec.FromNanoSecond(duration.NanoSecond(*t))
	return
}
func (t *Time) NanoInSecondElapsed() (d duration.NanoInSecond) {
	d.FromSec(duration.NanoSecond(*t))
	return
}

//memar:impl memar/string/protocol.Stringer
func (t *Time) ToString() (s string, err error_p.Error) {
	s = "TODO:::"
	return
}

//memar:impl memar/string/protocol.Stringer
func (t *Time) FromString(str string) (err error_p.Error) {
	// TODO:::
	return
}

func (t *Time) Now()                      { *t = Now() }
func (t *Time) Add(d duration.NanoSecond) { *t += Time(d) }

// Equal reports whether t and other represent the same time instant.
func (t Time) Equal(other Time) bool { return t == other }

// Pass reports whether the time instant t is after from.
func (t Time) Pass(from Time) bool { return t > from }

// PassNow reports whether the time instant t is after now.
func (t Time) PassNow() bool { return t > Now() }

// Since returns the time elapsed since t.
func (t Time) Since(from Time) (d duration.NanoSecond) { return duration.NanoSecond(from - t) }

// SinceNow returns the time elapsed since now.
func (t Time) SinceNow() (d duration.NanoSecond) { return duration.NanoSecond(Now() - t) }

// Until returns the duration t(the future) until to (before the t).
func (t Time) Until(to Time) (d duration.NanoSecond) { return duration.NanoSecond(t - to) }

// UntilNow returns the duration t(the future) until now.
func (t Time) UntilNow() (d duration.NanoSecond) { return duration.NanoSecond(t - Now()) }
