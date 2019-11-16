package asn1

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

type TestStruct struct {
	Id int
	E  float64
}

func TestStruct_1(t *testing.T) {

	value := TestStruct{32639, 0.15625}

	var bbuffer bytes.Buffer
	encodeStruct(&bbuffer, value)

	expected := []byte{ // SEQUENCE_TAG,
		0x80,                          // INDEFINITE FORM
		INTEGER_TAG, 0x02, 0x7F, 0x7F, // int
		REAL_TAG, 0x03, 0x80, 0xFB, 0x05, // float64
		END_OF_CONTENS_TAG, 0x00,
	}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue TestStruct

	decodeStruct(reflect.ValueOf(&evalue), &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %+v, actual %+v", value, evalue)
	}

}

func TestStruct_2(t *testing.T) {

	value := []TestStruct{TestStruct{}, TestStruct{1, math.Pi}}

	var bbuffer bytes.Buffer
	encodeStruct(&bbuffer, value)

	expected := []byte{
		// SEQUENCE_TAG
		0x80,                    // INDEFINITE FORM
		0x30,                    //  SEQUENCE_TAG + CONSTRUCTED
		0x80,                    // INDEFINITE FORM
		INTEGER_TAG, 0x01, 0x00, // INTEGER(0)
		REAL_TAG, 0x00, // REALD(0)
		END_OF_CONTENS_TAG, 0x00,
		0x30,                    //  SEQUENCE_TAG + CONSTRUCTED
		0x80,                    // INDEFINITE FORM
		INTEGER_TAG, 0x01, 0x01, // INTEGER(1)
		REAL_TAG, 0x09, 0x80, 0xd0, 0x03, 0x24, 0x3f, 0x6a, 0x88, 0x85, 0xa3, // REAL(3.141592654)
		END_OF_CONTENS_TAG, 0x00,
		END_OF_CONTENS_TAG, 0x00,
	}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue []TestStruct

	decodeStruct(reflect.ValueOf(&evalue), &bbuffer)

	if !reflect.DeepEqual(value, evalue) {
		t.Fatalf("Expected %+v, actual %+v", value, evalue)
	}
}
