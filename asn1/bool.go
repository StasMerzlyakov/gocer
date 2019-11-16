/**
 */
package asn1

import (
	"bytes"
	"fmt"
	"reflect"
)

// Encode function (value encoded without identifier byte)
func encodeBool(bbuffer *bytes.Buffer, value bool) {

	// Length octet
	err := bbuffer.WriteByte(0x01)
	if err != nil {
		panic(err)
	}
	// Contents octets
	if value {
		err = bbuffer.WriteByte(0xFF)
		if err != nil {
			panic(err)
		}
	} else {
		err = bbuffer.WriteByte(0x00)
		if err != nil {
			panic(err)
		}
	}

}

// Decode functon
// (no identifier byte expected)
func decodeBool(val reflect.Value, bbuffer *bytes.Buffer) {

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Bool {
		panic(DecodeError("Unexpected value. Expected bool"))
	}

	var next byte
	var err error
	if next, err = bbuffer.ReadByte(); err != nil {
		panic(err)
	}
	if next > 1 {
		panic(DecodeError(fmt.Sprintf("Wrong length for boolean %d", next)))
	}

	if next, err = bbuffer.ReadByte(); err != nil {
		panic(err)
	}

	if next > 0 {
		val.SetBool(true)
	} else {
		val.SetBool(false)
	}
}
