/* For license and copyright information please see the LEGAL file in the code repository */

package capsule_p

import (
	error_p "memar/error/protocol"
)

type Accessor[T any] interface {
	Get() T

	// It will check(validate) given value and return proper error for
	Set(new T) (err error_p.Error)
}

type AtomicAccessor[T any] interface {
	Load() T
	Store(new T) (err error_p.Error)
	Swap(new T) (old T, err error_p.Error)
	CompareAndSwap(old, new T) (err error_p.Error) // return more than swapped(bool)
}

type DefaultValue[T any] interface {
	Default() T
	SetDefault() // default value
}

type Optional interface {
	// Base on data or function false means:
	// - data required and must be exist.
	// 		e.g. user gender is optional, user email is not optional
	// - Function (usually a service) MUST be successful in chain of calls or MUST rollback all calls.
	// 		e.g. SendSMS is optional but WithdrawMoney is not optional in TransferMoney chain
	// 		SendSMS is not optional in SendOTP chain
	Optional() bool
}
