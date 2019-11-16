/**
 */
package asn1

import (
	"bytes"
	"fmt"
)

type LENGTH_FORM int

const (
	SHORT_FORM      = 0
	LONG_FORM       = 1
	INDEFINITE_FORM = 2
)

func encodeLength(bbuffer *bytes.Buffer, lenght int) {

	if lenght < 0 {
		// INDEFINITE FORM
		bbuffer.WriteByte(0x80)
		return
	}

	if lenght < 127 {
		// Short form. One octet. Bit 8 has value "0" and bits 7-1 give the length.
		if err := bbuffer.WriteByte(byte(lenght)); err != nil {
			panic(err)
		}
	} else {
		// Long form. Two to 127 octets. Bit 8 of first octet has value "1" and bits 7-1 give
		// the number of additional length octets. Second and following octets give the length,
		// base 256, most significant digit first.

		count := 0
		less_significant := byte(0)
		var tempOut bytes.Buffer
		tlen := lenght
		for {
			if tlen == 0 {
				break
			}
			less_significant = byte(tlen)
			count += 1
			tempOut.WriteByte(less_significant)
			tlen = tlen >> 8
		}

		// выставляем в 1 первый бит
		if count > 127 {
			panic(LengthError(fmt.Sprintf("Wrong length too long %d\n", lenght)))
		}
		tempOut.WriteByte(byte(count) | 0x80)

		// записываем в обратном порядке
		for index := tempOut.Len() - 1; index >= 0; index-- {
			//bout.Write([]byte{})
			bbuffer.WriteByte(tempOut.Bytes()[index])
		}
	}
}

func decodeLength(bin *bytes.Buffer) int {
	value := 0
	next, err := bin.ReadByte()

	if err != nil {
		panic(err)
	}

	if next&0x80 == 0 {
		//  Short form. One octet. Bit 8 has value "0" and bits 7-1 give the length.
		value = int(next)
	} else {

		// Long form or Indefinite form

		// Long form. Two to 127 octets. Bit 8 of first octet has value "1" and bits 7-1 give
		// the number of additional length octets. Second and following octets give the length,
		// base 256, most significant digit first.
		length := next & 0x7f

		if length == 0 {
			// indefinite form
			return -1
		}

		for {
			if length == 0 {
				break
			}
			length--
			value <<= 8
			next, err := bin.ReadByte()
			if err != nil {
				panic(err)
			}
			value += int(next)
		}
	}

	return value
}
