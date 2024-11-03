/* For license and copyright information please see the LEGAL file in the code repository */

package primitive_p

// https://en.wikipedia.org/wiki/Logical_conjunction
// Logical_Connective_AND
//
// Some other languages:
// - https://doc.rust-lang.org/stable/std/ops/trait.BitAnd.html
type Conjunction[T, RES any] interface {
	Conjunction(with T) (res RES)
	// AND()
}

// https://en.wikipedia.org/wiki/Logical_disjunction
// Logical_Connective_OR
//
// Some other languages:
// - https://doc.rust-lang.org/stable/std/ops/trait.BitOr.html
type Disjunction[T, RES any] interface {
	Disjunction(with T) (res RES)
	// OR()
}

// In logic and mathematics, statements are said to be logically equivalent if they have the same truth value in every model.
// It will use when qualities that are comparable.
// https://en.wikipedia.org/wiki/Logical_equivalence
// Logical_Connective_Equivalent
type Equivalence[T any] interface {
	// The statement P≡Q asserts the logical equivalence of P and Q
	// The statement P⟺Q asserts the (not necessarily logical) equivalence of P and Q.
	// It is not equal(=) sign
	// For example, the sentences 0=0 and 32=9 are mathematically but not logically equivalent.
	Equivalence(with T) bool
	// Equal(with T) bool
	// EQV()
}

// BiConditional is a relation between two propositions that is true only when both propositions are simultaneously true or false
// https://en.wikipedia.org/wiki/Logical_biconditional
// Logical_Connective_Equivalent
type BiConditional interface {
	// BiConditional is represented by ↔
	// This means that that P↔Q  is either analytically (in an axiom system) or synthetically true ,
	// i.e., that P and Q have the same truth value in the given interpretation.
	BiConditional()
	// BIC()
}

// In logic, negation, also called the logical not or logical complement,
// is an operation that takes a proposition to another proposition.
// https://en.wikipedia.org/wiki/Negation
// Logical_Connective_NOT
//
// Some other languages:
// - https://doc.rust-lang.org/stable/std/ops/trait.Not.html
type Negation[T any] interface {
	Negation(p T) (ap T)
	// NOT()
}
