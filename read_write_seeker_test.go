package bytesextra_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xaionaro-go/bytesextra"
)

func TestReadWriteSeeker(t *testing.T) {
	b := []byte{1, 2, 3, 4, 5}
	w := bytesextra.NewReadWriteSeeker(b)

	n, err := w.Write([]byte{8, 9})
	assert.NoError(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, uint(2), w.CurrentPosition)

	r := make([]byte, 2)
	n, err = w.Read(r)
	assert.NoError(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, []byte{3, 4}, r)
	assert.Equal(t, uint(4), w.CurrentPosition)

	n, err = w.Write([]byte{10, 11})
	assert.Equal(t, io.ErrShortWrite, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, uint(5), w.CurrentPosition)

	n, err = w.Write([]byte{10, 11})
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 0, n)
	assert.Equal(t, uint(5), w.CurrentPosition)

	pos, err := w.Seek(4, io.SeekStart)
	assert.NoError(t, err)
	assert.Equal(t, int64(4), pos)
	assert.Equal(t, uint(4), w.CurrentPosition)

	n, err = w.Read(r)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, []byte{10}, r[:n])
	assert.Equal(t, uint(5), w.CurrentPosition)

	n, err = w.Read(r)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 0, n)
	assert.Equal(t, uint(5), w.CurrentPosition)

	pos, err = w.Seek(-1, io.SeekEnd)
	assert.NoError(t, err)
	assert.Equal(t, int64(4), pos)
	assert.Equal(t, uint(4), w.CurrentPosition)

	pos, err = w.Seek(-1, io.SeekCurrent)
	assert.NoError(t, err)
	assert.Equal(t, int64(3), pos)
	assert.Equal(t, uint(3), w.CurrentPosition)

	assert.Equal(t, []byte{8, 9, 3, 4, 10}, b)
}
