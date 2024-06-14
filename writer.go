package dos2unix

import "io"

// WriteFlusher combines the io.Writer interface with a buffer Flush method.
type WriteFlusher interface {
	io.Writer
	Flush() error
}

var cr = []byte{'\r'}

type dos2unixWriter struct {
	io.Writer
	lastCR bool
}

func (d *dos2unixWriter) Write(p []byte) (int, error) {
	var n int

	if d.lastCR && len(p) > 0 {
		if p[0] != '\n' {
			_, err := d.Writer.Write(cr)
			if err != nil {
				return 0, err
			}
		}

		d.lastCR = false
	}

	for pos := 0; pos < len(p); pos++ {
		if d.lastCR {
			if p[pos] == '\n' {
				if pos > 1 { // more than just \r\n
					m, err := d.Writer.Write(p[:pos-1])
					n += m

					if err != nil {
						return n, err
					}
				}

				n++
				p = p[pos:]
				pos = 0
			}
		}

		d.lastCR = p[pos] == '\r'
	}
	if len(p) > 0 {
		if d.lastCR {
			p = p[:len(p)-1]
			n++
		}

		if len(p) > 0 {
			m, err := d.Writer.Write(p)
			n += m

			if err != nil {
				return n, err
			}
		}
	}

	return n, nil
}

func (d *dos2unixWriter) Flush() error {
	if !d.lastCR {
		return nil
	}

	d.lastCR = false
	_, err := d.Writer.Write(cr)

	return err
}

// DOS2UnixWriter wraps a writer to convert \r\n into \n. It is advisable to
// call the Flush method upon completion as a final \r may be buffered.
func DOS2UnixWriter(w io.Writer) WriteFlusher {
	return &dos2unixWriter{
		Writer: w,
	}
}

type unix2dosWriter struct {
	io.Writer
}

func (u unix2dosWriter) Write(p []byte) (int, error) {
	var n int

	for pos := 0; pos < len(p); pos++ {
		if p[pos] == '\n' {
			if pos > 0 {
				m, err := u.Writer.Write(p[:pos])
				n += m

				if err != nil {
					return n, err
				}

				p = p[pos:]
				pos = 0
			}

			if _, err := u.Writer.Write(cr); err != nil {
				return n, err
			}
		}
	}

	if len(p) > 0 {
		m, err := u.Writer.Write(p)
		n += m

		if err != nil {
			return n, err
		}
	}

	return n, nil
}

// Unix2DOSWriter wraps a io.Writer to convert \n into \r\n.
func Unix2DOSWriter(w io.Writer) io.Writer {
	return unix2dosWriter{
		Writer: w,
	}
}
