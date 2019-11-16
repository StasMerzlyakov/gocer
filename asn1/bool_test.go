/**

 */
package asn1

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeBool_1(t *testing.T) {
	var bbuffer bytes.Buffer
	value := true

	encodeBool(&bbuffer, value)

	expected := []byte{0x01, 0xFF}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	var evalue bool
	decodeBool(reflect.ValueOf(&evalue), &bbuffer)
	if value != evalue {
		t.Fatalf("Expected%t, actual %t", value, evalue)
	}

}

func TestEncodeBool_2(t *testing.T) {
	var bbuffer bytes.Buffer
	value := false
	encodeBool(&bbuffer, value)
	expected := []byte{0x01, 0x00}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}
	var evalue bool
	decodeBool(reflect.ValueOf(&evalue), &bbuffer)
	if value != evalue {
		t.Fatalf("Expected%t, actual %t", value, evalue)
	}

}
