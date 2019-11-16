/**

 */
package asn1

type ASN1Type uint

const (
	END_OF_CONTENS_TAG    = 0x00 //  end-of-contents
	BOOLEAN_TAG           = 0x01 // primitive
	INTEGER_TAG           = 0x02 // primitive
	BIT_STRING_TAG        = 0x03 // either primitive or constructed
	OCTET_STRING_TAG      = 0x04 // primitive
	NULL_TAG              = 0x05 // primitive
	OBJECT_IDENTIFIER_TAG = 0x06 // primitive
	REAL_TAG              = 0x09 // primitive
	UTF8String_TAG        = 0x0C // either primitive or constructed
	SEQUENCE_TAG          = 0x10 // constructed
	SET_TAG               = 0x11 // constructed
)
