package dos2unix_test

import (
	"fmt"
	"io"
	"strings"

	"vimagination.zapto.org/dos2unix"
)

// ExampleDOS2Unix converts CRLF line terminations to LF.
func ExampleDOS2Unix() {
	input := strings.NewReader("hello\r\nworld\r\n")
	r := dos2unix.DOS2Unix(input)
	out, _ := io.ReadAll(r)

	fmt.Printf("%q\n", out)
	// Output: "hello\nworld\n"
}

// ExampleUnix2DOS converts LF line terminations to CRLF.
func ExampleUnix2DOS() {
	input := strings.NewReader("hello\nworld\n")
	r := dos2unix.Unix2DOS(input)
	out, _ := io.ReadAll(r)

	fmt.Printf("%q\n", out)
	// Output: "hello\r\nworld\r\n"
}
