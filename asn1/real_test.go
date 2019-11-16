package asn1

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

func TestEncodeMantissa(t *testing.T) {
	value := uint64(273)
	var bbuffer bytes.Buffer
	var err error
	if err = encodeMantissa(&bbuffer, value); err != nil {
		t.Fatalf(err.Error())
	}
	expected := []byte{0x01, 0x11}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
		panic("!!!@")
	}

	evalue, _ := decodeMantissa(&bbuffer, len(expected))

	if value != evalue {
		t.Fatalf("Expected %08x, actual %08x", value, evalue)
	}
}

func TestEncodeReal_1(t *testing.T) {
	var bbuffer bytes.Buffer
	value := 0.

	encodeReal(&bbuffer, value)
	expected := []byte{0x00}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue float64
	decodeReal(reflect.ValueOf(&evalue), &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %f, actual %f", value, evalue)
	}

	bbuffer.Reset()
	encodeReal(&bbuffer, math.Inf(-1))

	expected = []byte{0x01, MINUS_INFINITY}
	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	decodeReal(reflect.ValueOf(&evalue), &bbuffer)

	if !math.IsInf(evalue, -1) {
		t.Fatalf("Expected %f, actual %f", math.Inf(-1), evalue)
	}

	bbuffer.Reset()
	encodeReal(&bbuffer, math.Inf(1))

	expected = []byte{0x01, PLUS_INFINITY}
	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	decodeReal(reflect.ValueOf(&evalue), &bbuffer)

	if !math.IsInf(evalue, 1) {
		t.Fatalf("Expected %f, actual %f", math.Inf(1), evalue)
	}

}

func TestEncodeReal_2(t *testing.T) {
	var bbuffer bytes.Buffer
	value := 0.15625
	encodeReal(&bbuffer, value)
	expected := []byte{0x03, 0x80, 0xFB, 0x05}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue float64
	decodeReal(reflect.ValueOf(&evalue), &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %f, actual %f", value, evalue)
	}

}

func TestEncodeReal_3(t *testing.T) {
	value := 40.15625

	var bbuffer bytes.Buffer
	encodeReal(&bbuffer, value)

	expected := []byte{0x04, 0x80, 0xFB, 0x05, 0x05}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue float64
	decodeReal(reflect.ValueOf(&evalue), &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %f, actual %f", value, evalue)
	}
}

func TestEncodeReal_4(t *testing.T) {
	value := math.MaxFloat64

	var bbuffer bytes.Buffer
	encodeReal(&bbuffer, value)

	expected := []byte{0x0a, 0x81, 0x03, 0xcb, 0x1f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff} // 9007199254740991 * 2^971

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue float64
	decodeReal(reflect.ValueOf(&evalue), &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %f, actual %f", value, evalue)
	}
}

func TestEncodeReal_5(t *testing.T) {
	value := math.SmallestNonzeroFloat64

	var bbuffer bytes.Buffer
	encodeReal(&bbuffer, value)

	expected := []byte{0x04, 0x81, 0xfb, 0xce, 0x01} // 1 * 2^-1074

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue float64
	decodeReal(reflect.ValueOf(&evalue), &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %016x, actual %016x", math.Float64bits(value), math.Float64bits(evalue))
	}
}

func TestEncodeReal_6(t *testing.T) {
	value := math.Float64frombits(0x0000001110000000) // 273 * 2^-1046

	var bbuffer bytes.Buffer
	encodeReal(&bbuffer, value)

	expected := []byte{0x05, 0x81, 0xfb, 0xea, 0x01, 0x11}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue float64
	decodeReal(reflect.ValueOf(&evalue), &bbuffer)

	if value != evalue {
		t.Fatalf("Expected %f, actual %f", value, evalue)
	}
}

func TestEncodeReal_7(t *testing.T) {
	value := 47.

	var bbuffer bytes.Buffer
	encodeReal(&bbuffer, value)

	expected := []byte{0x03, 0x80, 0x00, 0x2f}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue float64
	decodeReal(reflect.ValueOf(&evalue), &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %f, actual %f", value, evalue)
	}
}

func TestEncodeReal_8(t *testing.T) {
	value := .5097897194236047

	var bbuffer bytes.Buffer
	encodeReal(&bbuffer, value)

	expected := []byte{0x08, 0x80, 0xd0, 0x82, 0x81, 0x94, 0x3c, 0xc2, 0xeb}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue float64
	decodeReal(reflect.ValueOf(&evalue), &bbuffer)
	if value != evalue {
		t.Fatalf("Expected %f, actual %f", value, evalue)
	}
}
