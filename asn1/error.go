/**
 */

package asn1

import "errors"

type asn1Error struct{ error }

func LengthError(desc string) error {
	return asn1Error{errors.New(desc)}
}

func NotImplemented(desc string) error {
	return asn1Error{errors.New(desc)}
}

func NotSupported(desc string) error {
	return asn1Error{errors.New(desc)}
}

func DecodeError(desc string) error {
	return asn1Error{errors.New(desc)}
}

func EncodeError(desc string) error {
	return asn1Error{errors.New(desc)}
}
