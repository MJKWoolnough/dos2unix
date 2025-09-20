# dos2unix

[![CI](https://github.com/MJKWoolnough/dos2unix/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/dos2unix/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/dos2unix.svg)](https://pkg.go.dev/vimagination.zapto.org/dos2unix)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/dos2unix)](https://goreportcard.com/report/vimagination.zapto.org/dos2unix)

--
    import "vimagination.zapto.org/dos2unix"

Package dos2unix provides functions to convert between dos and unix line termination styles.

## Highlights

 - DOS2Unix & Unix2DOS functions which wrap `io.Reader`s to automatically convert line terminations.
 - DOS2UnixWriter & Unix2DOSWriter functions which wrap `io.Writer`s to automatically convert line terminations. NB: The `Flush()` method must be called once writing is finished.

## Usage

```go
package main

import (
	"fmt"
	"io"
	"strings"

	"vimagination.zapto.org/dos2unix"
)

func main() {
	input := strings.NewReader("hello\r\nworld\r\n")
	du := dos2unix.DOS2Unix(input)
	out, _ := io.ReadAll(du)

	fmt.Printf("%q\n", out)
	// Output: "hello\nworld\n"

	input = strings.NewReader("hello\nworld\n")
	ud := dos2unix.Unix2DOS(input)
	out, _ := io.ReadAll(ud)

	fmt.Printf("%q\n", out)
	// Output: "hello\r\nworld\r\n"
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/dos2unix
