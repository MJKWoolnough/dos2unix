package dos2unix

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

type reader struct {
	io.Reader
}

func TestDOS2UnixWriter(t *testing.T) {
	buf := make([]byte, 10)
	output := bytes.NewBuffer(make([]byte, 0, 100))
	for n, test := range []struct {
		Input, Output string
	}{
		{
			"Hello",
			"Hello",
		},
		{
			"Hello\r\n",
			"Hello\n",
		},
		{
			"Hello\r\nWorld",
			"Hello\nWorld",
		},
		{
			"qwertyuiop\r\nasdfghjkl\r\nzxcvbnm\r\n",
			"qwertyuiop\nasdfghjkl\nzxcvbnm\n",
		},
		{
			"qwertyuiop\rasdfgkl\rzxcbnm\r",
			"qwertyuiop\rasdfgkl\rzxcbnm\r",
		},
	} {
		for i := 1; i < 10; i++ {
			output.Reset()
			d := DOS2UnixWriter(output)
			_, err := io.CopyBuffer(d, reader{strings.NewReader(test.Input)}, buf[:i])
			if err != nil {
				t.Errorf("test %d.%d: unexpected write error: %s", n+1, i, err)
			} else if err := d.Flush(); err != nil {
				t.Errorf("test %d.%d: unexpected flush error: %s", n+1, i, err)
			} else if !bytes.Equal(output.Bytes(), []byte(test.Output)) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}

func TestUnix2DOSWriter(t *testing.T) {
	buf := make([]byte, 10)
	output := bytes.NewBuffer(make([]byte, 0, 100))
	for n, test := range []struct {
		Input, Output string
	}{
		{
			"Hello",
			"Hello",
		},
		{
			"Hello\n",
			"Hello\r\n",
		},
		{
			"Hello\nWorld",
			"Hello\r\nWorld",
		},
		{
			"qwertyuiop\nasdfghjkl\nzxcvbnm\n",
			"qwertyuiop\r\nasdfghjkl\r\nzxcvbnm\r\n",
		},
	} {
		for i := 1; i < 10; i++ {
			output.Reset()
			_, err := io.CopyBuffer(Unix2DOSWriter(output), reader{strings.NewReader(test.Input)}, buf[:i])
			if err != nil {
				t.Errorf("test %d.%d: unexpected error: %s", n+1, i, err)
			} else if !bytes.Equal(output.Bytes(), []byte(test.Output)) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}
