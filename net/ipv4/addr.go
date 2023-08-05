/* For license and copyright information please see the LEGAL file in the code repository */

package ipv4

import (
	"memar/protocol"
)

// Well-known IPv4 addresses
var (
	AddrZero       = Addr{0, 0, 0, 0}         // all zeros
	AddrAllRouters = Addr{224, 0, 0, 2}       // all routers
	AddrAllSystems = Addr{224, 0, 0, 1}       // all systems
	AddrBroadcast  = Addr{255, 255, 255, 255} // limited broadcast
)

// An Addr is an IP address version 4.
type Addr [AddrLen]byte

//memar:impl memar/protocol.Stringer
func (addr Addr) ToString() string {
	const maxIPv4StringLen = len("255.255.255.255")

	var b = make([]byte, maxIPv4StringLen)

	var n = ubtoa(b, 0, addr[0])
	b[n] = '.'
	n++

	n += ubtoa(b, n, addr[1])
	b[n] = '.'
	n++

	n += ubtoa(b, n, addr[2])
	b[n] = '.'
	n++

	n += ubtoa(b, n, addr[3])
	return string(b[:n])
}
func (addr *Addr) FromString(ipv4 string) (err protocol.Error) {
	for i := 0; i < AddrLen; i++ {
		if len(ipv4) == 0 {
			// err = Missing octets
			return
		}
		if i > 0 {
			if ipv4[0] != '.' {
				// err =
				return
			}
			ipv4 = ipv4[1:]
		}
		n, c, ok := dtoi(ipv4)
		if !ok || n > 0xFF {
			// err =
			return
		}
		if c > 1 && ipv4[0] == '0' {
			// err = Reject non-zero components with leading zeroes.
			return
		}
		ipv4 = ipv4[c:]
		addr[i] = byte(n)
	}
	if len(ipv4) != 0 {
		// Mask??
		// err =
	}
	return
}

// ubtoa encodes the string form of the integer v to dst[start:] and
// returns the number of bytes written to dst. The caller must ensure
// that dst has sufficient length.
func ubtoa(dst []byte, start int, v byte) int {
	if v < 10 {
		dst[start] = v + '0'
		return 1
	} else if v < 100 {
		dst[start+1] = v%10 + '0'
		dst[start] = v/10 + '0'
		return 2
	}

	dst[start+2] = v%10 + '0'
	dst[start+1] = (v/10)%10 + '0'
	dst[start] = v/100 + '0'
	return 3
}

// Decimal to integer.
// Returns number, characters consumed, success.
func dtoi(s string) (n int, i int, ok bool) {
	// Bigger than we need, not too big to worry about overflow
	const big = 0xFFFFFF

	n = 0
	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		n = n*10 + int(s[i]-'0')
		if n >= big {
			return big, i, false
		}
	}
	if i == 0 {
		return 0, 0, false
	}
	return n, i, true
}
