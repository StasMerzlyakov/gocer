package asn1

/**
 */

import (
	"bytes"
	"fmt"
	"reflect"
)

const (
	MAX_STRING_LENGTH = 1000
)

// Encode function (value encoded without identifier byte)
func encodeStringPrimitive(bbuffer *bytes.Buffer, value string) {

	// Length
	length := len(value)
	encodeLength(bbuffer, length)
	if length > 0 {
		if length > MAX_STRING_LENGTH {
			panic(EncodeError("String length too long. Use constructed encoding"))
		}
		bbuffer.Write([]byte(value))
	}
}

// Encode function (value encoded without identifier byte)
func encodeStringConstructed(bbuffer *bytes.Buffer, value string) {

	// Length
	length := len(value)

	if length <= MAX_STRING_LENGTH {
		panic(EncodeError("String length too short. Use primitive encoding"))
	}

	encodeLength(bbuffer, -1)

	ba := []byte(value)
	for i := 0; length > 0; length, i = length-MAX_STRING_LENGTH, i+MAX_STRING_LENGTH {
		tail := MAX_STRING_LENGTH
		if length < tail {
			tail = length
		}

		//  The string fragments contained in the constructed encoding
		// shall be encoded with a primitive encoding
		encodePrimitiveIdentifier(bbuffer, OCTET_STRING_TAG)
		encodeLength(bbuffer, tail)

		if n, err := bbuffer.Write(ba[i : i+tail]); err != nil {
			panic(err)
		} else {
			if n != tail {
				panic(fmt.Sprintf("actual write %d, expected %d\n", n, tail))
			}
		}
	}

	encodeEndOfContent(bbuffer)
}

// Decode functon
// (no identifier byte expected)
func decodeStringPrimitive(val reflect.Value, bbuffer *bytes.Buffer) {

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.String {
		panic(DecodeError("Unexpected value. Expected string"))
	}
	lenght := decodeLength(bbuffer)

	bbytes := make([]byte, lenght, lenght)
	if lenght > 0 {
		n, err := bbuffer.Read(bbytes)
		if err != nil {
			panic(err)
		}
		if n != lenght {
			panic(DecodeError(fmt.Sprintf("Expected %d byte to read, actual: %d", lenght, n)))
		}
	}

	val.SetString(string(bbytes))
}

// Decode function (value encoded without identifier byte)
func decodeStringConstructed(val reflect.Value, bbuffer *bytes.Buffer) {

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.String {
		panic(DecodeError("Unexpected value. Expected string"))
	}
	var err error
	length := decodeLength(bbuffer)

	if length != -1 {
		panic(DecodeError("Expected length: -1 (IndefiniteForm)"))
	}

	var stringBuf bytes.Buffer
L:
	for {
		var asn1Type ASN1Type
		asn1Type, _ = decodeIdentifier(bbuffer)

		switch asn1Type {
		case OCTET_STRING_TAG:
			lenght := decodeLength(bbuffer)
			bbytes := make([]byte, lenght, lenght)
			if lenght > 0 {
				n, err := bbuffer.Read(bbytes)
				if err != nil {
					panic(DecodeError(err.Error()))
				}
				if n != lenght {
					panic(DecodeError(fmt.Sprintf("Expected %d byte to read, actual: %d", lenght, n)))
				}
			}
			stringBuf.Write(bbytes)
		case END_OF_CONTENS_TAG:
			var next byte
			if next, err = bbuffer.ReadByte(); err != nil {
				panic(err)
			}
			if next != 0x00 {
				panic(DecodeError(fmt.Sprintf("Expected 0x00, found %02x", next)))
			}
			break L
		default:
			panic(DecodeError(fmt.Sprintf("Unexpected tag %02x. Expected %02x or %02x", asn1Type, UTF8String_TAG, END_OF_CONTENS_TAG)))
		}
	}

	val.SetString(string(stringBuf.Bytes()))
}
