# Simple Go ASN1 Canonical Encoding Rules (CER) encoding
- [x] int <-> INTEGER
- [x] float64 <-> REAL
- [x] bool <-> BOOLEAN
- [x] struct <-> SEQUENCE
- [x] slice <-> SEQUENCEOF
- [x] string <-> UTF8String
- [x] nil <-> NULL  // for slice




## Usage
```
import (
	"github.com/StasMerzlyakov/gocer/asn1"
)
type TestStuct struct {
	Id int
	E  float64
}

// encode to bytes.Buffer
value := []TestStruct{TestStruct{}, TestStruct{1, math.Pi}}
var bbuffer bytes.Buffer
asn1.Encode(value, &bbuffer)

// decode from bytes.Buffer
var evalue []TestStruct
asn1.Decode(&evalue, &bbuffer)
```
## Authors
- **Stas Merzlyakov** - *Initial work* - [info](https://github.com/StasMerzlyakov)
## Links
- ASN1 specification [ASN1 Spec](https://www.itu.int/ITU-T/studygroups/com17/languages/X.690-0207.pdf)

