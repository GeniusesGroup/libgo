/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/binary"
	"libgo/protocol"
)

type optionAltChecksum []byte

func (o optionAltChecksum) Length() byte        { return o[0] }
func (o optionAltChecksum) AltChecksum() uint16 { return binary.BigEndian.Uint16(o[1:]) }
func (o optionAltChecksum) NextOption() []byte  { return o[3:] }

func (o optionAltChecksum) Process(s *Stream) (err protocol.Error) {
	return
}
