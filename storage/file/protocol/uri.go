/* For license and copyright information please see the LEGAL file in the code repository */

package file_p

import (
	uri_p "memar/net/uri/protocol"
	string_p "memar/string/protocol"
)

// https://datatracker.ietf.org/doc/html/rfc8089
// https://en.wikipedia.org/wiki/File_URI_scheme
type URI[STR string_p.String] interface {
	// URI[STR] // always return full file uri as "{{file}}://{{domain}}/{{path/to/{{{{the file}}.{{html}}}}}}"
	uri_p.Scheme[STR] // always return "file"
	Domain() string   // same as URI.Authority
	// File location from root directory include file name if not point to a directory.
	uri_p.Path[STR]

	FileName[STR]
	FileExtension

	// Helper functions
	IsDirectory() bool // Path end with "/" if it is a file directory
}

type FileName[STR string_p.String] interface {
	FileName() STR // Full name with extension if exist
	FileNameWithoutExtension() STR
}

// A filename extension, file name extension or file extension is
// a suffix to the name of a computer file (for example, .txt, .docx, .md).
// The extension indicates a characteristic of the file contents or its intended use.
// A filename extension is typically delimited from the rest of the filename with a period,
// but in some systems it is separated with spaces.
// https://en.wikipedia.org/wiki/Filename_extension
type FileExtension interface {
	// FileExtension return extension name without a period or space or ...
	FileExtension() string
}
