package base64helper

import (
	"encoding/base64"
	"errors"
	"unicode/utf8"
)

// DecodeStringToString returns the string represented by the base64 string 's'.
// Decoded from base64 's' must consist entirely of valid UTF-8-encoded runes.
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
