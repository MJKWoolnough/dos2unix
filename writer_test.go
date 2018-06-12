package dos2unix

import (
	"bytes"
	"io"
	"testing"

	"vimagination.zapto.org/memio"
)

type jr struct {
	io.Reader
}

func TestDOS2UnixWriter(t *testing.T) {
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
			d := DOS2UnixWriter(&output)
			_, err := io.CopyBuffer(d, jr{&input}, buf[:i])
			if err != nil {
				t.Errorf("test %d.%d: unexpected write error: %s", n+1, i, err)
			} else if err := d.Flush(); err != nil {
				t.Errorf("test %d.%d: unexpected flush error: %s", n+1, i, err)
			} else if !bytes.Equal(output, test.Output) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}

func TestUnix2DOSWriter(t *testing.T) {
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
			_, err := io.CopyBuffer(Unix2DOSWriter(&output), jr{&input}, buf[:i])
			if err != nil {
				t.Errorf("test %d.%d: unexpected error: %s", n+1, i, err)
			} else if !bytes.Equal(output, test.Output) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}
