/* For license and copyright information please see the LEGAL file in the code repository */

package unix

import (
	error_p "memar/error/protocol"
	string_p "memar/string/protocol"
	"memar/time/duration"
	"memar/time/earth"
	time_p "memar/time/protocol"
)

func Now() (t Time) { t.Now(); return }

// A Time specifies second elapsed of January 1 of the absolute year.
// January 1 of the absolute year(1970), like January 1 of 2001, was a Monday.
type Time struct {
	sec  duration.Second
	nsec duration.NanoInSecond
}

//memar:impl memar/time/protocol.Time
func (t *Time) Epoch() time_p.Epoch                        { return &Epoch }
func (t *Time) SecondElapsed() duration.Second             { return t.sec }
func (t *Time) NanoInSecondElapsed() duration.NanoInSecond { return t.nsec }

//memar:impl memar/string/protocol.Stringer
func (t *Time) ToString() (str string_p.String, err error_p.Error) {
	// TODO:::
	return
}
func (t *Time) FromString(str string_p.String) (err error_p.Error) {
	// TODO:::
	return
}

func (t *Time) ChangeTo(sec duration.Second, nsecElapsed duration.NanoInSecond) {
	t.sec, t.nsec = sec, nsecElapsed
}
func (t *Time) Now() {
	var sec, nsec, _ = now()
	t.sec = duration.Second(sec)
	t.nsec = duration.NanoInSecond(nsec)
}
func (t Time) Pass(from Time) (pass bool) {
	if (t.sec > from.sec) || (t.sec == from.sec && t.nsec > from.nsec) {
		pass = true
	}
	return
}
func (t Time) PassNow() bool { return t.Pass(Now()) }
func (t *Time) Add(d duration.NanoSecond) {
	var sec, nsec = d.ToSecAndNano()
	t.sec += sec
	t.nsec += nsec
}
func (t Time) Since(baseTime Time) (d duration.NanoSecond) {
	var secPass = baseTime.sec - t.sec
	var nsecPass = duration.NanoSecond(baseTime.nsec - t.nsec)
	d = secPass.ToNanoSecond() + nsecPass
	return
}
func (t Time) SinceNow() (d duration.NanoSecond) { return t.Since(Now()) }

// fast way: unix.Now().NanoElapsed()	|| Go way: time.Now().UnixNano()
func (t Time) NanoElapsed() (d duration.NanoSecond) {
	d.FromSecAndNano(t.sec, t.nsec)
	return
}

// fast way: unix.Now().MicroElapsed()	|| Go way: time.Now().UnixMicro()
func (t Time) MicroElapsed() (d duration.MicroSecond) {
	d.FromSecAndNano(t.sec, t.nsec)
	return
}

// fast way: unix.Now().MilliElapsed()	|| Go way: time.Now().UnixMilli()
func (t Time) MilliElapsed() (d duration.MilliSecond) {
	d.FromSecAndNano(t.sec, t.nsec)
	return
}

// fast way: unix.Now().SecElapsed()		|| Go way: time.Now().Unix()
func (t Time) SecElapsed() duration.Second { return t.sec }
func (t Time) HourElapsed() (h earth.Hour) { h.FromSecond(t.sec); return }
func (t Time) DayElapsed() (day earth.Day) { day.FromSecond(t.sec); return }

// ElapsedByNanoSecond returns the result of rounding t to the nearest multiple of d.
func (t Time) ElapsedByNanoSecond(d duration.NanoSecond) (period int64) {
	if d < 0 {
		return
	}
	if d == 0 {
		return int64(t.NanoElapsed())
	}
	var sec, nsec = d.ToSecAndNano()
	if sec > 0 {
		period = int64(t.sec / sec)
		if nsec > 0 {
			period += (int64(t.nsec/nsec) / int64(duration.NanoSecondInSecond))
		}
	} else {
		period = int64(duration.NanoSecond(t.sec)*duration.NanoSecondInSecond) / int64(nsec)
		period += int64(t.nsec / nsec)
	}
	return
}

// ElapsedBySecond returns the result of rounding t to the nearest multiple of d.
func (t Time) ElapsedBySecond(d duration.Second) (period int64) {
	if d < 0 {
		return
	}
	if d == 0 {
		return int64(t.NanoElapsed())
	}

	period = int64(t.sec / d)
	return
}
