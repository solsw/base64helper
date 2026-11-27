# base64helper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/base64helper.svg)](https://pkg.go.dev/github.com/solsw/base64helper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/base64helper)

Package `base64helper` contains helpers for Go's [encoding/base64](https://pkg.go.dev/encoding/base64) package.

## Installation

```sh
go get github.com/solsw/base64helper
```

```go
import "github.com/solsw/base64helper"
```

## API

### DecodeStringToString

```go
func DecodeStringToString(enc *base64.Encoding, s string) (string, error)
```

Decodes `s` using `enc` and returns the decoded text as a UTF-8 string. On error, it returns an empty string.

- Pass `nil` for the standard padded encoding, or `base64.RawStdEncoding` for the standard alphabet without padding.
- For malformed input, the function returns the underlying decoder's `base64.CorruptInputError`. Its value is the zero-based byte offset reported by the decoder.
- If decoding succeeds but the bytes contain invalid UTF-8, the function returns `ErrInvalidUTF8`.

### ErrInvalidUTF8

```go
var ErrInvalidUTF8 = errors.New("decoded bytes are not valid UTF-8")
```

Returned when the decoded bytes are not valid UTF-8. Test for it with [errors.Is](https://pkg.go.dev/errors#Is).

## Example

```go
package main

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/solsw/base64helper"
)

func main() {
	// Decode using the standard encoding (nil enc defaults to base64.StdEncoding).
	s, err := base64helper.DecodeStringToString(nil, "0L/RgNC40LLQtdGC")
	if err != nil {
		panic(err)
	}
	fmt.Println(s) // привет

	s, err = base64helper.DecodeStringToString(base64.RawStdEncoding, "MQ")
	if err != nil {
		panic(err)
	}
	fmt.Println(s) // 1

	_, err = base64helper.DecodeStringToString(nil, "qwerty")
	var offset base64.CorruptInputError
	if errors.As(err, &offset) {
		fmt.Println(int64(offset)) // 4
	}

	// Decoded bytes must be valid UTF-8.
	_, err = base64helper.DecodeStringToString(nil, "//79")
	fmt.Println(errors.Is(err, base64helper.ErrInvalidUTF8)) // true
}
```

## License

See [LICENSE](https://github.com/solsw/base64helper?tab=MIT-1-ov-file#MIT-1-ov-file).
