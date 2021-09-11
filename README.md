# dos2unix
--
    import "vimagination.zapto.org/dos2unix"

Package dos2unix provides functions to convert between dos and unix line
termination styles

## Usage

#### func  DOS2Unix

```go
func DOS2Unix(r io.Reader) io.Reader
```
DOS2Unix wraps a byte reader with a reader that replaces all instances of \r\n
with \n

#### func  Unix2DOS

```go
func Unix2DOS(r io.Reader) io.Reader
```
Unix2DOS wraps a byte reader with a reader that replaces all instances of \n
with \r\n

#### func  Unix2DOSWriter

```go
func Unix2DOSWriter(w io.Writer) io.Writer
```
Unix2DOSWriter wraps a io.Writer to convert \n into \r\n

#### type WriteFlusher

```go
type WriteFlusher interface {
	io.Writer
	Flush() error
}
```

WriteFlusher combines the io.Writer interface with a buffer Flush method

#### func  DOS2UnixWriter

```go
func DOS2UnixWriter(w io.Writer) WriteFlusher
```
DOS2UnixWriter wraps a writer to convert \r\n into \n. It is advisable to call
the Flush method upon completion as a final \r may be buffered
