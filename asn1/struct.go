/**
 */
package asn1

import (
	"bytes"
	"fmt"
	"reflect"
)

// Encode function (value encoded without identifier byte)
func encodeStruct(bbuffer *bytes.Buffer, value interface{}) {

	val := reflect.ValueOf(value)

	encodeLength(bbuffer, -1) // INDEFINITE FORM

	switch val.Kind() {

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			iface := val.Field(i).Interface()
			encodeV(bbuffer, iface)
		}

	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			iface := val.Index(i).Interface()
			encodeV(bbuffer, iface)
		}
	default:
		panic(NotImplemented("TODO !!!"))

	}

	encodeEndOfContent(bbuffer)

}

// Decode functon
// (no identifier byte expected)
func decodeStruct(val reflect.Value, bbuffer *bytes.Buffer) {

	// reald length
	length := decodeLength(bbuffer)
	if length != -1 {
		panic(DecodeError("Expected length: -1 (IndefiniteForm)"))
	}

	if val.Kind() == reflect.Ptr {
		// TODO see "encoding/json:inderect" for walks down
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			switch field.Kind() {
			case reflect.Int, reflect.Float64, reflect.Bool:
				decodeV(field, bbuffer)
			case reflect.Struct, reflect.Slice:
				evalue := field.Addr().Interface() // pass by address !!
				rv := reflect.ValueOf(evalue)
				decodeV(rv, bbuffer)
			default:
				panic(DecodeError(fmt.Sprintf("Unexpected type: %s. Expected Int, Float64, Bool, Struct, Slice", field.Type().Name())))
			}
		}

	case reflect.Slice:
		// Init if null
		if val.IsNil() {
			val.Set(reflect.MakeSlice(val.Type(), 0, 10))
		}

		// Reset slice
		val.Set(reflect.Zero(val.Type()))

		for {
			// check next byte
			var next byte
			var err error
			if next, err = bbuffer.ReadByte(); err != nil {
				panic(err)
			}
			bbuffer.UnreadByte()
			if next == END_OF_CONTENS_TAG {
				break
			}
			// read value
			v := reflect.New(val.Type().Elem())
			evalue := v.Interface() // pass by address !!
			rv := reflect.ValueOf(evalue)
			decodeV(rv, bbuffer)
			val.Set(reflect.Append(val, v.Elem()))
		}
	default:
		panic(DecodeError(fmt.Sprintf("Unexpected type %v. Expected Slice or Struct", val.Type())))

	}

	// read end_of_contens_struct
	decodeEndOfContent(bbuffer)
}
