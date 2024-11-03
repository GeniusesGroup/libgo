/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	container_p "memar/adt/container/protocol"
	string_p "memar/string/protocol"
	time_p "memar/time/protocol"
)

// Metadata is the interface that must implement by any file and directory.
type Metadata interface {
	ParentDirectory() Directory
	URI[string_p.String]
	Size() container_p.NumberOfElement // in Byte??
	Created() time_p.Time
	Accessed() time_p.Time
	Modified() time_p.Time
}
