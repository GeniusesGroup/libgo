/* For license and copyright information please see the LEGAL file in the code repository */

package reference_p

import (
	capsule_p "memar/computer/capsule/protocol"
	error_p "memar/error/protocol"
)

type Reference[T any] interface {
	// Returns a copy of the type itself
	Dereference() (t T, err error_p.Error)

	// TODO::: we need out of scope mechanism like RUST, how to implement this in the library instead of language syntax??
	capsule_p.LifeCycle
}
