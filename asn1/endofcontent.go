/**
 */
package asn1

import (
	"bytes"
	"fmt"
)

// Encode function
// with identifiers
func encodeEndOfContent(bbuffer *bytes.Buffer) {
	encodePrimitiveIdentifier(bbuffer, END_OF_CONTENS_TAG)
	encodeLength(bbuffer, 0)
}

// Decode functon
// (with identifier )
func decodeEndOfContent(bbuffer *bytes.Buffer) error {
	var next byte
	var err error
	if next, err = bbuffer.ReadByte(); err != nil {
		return err
	}
	if next != END_OF_CONTENS_TAG {
		return DecodeError(fmt.Sprintf("Expected 0x00, found %x", next))
	}

	if next, err = bbuffer.ReadByte(); err != nil {
		return err
	}
	if next != 0x00 {
		return DecodeError(fmt.Sprintf("Expected 0x00, found %x", next))
	}
	return nil

}
