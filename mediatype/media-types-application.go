/* For license and copyright information please see LEGAL file in repository */

package mediatype

var (
	WASM       = newMediaType("", "application/wasm", "wasm", "WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine")
	ARC        = newMediaType("", "application/octet-stream", "arc", "Archive document (multiple files embedded)")
	JavaScript = newMediaType("", "application/javascript", "js", "JavaScript (ECMAScript) programming language")
	JSON       = newMediaType("", "application/json", "json", "JSON (JavaScript Object Notation) format")
	XML        = newMediaType("", "application/xml", "xml", "XML format")
	XHTML      = newMediaType("", "application/xhtml+xml", "xhtml", "")
	PDF        = newMediaType("", "application/pdf", "pdf", "Adobe Portable Document Format")
	GZ         = newMediaType("", "application/gzip", "gz", "strictly speaking under MIME gzip would only be used as an encoding, not a content-type, but it's common to have .gz files")
	TAR        = newMediaType("", "application/tar", "tar", "Tape Archive")
	ABW        = newMediaType("", "application/x-abiword", "abw", "AbiWord document")
	AZW        = newMediaType("", "application/vnd.amazon.ebook", "azw", "Amazon Kindle eBook format")
	BIN        = newMediaType("", "application/octet-stream", "bin", "Any kind of binary data")
	BZ         = newMediaType("", "application/x-bzip", "bz", "BZip archive")
	BZ2        = newMediaType("", "application/x-bzip2", "bz2", "BZip2 archive")
	CSH        = newMediaType("", "application/x-csh", "csh", "C-Shell script")
	DOC        = newMediaType("", "application/msword", "doc", "Microsoft Word")
	EPUB       = newMediaType("", "application/epub+zip", "epub", "Electronic publication")
	MPKG       = newMediaType("", "application/vnd.apple.installer+xml", "mpkg", "Apple Installer Package")
	ODP        = newMediaType("", "application/vnd.oasis.opendocument.presentation", "odp", "OpenDocuemnt presentation document")
	ODS        = newMediaType("", "application/vnd.oasis.opendocument.spreadsheet", "ods", "OpenDocuemnt spreadsheet document")
	ODT        = newMediaType("", "application/vnd.oasis.opendocument.text", "odt", "OpenDocument text document")
	OGX        = newMediaType("", "application/ogg", "ogx", "OGG")
	PPT        = newMediaType("", "application/vnd.ms-powerpoint", "ppt", "Microsoft PowerPoint")
	RAR        = newMediaType("", "application/x-rar-compressed", "rar", "RAR archive")
	RTF        = newMediaType("", "application/rtf", "rtf", "Rich Text Format")
	SH         = newMediaType("", "application/x-sh", "sh", "Bourne shell script")
	SWF        = newMediaType("", "application/x-shockwave-flash", "swf", "Small web format (SWF) or Adobe Flash document")
	ZIP        = newMediaType("", "application/zip", "zip", "ZIP archive")
	JAR        = newMediaType("", "application/java-archive", "jar", "Java Archive")
	VSD        = newMediaType("", "application/vnd.visio", "vsd", "Microsoft Visio")
	XLS        = newMediaType("", "application/vnd.ms-excel", "xls", "Microsoft Excel")
	XUL        = newMediaType("", "application/vnd.mozilla.xul+xml", "xul", "")
	SevenZ     = newMediaType("", "application/x-7z-compressed", "7z", "7-zip archive")
)
