# base64helper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/base64helper.svg)](https://pkg.go.dev/github.com/solsw/base64helper)
[![Documentation](https://img.shields.io/badge/docs-github.io-green)](https://solsw.github.io/base64helper/)

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

Returns the string represented by the `enc`-encoded base64 string `s`.

- If `enc` is `nil`, [base64.StdEncoding](https://pkg.go.dev/encoding/base64#pkg-variables) is used.
- The bytes decoded from `s` must consist entirely of valid UTF-8-encoded runes; otherwise `ErrInvalidUTF8` is returned.
- If `s` is not valid base64, the error from the underlying [Encoding.DecodeString](https://pkg.go.dev/encoding/base64#Encoding.DecodeString) is returned.

### ErrInvalidUTF8

```go
var ErrInvalidUTF8 = errors.New("decoded bytes are not valid UTF-8")
```

Returned when the decoded bytes are not valid UTF-8. Test for it with [errors.Is](https://pkg.go.dev/errors#Is).

## Example

```go
package main

import (
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

	// Decoded bytes must be valid UTF-8.
	_, err = base64helper.DecodeStringToString(nil, "//79")
	fmt.Println(errors.Is(err, base64helper.ErrInvalidUTF8)) // true
}
```

## License

See [LICENSE](https://github.com/solsw/base64helper?tab=MIT-1-ov-file#MIT-1-ov-file).
