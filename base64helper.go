package base64helper

import (
	"encoding/base64"
	"errors"
	"unicode/utf8"
)

// ErrInvalidUTF8 is returned when decoded bytes are not valid UTF-8.
var ErrInvalidUTF8 = errors.New("decoded bytes are not valid UTF-8")

// DecodeStringToString returns the string represented by the 'enc'-encoded base64 string 's'.
// Decoded from 's' bytes must consist entirely of valid UTF-8-encoded runes.
// If 'enc' is nil, [base64.StdEncoding] is used.
func DecodeStringToString(enc *base64.Encoding, s string) (string, error) {
	if enc == nil {
		enc = base64.StdEncoding
	}
	bb, err := enc.DecodeString(s)
	if err != nil {
		return "", err
	}
	if !utf8.Valid(bb) {
		return "", ErrInvalidUTF8
	}
	return string(bb), nil
}
