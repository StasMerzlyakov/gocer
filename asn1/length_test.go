package asn1

import (
	"bytes"
	"testing"
)

//	Short form test
func TestShortFormLengthEncoding(t *testing.T) {
	var bbuffer bytes.Buffer

	testValue := 123
	encodeLength(&bbuffer, testValue)

	expected := []byte{0x7b}

	res := bytes.Compare(expected, bbuffer.Bytes())

	if res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	value := decodeLength(&bbuffer)

	if testValue != value {
		t.Fatalf("Expected %d, actual %d", testValue, value)
	}

}

//	Two bytes length
func TestLongFormLengthEncoding(t *testing.T) {
	var bbuffer bytes.Buffer

	testValue := 3341
	expected := []byte{0x82, 0x0d, 0x0d}
	encodeLength(&bbuffer, testValue)

	res := bytes.Compare(expected, bbuffer.Bytes())
	if res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}
	value := decodeLength(&bbuffer)

	if testValue != value {
		t.Fatalf("Expected %d, actual %d", testValue, value)
	}
}

//	Three bytes length
func TestLongFormLengthEncoding_2(t *testing.T) {
	var bbuffer bytes.Buffer

	testValue := 0x10000
	expected := []byte{0x83, 0x01, 0x00, 0x00}
	encodeLength(&bbuffer, testValue)

	res := bytes.Compare(expected, bbuffer.Bytes())
	if res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}
	value := decodeLength(&bbuffer)
	if testValue != value {
		t.Fatalf("Expected %d, actual %d", testValue, value)
	}
}

//  Indefinite form
func TestInfinitFormEncoding_1(t *testing.T) {
	var bbuffer bytes.Buffer
	testValue := -1
	encodeLength(&bbuffer, testValue)
	expected := []byte{0x80}
	res := bytes.Compare(expected, bbuffer.Bytes())
	if res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	value := decodeLength(&bbuffer)
	if testValue != value {
		t.Fatalf("Expected %d, actual %d", testValue, value)
	}
}
