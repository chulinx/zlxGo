package endecryption

import (
	"encoding/base64"
	"strings"
)

const defaultConfusion = "^Z*x5lGo"

type Base64 struct {
	encoding  base64.Encoding
	confusion string
}

func NewBase64(confusions ...string) *Base64 {
	var confusion string
	encode := base64.RawStdEncoding
	if len(confusions) == 0 {
		confusion = defaultConfusion
		return &Base64{
			encoding:  *encode,
			confusion: confusion,
		}
	}
	confusion = confusions[0]
	return &Base64{
		encoding:  *encode,
		confusion: confusion,
	}
}

// Base64Encode two decode
func (b Base64) Base64Encode(s string) string {
	stringNotNull(s)
	oneBase64 := b.encoding.EncodeToString([]byte(confusionString(s, b.confusion)))
	return b.encoding.EncodeToString([]byte(oneBase64))
}

func (b Base64) Base64Decode(string2 string) (string, error) {
	stringNotNull(string2)
	decodeString, err := b.encoding.DecodeString(string2)
	if err != nil {
		return "", err
	}
	decodeString, err = b.encoding.DecodeString(string(decodeString))
	if err != nil {
		return "", err
	}
	return clarifyString(string(decodeString), b.confusion), nil
}

func stringNotNull(s string) {
	if s == "" {
		panic("s not allow is nil")
	}
}

// clarifyString clear confusionString
// Param: s: ccwork confusion: zlx
// Example: confusionString: ===> zlccworxk
// clarifyString: zlccworxk ===> ccwork
func confusionString(s, confusion string) string {
	cl := len(confusion)
	sl := len(s)
	return confusion[0:cl-1] + s[0:sl-1] + confusion[cl-1:] + s[sl-1:]
}

func clarifyString(s, confusion string) string {
	cl := len(confusion)
	s1 := strings.TrimPrefix(s, confusion[0:cl-1])
	splitS := strings.Split(s1, "")
	splitS[len(splitS)-2] = ""
	result := strings.Join(splitS, "")
	return result
}
