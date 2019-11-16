/**

 */
package asn1

import (
	"bytes"
	"reflect"
)

func encodeV(bbuffer *bytes.Buffer, value interface{}) {
	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.Int:
		encodePrimitiveIdentifier(bbuffer, INTEGER_TAG)
		encodeInteger(bbuffer, int(val.Int()))
	case reflect.Float64:
		encodePrimitiveIdentifier(bbuffer, REAL_TAG)
		encodeReal(bbuffer, val.Float())
	case reflect.Bool:
		encodePrimitiveIdentifier(bbuffer, BOOLEAN_TAG)
		encodeBool(bbuffer, val.Bool())
	case reflect.String:
		if len(val.String()) <= 1000 {
			encodePrimitiveIdentifier(bbuffer, UTF8String_TAG)
			encodeStringPrimitive(bbuffer, val.String())
		} else {
			encodeConstructedIdentifier(bbuffer, UTF8String_TAG)
			encodeStringConstructed(bbuffer, val.String())
		}

	case reflect.Struct:
		encodeConstructedIdentifier(bbuffer, SEQUENCE_TAG)
		encodeStruct(bbuffer, value)
	case reflect.Slice:
		if !val.IsNil() {
			encodeConstructedIdentifier(bbuffer, SEQUENCE_TAG)
			encodeStruct(bbuffer, value)
		} else {
			encodePrimitiveIdentifier(bbuffer, NULL_TAG)
			encodeNull(bbuffer)
		}

	default:
		panic(NotImplemented("TODO"))
	}
}

// Encode value
func Encode(bbuffer *bytes.Buffer, value interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if je, ok := r.(asn1Error); ok {
				err = je.error
			} else {
				panic(r)
			}
		}
	}()
	encodeV(bbuffer, value)
	return nil
}

func decodeV(val reflect.Value, bbuffer *bytes.Buffer) {

	asn1Type, _ := decodeIdentifier(bbuffer)

	switch asn1Type {
	case INTEGER_TAG:
		decodeInteger(val, bbuffer)
	case REAL_TAG:
		decodeReal(val, bbuffer)
	case BOOLEAN_TAG:
		decodeBool(val, bbuffer)
	case SEQUENCE_TAG:
		decodeStruct(val, bbuffer)
	case UTF8String_TAG:
		decodeStringPrimitive(val, bbuffer)
	case UTF8String_TAG + 0x20: // constructed
		decodeStringConstructed(val, bbuffer)
	case NULL_TAG:
		decodeNull(val, bbuffer)
	default:
		panic(NotImplemented("TODO"))
	}
}

func Decode(evalue interface{}, bbuffer *bytes.Buffer) (err error) {

	defer func() {
		if r := recover(); r != nil {
			if je, ok := r.(asn1Error); ok {
				err = je.error
			} else {
				panic(r)
			}
		}
	}()

	val := reflect.ValueOf(evalue)
	if val.Kind() == reflect.Ptr {
		// TODO see "encoding/json:inderect" for walks down
		val = val.Elem()
	}

	decodeV(val, bbuffer)
	return nil
}
