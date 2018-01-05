# dos2unix
--
    import "github.com/MJKWoolnough/dos2unix"

Package dos2unix provides functions to convert between dos and unix line
termination styles

## Usage

#### func  DOS2Unix

```go
func DOS2Unix(r io.ByteReader) io.Reader
```
DOS2Unix wraps a byte reader with a reader that replaces all instances of \r\n
with \n

#### func  Unix2DOS

```go
func Unix2DOS(r io.ByteReader) io.Reader
```
Unix2DOS wraps a byte reader with a reader that replaces all instances of \n
with \r\n
