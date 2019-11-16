/**

 */
package asn1

import (
	"bytes"
	"fmt"
	"math"
	"math/bits"
	"reflect"
)

const PLUS_INFINITY = 0x40
const MINUS_INFINITY = 0x41
const IEEE_754_BIAS = 0x3FF
const IEEE_754_EMAX = 0x3FF
const IEEE_754_EMIN = 1 - 0x3FF
const IEEE_754_T = 52
const IEEE_754_W = 11

func encodeMantissa(bbuffer *bytes.Buffer, mantissa uint64) error {

	flag := false
	for ; mantissa > 0; mantissa <<= 8 {
		hb := byte((mantissa & 0xff00000000000000) >> 56)
		if hb > 0 {
			flag = true
		}
		if flag {
			if err := bbuffer.WriteByte(hb); err != nil {
				return err
			}
		}
	}
	return nil
}

func decodeMantissa(bbuffer *bytes.Buffer, mlen int) (uint64, error) {
	var result uint64

	for {
		if mlen == 0 {
			break
		}
		result <<= 8
		if next, err := bbuffer.ReadByte(); err != nil {
			return 0, err
		} else {
			result += uint64(next)
			mlen -= 1
		}
	}
	return result, nil
}

// Encode function (value encoded without identifier byte)
func encodeReal(bbuffer *bytes.Buffer, value float64) {

	if math.IsNaN(value) {
		panic(NotImplemented("NaN not supported"))
	}

	if math.IsInf(value, 1) {
		encodeLength(bbuffer, 1)
		bbuffer.WriteByte(PLUS_INFINITY)
		return
	}

	if math.IsInf(value, -1) {
		encodeLength(bbuffer, 1)
		bbuffer.WriteByte(MINUS_INFINITY)
		return
	}

	if value == 0 {
		// If the real value is the value zero, there shall be no contents octets in the encoding.
		encodeLength(bbuffer, 0)
	} else {
		// IEEE 754
		vbits := math.Float64bits(value)
		t := vbits & 0x000FFFFFFFFFFFFF
		s := vbits >> (IEEE_754_T + IEEE_754_W)
		e := vbits & 0x7FF0000000000000 >> IEEE_754_T
		var valueBuf bytes.Buffer
		var exp int
		if e == 0 {
			// If E=0 and T≠0, then r is (S, emin, (0+2^(1−p)×T));
			tlen := IEEE_754_T
			for ; t%2 == 0 && t != 0; t, tlen = t>>1, tlen-1 {
				// shift t

			}
			exp = IEEE_754_EMIN - tlen
		} else {
			// If 1≤E≤2w−2, then r is (S, (E−bias), (1+2^(1−p)×T));

			t = t + 0x0010000000000000

			for ; t%2 == 0 && t != 0; t = t >> 1 {
				// shift t
			}
			exp = int(e) - int(IEEE_754_BIAS) - (bits.Len64(t) - 1)
		}
		expbuffer := encodeIntegerV(exp)

		switch len(expbuffer.Bytes()) {
		case 1:
			// if bits 2 to 1 are 00, then the second contents octet encodes the value of the exponent
			// as a two's complement binary number;
			valueBuf.WriteByte(byte(0x80 + 0x40*s))
		case 2:
			// if bits 2 to 1 are 01, then the second and third contents octets encode the value of
			// the exponent as a two's complement binary number;
			valueBuf.WriteByte(byte(0x80 + 0x40*s + 0x01))
		case 3:
			// if bits 2 to 1 are 10, then the second, third and fourth contents octets encode the
			// value of the exponent as a two's complement binary number;
			valueBuf.WriteByte(byte(0x80 + 0x40*s + 0x10))
		default:
			valueBuf.WriteByte(byte(0x80 + 0x40*s + 0x11))
			valueBuf.WriteByte(byte(uint(len(expbuffer.Bytes()))))
		}
		valueBuf.Write(expbuffer.Bytes())
		encodeMantissa(&valueBuf, t)
		vlen := valueBuf.Len()
		encodeLength(bbuffer, vlen)
		bbuffer.Write(valueBuf.Bytes())
	}
}

