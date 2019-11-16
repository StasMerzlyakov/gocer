package asn1

/**
 */

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeInteger_1(t *testing.T) {
	var bbuffer bytes.Buffer
	value := 32639
	encodeInteger(&bbuffer, value)

	expected := []byte{0x02, 0x7F, 0x7F}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue int
	decodeInteger(reflect.ValueOf(&evalue), &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %d, actual %d", value, evalue)
	}

}

func TestEncodeInteger_2(t *testing.T) {
	var bbuffer bytes.Buffer
	value := -32639
	encodeInteger(&bbuffer, value)
	expected := []byte{0x02, 0x80, 0x81}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue int
	decodeInteger(reflect.ValueOf(&evalue), &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %d, actual %d", value, evalue)
	}
}

func TestEncodeInteger_3(t *testing.T) {
	var bbuffer bytes.Buffer
	value := 0

	encodeInteger(&bbuffer, value)

	expected := []byte{0x01, 0x00}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue int
	decodeInteger(reflect.ValueOf(&evalue), &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %d, actual %d", value, evalue)
	}

}
