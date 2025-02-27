package dos2unix

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

type writer struct {
	io.Writer
}

func TestDOS2Unix(t *testing.T) {
	buf := make([]byte, 10)
	output := bytes.NewBuffer(make([]byte, 0, 100))
	w := &writer{output}
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
		{
			"Hello\r\r\n",
			"Hello\r\n",
		},
		{
			"Hello\r\r\r\n",
			"Hello\r\r\n",
		},
		{
			"Hello\r\r\r\r\n",
			"Hello\r\r\r\n",
		},
	} {
		for i := 1; i < 10; i++ {
			output.Reset()
			io.CopyBuffer(w, DOS2Unix(strings.NewReader(test.Input)), buf[:i])

			if !bytes.Equal(output.Bytes(), []byte(test.Output)) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}

func TestUnix2DOS(t *testing.T) {
	buf := make([]byte, 10)
	output := bytes.NewBuffer(make([]byte, 0, 100))
	w := &writer{output}

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
			io.CopyBuffer(w, Unix2DOS(strings.NewReader(test.Input)), buf[:i])

			if !bytes.Equal(output.Bytes(), []byte(test.Output)) {
				t.Errorf("test %d.%d: expected output: %q, got %q", n+1, i, test.Output, output)
			}
		}
	}
}
