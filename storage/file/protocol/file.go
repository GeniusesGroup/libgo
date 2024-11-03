/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	buffer_p "memar/buffer/protocol"
	error_p "memar/error/protocol"
)

// File is the descriptor interface that must implement by any to be an file.
// File owner is one app so it must handle concurrent protection internally not by file it self.
type File interface {
	Metadata() Metadata
	Data() buffer_p.Buffer

	File_Methods
}

type File_Methods interface {
	// Depend on OS, file data can be cache on ram until Save() called.
	Save() (err error_p.Error)

	Rename(newName string)
}
