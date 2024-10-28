/* For license and copyright information please see the LEGAL file in the code repository */

package monotonic

import (
	"sync/atomic"

	error_p "memar/error/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
)

// Atomic same as Time is monotonic clock is for measuring time.
// Just due to 32bit hardwares alignment problem,
// we suggest don't use Time for atomic operation and Use this type for any atomic purposes.
type Atomic struct {
	atomic.Int64
}

func (a *Atomic) Load() Time               { return Time(a.Int64.Load()) }
func (a *Atomic) Store(t Time)             { a.Int64.Store(int64(t)) }
func (a *Atomic) Swap(new Time) (old Time) { return Time(a.Int64.Swap(int64(new))) }
func (a *Atomic) CompareAndSwap(old, new Time) (swapped bool) {
	return a.Int64.CompareAndSwap(int64(old), int64(new))
}
func (a *Atomic) Add(d duration.NanoSecond) { a.Int64.Add(int64(d)) }

//memar:impl memar/time/protocol.Time
func (a *Atomic) Epoch() time_p.Epoch { return &Epoch }
func (a *Atomic) SecondElapsed() duration.Second {
	var t = a.Load()
	return t.SecondElapsed()
}
func (a *Atomic) NanoInSecondElapsed() duration.NanoInSecond {
	var t = a.Load()
	return t.NanoInSecondElapsed()
}

//memar:impl memar/string/protocol.Stringer
func (a *Atomic) ToString() (s string, err error_p.Error) {
	var t = a.Load()
	return t.ToString()
}
func (a *Atomic) FromString(str string) (err error_p.Error) {
	var t Time
	err = t.FromString(str)
	a.Store(t)
	return
}
