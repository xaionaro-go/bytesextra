package bytesextra

import (
	"io"
)

// Writer is a wrapper around `[]byte` to convert it into `io.Writer`
//
// `Write` will write into the `[]byte`.
type Writer struct {
	Storage         []byte
	CurrentPosition uint
}

// NewWriter returns a new instance of `Writer`.
//
// Writer is a wrapper around `[]byte` (`storage`) to convert it into `io.Writer`.
//
// `Write` will write into the `storage`.
func NewWriter(storage []byte) *Writer {
	return &Writer{
		Storage: storage,
	}
}

// Write writes `b` into the `Storage`.
//
// If the `Storage` is full then returns io.EOF as `err`.
// On partial write returns `io.ErrShortWrite` as `err`.
func (w *Writer) Write(b []byte) (n int, err error) {
	if w.CurrentPosition >= uint(len(w.Storage)) {
		return 0, io.EOF
	}
	copy(w.Storage[w.CurrentPosition:], b)
	n = min(
		int(uint(len(w.Storage))-w.CurrentPosition),
		len(b),
	)
	if n < len(b) {
		err = io.ErrShortWrite
	}
	w.CurrentPosition += uint(n)
	return
}
