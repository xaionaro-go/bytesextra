package bytesextra

import (
	"fmt"
	"io"
)

var _ io.ReadWriteSeeker = (*ReadWriteSeeker)(nil)

// ReadWriteSeeker is a wrapper for a []byte slice which implements io.ReadWrite.Seeker.
type ReadWriteSeeker struct {
	Storage         []byte
	CurrentPosition uint
}

// NewReadWriteSeeker wraps `storage` with an implementation of io.ReadWriteSeeker.
func NewReadWriteSeeker(storage []byte) *ReadWriteSeeker {
	return &ReadWriteSeeker{
		Storage:         storage,
		CurrentPosition: 0,
	}
}

// Write implements io.Writer
//
// Write writes `b` into the `Storage`.
//
// If the `Storage` is full then returns io.EOF as `err`.
// On partial write returns `io.ErrShortWrite` as `err`.
func (w *ReadWriteSeeker) Write(b []byte) (n int, err error) {
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

// Read implements io.Reader.
//
// Read reads `Storage` into `b`.
//
// If nothing else left to read then io.EOF is returned
// as `err`.
func (w *ReadWriteSeeker) Read(b []byte) (n int, err error) {
	if w.CurrentPosition >= uint(len(w.Storage)) {
		return 0, io.EOF
	}
	copy(b, w.Storage[w.CurrentPosition:])
	n = min(
		int(uint(len(w.Storage))-w.CurrentPosition),
		len(b),
	)
	w.CurrentPosition += uint(n)
	return
}

// Seek implements io.Seeker.
func (w *ReadWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	var newPos int64
	switch whence {
	case io.SeekStart:
		newPos = offset
	case io.SeekCurrent:
		newPos = int64(w.CurrentPosition) + offset
	case io.SeekEnd:
		newPos = int64(len(w.Storage)) + offset
	}

	if newPos < 0 {
		return int64(w.CurrentPosition), fmt.Errorf("requested position is negative: %d < 0", newPos)
	}
	if newPos > int64(len(w.Storage)) {
		return int64(w.CurrentPosition), fmt.Errorf("requested position is outside of the buffer: %d > %d", newPos, len(w.Storage))
	}

	w.CurrentPosition = uint(newPos)
	return int64(w.CurrentPosition), nil
}
