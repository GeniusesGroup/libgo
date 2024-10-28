/* For license and copyright information please see the LEGAL file in the code repository */

package time_p

import (
	datatype_p "memar/datatype/protocol"
	string_p "memar/string/protocol"
)

// Epoch is the interface that must implement by any time capsule.
// It is base on Epoch and Second terms to work anywhere (in any planet in the universe).
// https://en.wikipedia.org/wiki/Epoch
type Epoch interface {
	datatype_p.DataType
	string_p.Stringer_To[string_p.String]
}
