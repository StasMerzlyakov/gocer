/**

 */
package asn1

import (
	"bytes"
	"testing"
)

func TestEncodeEndOfContent(t *testing.T) {
	var bbuffer bytes.Buffer

	encodeEndOfContent(&bbuffer)

	expected := []byte{END_OF_CONTENS_TAG, 0x00}

	if res := bytes.Compare(expected, bbuffer.Bytes()); res != 0 {
		t.Fatalf("Expected [% x], actual [% x]", expected, bbuffer.Bytes())
	}

	decodeEndOfContent(&bbuffer)
}
