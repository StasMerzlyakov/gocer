/**

 */
package asn1

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

func TestEncode_1_Float64(t *testing.T) {

	value := math.Float64frombits(0x0000001110000000) // 273 * 2^-1046

	var bbuffer bytes.Buffer

	Encode(&bbuffer, value)

	expected := []byte{REAL_TAG, 0x05, 0x81, 0xfb, 0xea, 0x01, 0x11}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue float64
	Decode(&evalue, &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %016x, actual %016x", math.Float64bits(value), math.Float64bits(evalue))
	}

}

func TestEncode_2_Int(t *testing.T) {
	var bbuffer bytes.Buffer
	value := -32639
	Encode(&bbuffer, value)

	expected := []byte{INTEGER_TAG, 0x02, 0x80, 0x81}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue int
	Decode(&evalue, &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %d, actual %d", value, evalue)
	}
}

func TestEncode_3_Bool(t *testing.T) {
	var bbuffer bytes.Buffer
	value := true

	Encode(&bbuffer, value)
	expected := []byte{BOOLEAN_TAG, 0x01, 0xFF}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue bool
	Decode(&evalue, &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %t, actual %t", value, evalue)
	}
}

type TestStruct2 struct {
	T1, T2 TestStruct
	V      int
	B      bool
}

func TestEncode_4_Struct(t *testing.T) {
	var bbuffer bytes.Buffer

	value := TestStruct2{TestStruct{32639, 0.15625}, TestStruct{32639, 0.15625}, 0, true}

	Encode(&bbuffer, value)

	expected := []byte{0x20 + SEQUENCE_TAG, // CONSTRUCTED
		0x80, // INDEFINITE FORM
		// T1
		0x20 + SEQUENCE_TAG,           // CONSTRUCTED
		0x80,                          // INDEFINITE FORM
		INTEGER_TAG, 0x02, 0x7F, 0x7F, // INTEGER
		REAL_TAG, 0x03, 0x80, 0xFB, 0x05, // float64
		END_OF_CONTENS_TAG, 0x00,

		// T2
		0x20 + SEQUENCE_TAG,           // CONSTRUCTED
		0x80,                          // INDEFINITE FORM
		INTEGER_TAG, 0x02, 0x7F, 0x7F, // INTEGER
		REAL_TAG, 0x03, 0x80, 0xFB, 0x05, // float64
		END_OF_CONTENS_TAG, 0x00,

		// V
		INTEGER_TAG, 0x01, 0x00,

		// B
		BOOLEAN_TAG, 0x01, 0xFF,

		END_OF_CONTENS_TAG, 0x00,
	}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	// var only declared !!
	var evalue TestStruct2

	Decode(&evalue, &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %+v, actual %+v", value, evalue)
	}

}

func TestEncode_4_Slice(t *testing.T) {

	value := []TestStruct{TestStruct{}, TestStruct{1, math.Pi}}

	var bbuffer bytes.Buffer
	Encode(&bbuffer, value)

	expected := []byte{
		0x30,                    //  SEQUENCE_TAG + CONSTRUCTED
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

	// only declaration
	var evalue []TestStruct

	Decode(&evalue, &bbuffer)

	if !reflect.DeepEqual(value, evalue) {
		t.Fatalf("Expected %+v, actual %+v", value, evalue)
	}

	//  variable may be initialized
	evalue = []TestStruct{TestStruct{2, 2 * math.Pi}}
	Decode(&evalue, bytes.NewBuffer(expected))

	if !reflect.DeepEqual(value, evalue) {
		t.Fatalf("Expected %+v, actual %+v", value, evalue)
	}

}

type TestStruct3 struct {
	P, T []int
}

func TestEncode_5_Slice(t *testing.T) {

	value := TestStruct3{
		[]int{0, 1, 2},
		[]int{0, 1, 2}}

	var bbuffer bytes.Buffer
	Encode(&bbuffer, value)

	// only declaration
	var evalue TestStruct3

	Decode(&evalue, &bbuffer)

	if !reflect.DeepEqual(value, evalue) {
		t.Fatalf("Expected %+v, actual %+v", value, evalue)
	}
}

func TestEncode_1_String(t *testing.T) {

	value := "\""
	var bbuffer bytes.Buffer
	Encode(&bbuffer, value)

	expected := []byte{
		UTF8String_TAG,
		0x01,
		0x22,
	}
	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}
	var evalue string

	Decode(&evalue, &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %s, actual %s", value, evalue)
	}
}

type TestStruct4 struct {
	P, T []string
}

func TestEncode_6_Slice(t *testing.T) {

	value := TestStruct4{
		[]string{"a", "b", "1", "2"},
		[]string{"a", "b", "1", "2"}}

	var bbuffer bytes.Buffer

	Encode(&bbuffer, value)

	// only declaration
	var evalue TestStruct4

	Decode(&evalue, &bbuffer)

	if !reflect.DeepEqual(value, evalue) {
		t.Fatalf("Expected %+v, actual %+v", value, evalue)
	}
}

func TestEncode_7_Slice(t *testing.T) {

	value := TestStruct4{
		nil,
		[]string{"a", "b", "1", "2"}}

	var bbuffer bytes.Buffer
	Encode(&bbuffer, value)

	// only declaration
	var evalue TestStruct4

	Decode(&evalue, &bbuffer)

	if !reflect.DeepEqual(value, evalue) {
		t.Fatalf("Expected %+v, actual %+v", value, evalue)
	}
}

func TestEncode_Panic(t *testing.T) {

	value := TestStruct4{
		nil,
		[]string{"a", "b", "1", "2"}}

	var bbuffer bytes.Buffer
	Encode(&bbuffer, value)

	// only declaration
	var evalue int

	if err := Decode(&evalue, &bbuffer); err == nil {
		t.Fatalf("Expectid error, but error not found.")
	}
}

func TestEncode_Panic2(t *testing.T) {

	value := TestStruct4{
		nil,
		[]string{"a", "b", "1", "2"}}

	if err := Encode(nil, &value); err == nil {
		t.Fatalf("Expectid error, but error not found.")
	}
}
