package asn1

import (
	"bytes"
	"testing"
)

func TestEncodeIdentifier_1(t *testing.T) {
	var bbuffer bytes.Buffer
	encodeIdentifier(&bbuffer, INTEGER_TAG, true)
	expected := []byte{0x02}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	asn1Type, isConstructed := decodeIdentifier(&bbuffer)

	if isConstructed {
		t.Fatalf("isConstructed  expected %t, actual %t", false, isConstructed)
	}

	if asn1Type != INTEGER_TAG {
		t.Fatalf("Expected %d, actual %d", INTEGER_TAG, asn1Type)
	}

	bbuffer.Reset()

	encodeIdentifier(&bbuffer, INTEGER_TAG, false)

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	asn1Type, isConstructed = decodeIdentifier(&bbuffer)

	if isConstructed {
		t.Fatalf("isConstructed  expected %t, actual %t", false, isConstructed)
	}

	if asn1Type != INTEGER_TAG {
		t.Fatalf("Expected %d, actual %d", INTEGER_TAG, asn1Type)
	}

}

func TestEncodeIdentifier_2(t *testing.T) {
	var bbuffer bytes.Buffer
	encodeIdentifier(&bbuffer, SEQUENCE_TAG, true)

	expected := []byte{0x30}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	asn1Type, isConstructed := decodeIdentifier(&bbuffer)

	if !isConstructed {
		t.Fatalf("isConstructed  expected %t, actual %t", true, isConstructed)
	}

	if asn1Type != SEQUENCE_TAG {
		t.Fatalf("Expected %d, actual %d", SEQUENCE_TAG, asn1Type)
	}

	bbuffer.Reset()

	encodeIdentifier(&bbuffer, SEQUENCE_TAG, false)

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	asn1Type, isConstructed = decodeIdentifier(&bbuffer)

	if !isConstructed {
		t.Fatalf("isConstructed  expected %t, actual %t", true, isConstructed)
	}

	if asn1Type != SEQUENCE_TAG {
		t.Fatalf("Expected %d, actual %d", SEQUENCE_TAG, asn1Type)
	}

}

func TestEncodeIdentifier_3(t *testing.T) {
	var bbuffer bytes.Buffer
	encodeIdentifier(&bbuffer, BIT_STRING_TAG, false)

	expected := []byte{0x03}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Errorf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	asn1Type, isConstructed := decodeIdentifier(&bbuffer)

	if isConstructed {
		t.Errorf("isConstructed  expected %t, actual %t", false, isConstructed)
	}

	if asn1Type != BIT_STRING_TAG {
		t.Errorf("Expected %d, actual %d", BIT_STRING_TAG, asn1Type)
	}

	bbuffer.Reset()

	encodeIdentifier(&bbuffer, BIT_STRING_TAG, true)

	expected = []byte{0x23}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Errorf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	asn1Type, isConstructed = decodeIdentifier(&bbuffer)

	if !isConstructed {
		t.Errorf("isConstructed  expected %t, actual %t", true, isConstructed)
	}

	if asn1Type != BIT_STRING_TAG {
		t.Errorf("Expected %d, actual %d", BIT_STRING_TAG, asn1Type)
	}
}

func TestEncodeIdentifier_4(t *testing.T) {

	var value ASN1Type = 0x1F

	var bbuffer bytes.Buffer
	encodeIdentifier(&bbuffer, value, false)

	expected := []byte{0x1F, 0x1F}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Errorf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	asn1Type, isConstructed := decodeIdentifier(&bbuffer)

	if isConstructed {
		t.Errorf("isConstructed  expected %t, actual %t", false, isConstructed)
	}

	if asn1Type != value {
		t.Errorf("Expected %d, actual %d", value, asn1Type)
	}

	bbuffer.Reset()

	encodeIdentifier(&bbuffer, value, true)

	expected = []byte{0x3F, 0x1F}
	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Errorf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	asn1Type, isConstructed = decodeIdentifier(&bbuffer)

	if !isConstructed {
		t.Errorf("isConstructed  expected %t, actual %t", true, isConstructed)
	}

	if asn1Type != value {
		t.Errorf("Expected %d, actual %d", value, asn1Type)
	}

}
