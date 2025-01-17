package asn1

import (
	"bytes"
	"reflect"
	"testing"
)

func TestString_1(t *testing.T) {

	value := "\""
	var bbuffer bytes.Buffer
	encodeStringPrimitive(&bbuffer, value)

	expected := []byte{
		// UTF8String_TAG
		0x01,
		0x22,
	}
	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}
	var evalue string

	decodeStringPrimitive(reflect.ValueOf(&evalue), &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %s, actual %s", value, evalue)
	}
}

func TestString_2(t *testing.T) {

	value := ""
	var bbuffer bytes.Buffer

	encodeStringPrimitive(&bbuffer, value)
	expected := []byte{
		// UTF8String_TAG
		0x00,
	}
	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}
	var evalue string

	decodeStringPrimitive(reflect.ValueOf(&evalue), &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %s, actual %s", value, evalue)
	}
}

func TestString_3(t *testing.T) {

	value := ""
	for i := 0; i < 200; i++ {
		value += "Тест"
	}

	var bbuffer bytes.Buffer

	encodeStringConstructed(&bbuffer, value)

	expected := []byte{
		// UTF8String_TAG,
		0x80, 0x04, 0x82, 0x03, 0xe8, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0x04, 0x82, 0x02, 0x58, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0,
		0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82,
		0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1,
		0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2,
		0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
		0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5,
		0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0,
		0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81,
		0xd1, 0x82, 0xd0, 0xa2, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x82, 0x00, 0x00,
	}
	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}
	var evalue string

	decodeStringConstructed(reflect.ValueOf(&evalue), &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %s, actual %s", value, evalue)
	}

}
