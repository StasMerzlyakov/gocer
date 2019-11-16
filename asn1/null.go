/**
 */
package asn1

import (
	"bytes"
	"fmt"
	"reflect"
)

// Encode function
// no identifiers
func encodeNull(bbuffer *bytes.Buffer) {
	encodeLength(bbuffer, 0)
}

// Decode functon
// (with identifier )
func decodeNull(val reflect.Value, bbuffer *bytes.Buffer) {
	// get the value that the pointer v points to.
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	var next byte
	var err error

	if next, err = bbuffer.ReadByte(); err != nil {
		panic(err)
	}
	if next != 0x00 {
		panic(DecodeError(fmt.Sprintf("Expected 0x00, found %x", next)))
	}
	val.Set(reflect.Zero(val.Type()))

}
