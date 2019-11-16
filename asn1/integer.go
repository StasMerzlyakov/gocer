/**
 */
package asn1

import (
	"bytes"
	"fmt"
	"reflect"
)

// Encode int function (int encoded to ASN1 identifier byte)
func encodeInteger(bbuffer *bytes.Buffer, value int) {
	tbuffer := encodeIntegerV(value)
	tbuflen := tbuffer.Len()
	encodeLength(bbuffer, tbuflen)
	bbuffer.Write(tbuffer.Bytes())
}

func encodeIntegerV(value int) bytes.Buffer {
	var tbuffer bytes.Buffer

	val := uint(value)
	if value < 0 {
		g := 0x80
		for ; g+value < 0; g <<= 8 {
		}
		val = uint(g+value) + uint(g)
	}

	tbuffer.WriteByte(byte(val & 0xFF))
	val >>= 8
	for ; val != 0; val >>= 8 {
		tbuffer.WriteByte(byte(val & 0xFF))
	}

	// reverse
	for i, j := 0, tbuffer.Len()-1; i < j; i, j = i+1, j-1 {
		tbuffer.Bytes()[i], tbuffer.Bytes()[j] = tbuffer.Bytes()[j], tbuffer.Bytes()[i]
	}

	return tbuffer
}

// Decode functon
// (no identifier byte expected)
func decodeInteger(val reflect.Value, bbuffer *bytes.Buffer) {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Int {
		panic(DecodeError(fmt.Sprintf("Unexpected type %v. Expected int", val.Type())))
	}

	n := decodeLength(bbuffer)
	n = decodeIntegerV(bbuffer, n)

	val.SetInt(int64(n))
}

func decodeIntegerV(bbuffer *bytes.Buffer, ilen int) int {

	if ilen < 0 {
		panic(DecodeError("error: expected lenght < 0"))
	}

	if ilen == 0 {
		return 0
	}

	var next byte
	var err error
	if next, err = bbuffer.ReadByte(); err != nil {
		panic(err)
	}

	ilen -= 1
	g := int(0x80 & next)
	value := int(0x7f & next)

	for ilen != 0 {
		if next, err = bbuffer.ReadByte(); err != nil {
			panic(err)
		} else {
			g <<= 8
			value = value<<8 + int(next&0xff)
			ilen -= 1
		}

	}
	return value - g
}
