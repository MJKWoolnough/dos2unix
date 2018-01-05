// Package dos2unix provides functions to convert between dos and unix line
// termination styles
package dos2unix

import "io"

type dos2unix struct {
	r    io.ByteReader
	b    bool
	char byte
}

func (d *dos2unix) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	var n int
	for len(b) > 0 {
		if d.b {
			b[0] = d.char
			d.b = false
			b = b[1:]
			n++
			continue
		}
		c, err := d.r.ReadByte()
		if err != nil {
			return n, err
		}
		if c == '\r' {
			d.char, err = d.r.ReadByte()
			if err != io.EOF {
				if err != nil {
					return n, err
				}
				if d.char == '\n' {
					c = '\n'
				} else {
					d.b = true
				}
			}
		}
		b[0] = c
		b = b[1:]
		n++
	}
	return n, nil
}

// DOS2Unix wraps a byte reader with a reader that replaces all instances of
// \r\n with \n
func DOS2Unix(r io.ByteReader) io.Reader {
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
func Unix2DOS(r io.ByteReader) io.Reader {
	return &unix2dos{r: r}
}
