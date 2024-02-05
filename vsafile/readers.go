package vsafile

import "io"

type readerWithIndex interface {
	io.Reader
	Index() int64
}

type countingReader struct {
	r io.Reader
	i int64
}

func newCountingReader(r io.Reader) readerWithIndex {
	return &countingReader{r: r}
}

// Read implements the io.Reader interface and increments the index
func (cr *countingReader) Read(p []byte) (n int, err error) {
	n, err = cr.r.Read(p)
	cr.i += int64(n)
	return
}

// Index returns the current index
func (cr *countingReader) Index() int64 {
	return cr.i
}