// Decode functon
// (no identifier byte expected)

// Decode functon
// (no identifier byte expected)
func decodeReal(val reflect.Value, bbuffer *bytes.Buffer) {
	// get the value that the pointer v points to.
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Float64 {
		panic(DecodeError("Unexpected value. Expected float64"))
	}

	f := decodeRealV(bbuffer)
	val.SetFloat(f)
}

func decodeRealV(bbuffer *bytes.Buffer) float64 {

	var err error
	var length int

	length = decodeLength(bbuffer)
	if length == 0 {
		// If the real value is the value zero, there shall be no contents octets in the encoding.
		return 0
	}

	var next byte
	if next, err = bbuffer.ReadByte(); err != nil {
		panic(err)
	}
	length--

	if next&0x80 == 0 {
		if next&0x40 == 0 {
			// if bit 8 = 0 and bit 7 = 0, then the decimal encoding
			panic(NotSupported("decimal encoding"))
		} else {
			// if bit 8 = 0 and bit 7 = 1, then a "SpecialRealValue"
			if length == 0 {
				if next == PLUS_INFINITY {
					return math.Inf(1)
				}

				if next == MINUS_INFINITY {
					return math.Inf(-1)
				}
				panic(DecodeError("Wrong special values"))
			}
			panic(DecodeError("Wrong length for special values"))
		}
	}
	//  the binary encoding

	// bit 7 - sign;
	s := uint64((next & 0x40) >> 7)

	// bits 6,5 - binary encoding type
	if next&0x30 != 0 {
		panic(NotSupported(fmt.Sprintf("Not supported binary encoding type %x", next&0x30)))
	}

	var exp int
	switch lentType := next & 0x03; lentType {
	case 0x00:
		// if bits 2 to 1 are 00, then the second contents octet encodes the value of the exponent
		// as a two's complement binary number;
		exp = decodeIntegerV(bbuffer, 1)
		length--
	case 0x01:
		// if bits 2 to 1 are 01, then the second and third contents octets encode the value of
		// the exponent as a two's complement binary number;
		exp = decodeIntegerV(bbuffer, 2)
		length -= 2
	case 0x02:
		// if bits 2 to 1 are 10, then the second, third and fourth contents octets encode the
		// value of the exponent as a two's complement binary number;
		exp = decodeIntegerV(bbuffer, 3)
		length -= 3
	default:
		if explen, err := bbuffer.ReadByte(); err != nil {
			panic(err)
		} else {
			length--
			exp = decodeIntegerV(bbuffer, int(explen))
			length -= int(explen)

		}
	}

	var e, t uint64

	if t, err = decodeMantissa(bbuffer, length); err != nil {
		panic(err)
	}

	if exp < IEEE_754_EMIN {
		// If E=0 and T≠0, then r is (S, emin, (0+2^(1−p)×T));
		// exp := EMIN  - tlen
		tlen := IEEE_754_EMIN - exp
		t <<= IEEE_754_T - tlen
		e = 0

	} else {
		// If 1≤E≤2w−2, then r is (S, (E−bias), (1+2^(1−p)×T));
		tlen := IEEE_754_T + 1 // (1+2^(1−p)×T)
		for ; t < 0x0010000000000000; t, tlen = t<<1, tlen-1 {

		}

		t -= 0x0010000000000000

		// exp = int(e) - int(BIAS) - tlen + 1
		e = uint64(exp - 1 + tlen + int(IEEE_754_BIAS))
	}

	// make IEEE 754
	vbits := s<<(IEEE_754_T+IEEE_754_W) + e<<IEEE_754_T + t
	return math.Float64frombits(vbits)

}
