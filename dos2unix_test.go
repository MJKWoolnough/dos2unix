package dos2unix

import (
	"bytes"
	"io"
	"testing"

	"github.com/MJKWoolnough/memio"
)

func TestDOS2Unix(t *testing.T) {
	buf := make([]byte, 10)
	for n, test := range []struct {
		Input, Output memio.Buffer
	}{
		{
			memio.Buffer("Hello"),
			memio.Buffer("Hello"),
		},
		{
			memio.Buffer("Hello\r\n"),
			memio.Buffer("Hello\n"),
		},
		{
			memio.Buffer("Hello\r\nWorld"),
			memio.Buffer("Hello\nWorld"),
		},
		{
			memio.Buffer("qwertyuiop\r\nasdfghjkl\r\nzxcvbnm\r\n"),
			memio.Buffer("qwertyuiop\nasdfghjkl\nzxcvbnm\n"),
		},
		{
			memio.Buffer("qwertyuiop\rasdfgkl\rzxcbnm\r"),
			memio.Buffer("qwertyuiop\rasdfgkl\rzxcbnm\r"),
		},
	} {
		for i := 1; i < 10; i++ {
			var output memio.Buffer
			input := test.Input
			io.CopyBuffer(&output, DOS2Unix(&input), buf[:i])
			if !bytes.Equal(output, test.Output) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}

func TestUnix2DOS(t *testing.T) {
	buf := make([]byte, 10)
	for n, test := range []struct {
		Input, Output memio.Buffer
	}{
		{
			memio.Buffer("Hello"),
			memio.Buffer("Hello"),
		},
		{
			memio.Buffer("Hello\n"),
			memio.Buffer("Hello\r\n"),
		},
		{
			memio.Buffer("Hello\nWorld"),
			memio.Buffer("Hello\r\nWorld"),
		},
		{
			memio.Buffer("qwertyuiop\nasdfghjkl\nzxcvbnm\n"),
			memio.Buffer("qwertyuiop\r\nasdfghjkl\r\nzxcvbnm\r\n"),
		},
	} {
		for i := 1; i < 10; i++ {
			var output memio.Buffer
			input := test.Input
			io.CopyBuffer(&output, Unix2DOS(&input), buf[:i])
			if !bytes.Equal(output, test.Output) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}
