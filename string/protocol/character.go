/* For license and copyright information please see the LEGAL file in the code repository */

package string_p

import (
	primitive_p "memar/computer/language/primitive/protocol"
)

// Character is base of any character encoding
type Character interface {
	primitive_p.Equivalence[Character]
	primitive_p.Clone[Character]
	primitive_p.Copy[Character]
}
