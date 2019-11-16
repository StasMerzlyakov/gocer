/**

 */
package asn1

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNull(t *testing.T) {
	var bbuffer bytes.Buffer

	encodeNull(&bbuffer)

	expected := []byte{0x00}

	evalue := []byte{}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	decodeNull(reflect.ValueOf(&evalue), &bbuffer)

	if evalue != nil {
		t.Fatalf("Expected nil, actual %v", evalue)
	}
}
