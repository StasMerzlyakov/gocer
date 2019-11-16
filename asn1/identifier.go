/**
 */
package asn1

import (
	"bytes"
	"fmt"
)

type IsConstructed bool

func encodePrimitiveIdentifier(bbuffer *bytes.Buffer, asn1Type ASN1Type) {
	encodeIdentifier(bbuffer, asn1Type, false)
}

func encodeConstructedIdentifier(bbuffer *bytes.Buffer, asn1Type ASN1Type) {
	encodeIdentifier(bbuffer, asn1Type, true)
}

func encodeIdentifier(bbuffer *bytes.Buffer, asn1Type ASN1Type, isConstructed IsConstructed) {

	// For tags with a number ranging from zero to 30 (inclusive), the identifier octets shall comprise a single octet
	// encoded as follows:
	//		a) bits 8 and 7 shall be encoded to represent the class ; HINT universal class support only
	//		b) bit 6 shall be set to zero if the encoding is primitive, and shall be set to one if the encoding is constructed.
	//		c) bits 5 to 1 shall encode the number of the tag as a binary integer with bit 5 as the most significant bit.

	var err error
	if asn1Type <= 0x1E {

		constructedFlag := byte(0x20)
		switch asn1Type {
		case BOOLEAN_TAG, INTEGER_TAG, NULL_TAG, OBJECT_IDENTIFIER_TAG, REAL_TAG:
			// primitive only
			constructedFlag = 0
		case SEQUENCE_TAG, SET_TAG:
			break
		default:
			if !isConstructed {
				constructedFlag = 0
			}
		}

		if err = bbuffer.WriteByte(constructedFlag | byte(asn1Type)); err != nil {
			panic(err)
		}

	} else {

		var tbuf bytes.Buffer

		for asn1Type > 0 {
			// reverse order
			if err = tbuf.WriteByte(byte(asn1Type & 0x7F)); err != nil {
				panic(err)
			}
			asn1Type = asn1Type >> 7
		}

		// Leading octet
		constructedFlag := byte(0)
		if isConstructed {
			constructedFlag = 0x20
		}

		if err = bbuffer.WriteByte(constructedFlag | 0x1F); err != nil {
			panic(err)
		}

		l := tbuf.Len() - 1

		for l >= 0 {
			if l > 0 {
				// bit 8 of each octet shall be set to one unless it is the last octet of the identifier octets;
				if err = bbuffer.WriteByte(byte(0x80 | tbuf.Bytes()[l])); err != nil {
					panic(err)
				}
			} else {
				// Last octet
				if err := bbuffer.WriteByte(tbuf.Bytes()[l]); err != nil {
					panic(err)
				}
			}
			l -= 1
		}

	}
}

func decodeIdentifier(bbuffer *bytes.Buffer) (ASN1Type, IsConstructed) {

	var isConstructed IsConstructed

	next, err := bbuffer.ReadByte()
	if err != nil {
		panic(err)
	}

	// bits 8 and 7
	if next&0xC0 > 0 {
		// HINT universal class support only
		panic(NotImplemented(fmt.Sprintf("universal class support only. found: %x", next)))
	}

	// bit 6
	if next&0x20 > 0 {
		isConstructed = true
	}

	var asn1Type ASN1Type

	if value := next & 0x1F; value <= 0x1E {
		asn1Type = ASN1Type(value)
	} else {
		for {
			next, err := bbuffer.ReadByte()
			if err != nil {
				panic(err)
			}
			asn1Type = asn1Type<<7 + ASN1Type(next&0x7F)
			if next&0x80 == 0 {
				break
			}
		}
	}

	return asn1Type, isConstructed
}
