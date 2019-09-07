package bytesextra_test

import (
	"io"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xaionaro-go/bytesextra"
)

func TestWriter(t *testing.T) {
	b := make([]byte, 1024)
	w := bytesextra.NewWriter(b)

	r := make([]byte, 2048)
	_, err := rand.Read(r)
	assert.NoError(t, err)

	n, err := w.Write(r)
	assert.Equal(t, 1024, n)
	assert.NoError(t, err)

	n, err = w.Write(r[1024:])
	assert.Equal(t, 0, n)
	assert.Equal(t, io.EOF, err)

	assert.Equal(t, r[:1024], b)
}
