/* For license and copyright information please see the LEGAL file in the code repository */

package monotonic

import (
	_ "unsafe" // for go:linkname
)

// TODO::: move assembly logic from go/src/runtime to this package to prevent use unsafe package
// TODO::: Is duration include sleep time or without it??

// now returns the current value of the runtime monotonic clock in nanoseconds.
// It isn't not wall clock, Use in tasks like timeout, ...
//
//go:linkname now runtime.nanotime
func now() int64
