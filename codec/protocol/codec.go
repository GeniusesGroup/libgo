/* For license and copyright information please see the LEGAL file in the code repository */

package codec_p

import (
	container_p "memar/adt/container/protocol"
	buffer_p "memar/buffer/protocol"
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
)

// Codec wraps some other interfaces that need an data structure be a codec.
// Protocols that have just one specific logic NEED to implement this interface e.g. HTTP, MP3, AVI, ...
// Others can implement other Codec e.g. Syllab, JSON, XML, HTML, CSS, ...
// https://en.wikipedia.org/wiki/Codec
type Codec /*[BUF Buffer]*/ interface {
	Decoder /*[BUF]*/
	Encoder /*[BUF]*/

	datatype_p.DataType
}

// Decoder is the interface that wraps the Decode method.
type Decoder /*[BUF Buffer]*/ interface {
	// Decode read and decode data until end of needed data or occur error.
	// Unlike io.ReadFrom() it isn't read until EOF and just read needed data.
	Decode(source buffer_p.Buffer) (err error_p.Error)
}

// Encoder is the interface that wraps the Encode & CodecLength methods.
type Encoder /*[BUF Buffer]*/ interface {
	// Encode writes serialized(encoded) data to destination until there's no more data to write.
	// Return any error that occur in buffer logic e.g. timeout error in socket, ...
	Encode(destination buffer_p.Buffer) (err error_p.Error)

	CodecLength
}

type CodecLength interface {
	// SerializationLength return value ln, that is the max number of bytes that will written as encode data by Encode()||Marshal()
	// 0 means no data and -1 means can't tell until full write.
	// Due to prevent performance penalty, Implementors can return max number instead of actual number of length.
	SerializationLength() (ln container_p.NumberOfElement)
}
