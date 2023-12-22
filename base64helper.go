package base64helper

import (
	"encoding/base64"
	"errors"
	"unicode/utf8"
)

// DecodeStringToString returns the string represented by the base64 string 's'.
func DecodeStringToString(enc *base64.Encoding, s string) (string, error) {
	bb, err := enc.DecodeString(s)
	if err != nil {
		return "", err
	}
	if !utf8.Valid(bb) {
		return "", errors.New("invalid UTF-8-encoded runes")
	}
	return string(bb), nil
}
