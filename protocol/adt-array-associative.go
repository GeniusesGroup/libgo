/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// In computer science, an associative array, map, symbol table, or dictionary is an abstract data type that
// stores a collection of (key, value) pairs, such that each possible key appears at most once in the collection.
// In mathematical terms, an associative array is a function with finite domain.
// https://en.wikipedia.org/wiki/Associative_array
type Array_Associative[KEY comparable, VALUE any] interface {
	ADT_Put[KEY, VALUE]
	ADT_Lookup[KEY, VALUE]
	ADT_Remove[KEY]

	ADT_Iteration_KV[KEY, VALUE]

	Capacity
	OccupiedLength
}

type ADT_Iteration_KV[K, V any] interface {
	Iterate_KV(startIndex ElementIndex, iterator Iterate_KV[K, V]) (err Error)
}

type Iterate_KV[K, V any] interface {
	// Iterate or traverse
	// In each iteration if err != nil, iteration will be stopped
	Iterate(index ElementIndex, key K, value V) (err Error)
}

type ADT_Put[KEY, VALUE any] interface {
	// Put or Insert()
	Put(key KEY, value VALUE) (err Error)
}

type ADT_Lookup[KEY comparable, VALUE any] interface {
	Lookup(key KEY) (value VALUE, err Error)
}

type ADT_Remove[KEY comparable] interface {
	Remove(key KEY) (err Error)
}
