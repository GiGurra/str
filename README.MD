# str

`str` is a Go library providing simple and reusable functions for working with strings.

## Installation

To install the package, use `go get`:

```sh
go get github.com/GiGurra/str
```

## Usage

### SplitIntoLines

The `SplitIntoLines` function splits a string into a slice of lines.

```go
package main

import (
	"fmt"
	"github.com/GiGurra/str"
)

func main() {
	text := "a\nb\nc"
	lines := str.SplitIntoLines(text)
	fmt.Println(lines) // Output: [a b c]
}
```

### SplitBufIntoLines

The `SplitBufIntoLines` function splits a `bytes.Buffer` into a slice of lines.

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/GiGurra/str"
)

func main() {
	buffer := bytes.NewBufferString("a\nb\nc")
	lines := str.SplitBufIntoLines(buffer)
	fmt.Println(lines) // Output: [a b c]
}
```

## Testing

To run the tests, use the `go test` command:

```sh
go test ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

*This README was AI-generated and may require further editing for accuracy and completeness.*
