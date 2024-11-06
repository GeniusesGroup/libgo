/* For license and copyright information please see the LEGAL file in the code repository */

package math_p

// Operations are calculation by mathematical methods.
//
// Inverse operations are operations which reverse or “undo” another operation.
// They are also sometimes referred to as 'opposite operations'.
//
// Implementors can provide more useful methods like
// - Increment() { i++ }
// - Decrement() { i-- }
type Operations[T any] interface {
	// Multiplication is an arithmetic operation that is the inverse of division; the product of two numbers is computed.
	Multiplication(by any)

	Multiple(of T)

	// Addition is the arithmetic operation of summing; calculating the sum of two or more numbers
	// Addition is inverse operation.
	Addition(from T)

	// Subtraction is an arithmetic operation in which the difference between two numbers is calculated
	// Subtraction us inverse operation.
	Subtraction(from T)

	// Division is an arithmetic operation that is the inverse of multiplication; the quotient of two numbers is computed.
	// Implementors MUST care for division by zero.
	Division(by T)

	// comparison??
}
