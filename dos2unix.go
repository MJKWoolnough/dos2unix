// Package dos2unix provides functions to convert between dos and unix line
// termination styles
package dos2unix // import "vimagination.zapto.org/dos2unix"

import (
	"io"
	"slices"
)

type byteReader struct {
	io.Reader
	buf [1]byte
}

func (b *byteReader) ReadByte() (byte, error) {
	_, err := io.ReadFull(b.Reader, b.buf[:])
	return b.buf[0], err
}

type dos2unix struct {
	r    io.Reader
	b    bool
	char [1]byte
}

func (d *dos2unix) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}

	var n int

	s := b

	if d.b {
		b[0] = d.char[0]
		s = b[1:]
		n = 1

		if d.b = len(s) == 0 && b[0] == '\r'; d.b {
			_, err := d.r.Read(d.char[:])
			if d.char[0] == '\n' {
				b[0] = '\n'
				d.b = false
			}

			return n, err
		}
	}

	m, err := d.r.Read(s)
	n += m

	lastIsCR := false

	for i := 0; i < n; i++ {
		c := b[i]
		if lastIsCR && c == '\n' {
			b = slices.Delete(b, i-1, i)
			n--
			i--
		}

		lastIsCR = c == '\r'
	}

	if lastIsCR && err == nil {
		n--
		d.char[0] = '\r'
	}

	d.b = lastIsCR

	return n, err
}

// DOS2Unix wraps a byte reader with a reader that replaces all instances of
// \r\n with \n
func DOS2Unix(r io.Reader) io.Reader {
	return &dos2unix{r: r}
}

type unix2dos struct {
	r  io.ByteReader
	lf bool
}

func (u *unix2dos) Read(b []byte) (int, error) {
	var n int
	for len(b) > 0 {
		if u.lf {
			b[0] = '\n'
			u.lf = false
			b = b[1:]
			n++
			continue
		}
		c, err := u.r.ReadByte()
		if err != nil {
			return n, err
		}
		if c == '\n' {
			u.lf = true
			c = '\r'
		}
		b[0] = c
		b = b[1:]
		n++
	}
	return n, nil
}

// Unix2DOS wraps a byte reader with a reader that replaces all instances of \n
// with \r\n
func Unix2DOS(r io.Reader) io.Reader {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = &byteReader{Reader: r}
	}
	return &unix2dos{r: br}
}
