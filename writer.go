package bytesextra

// Writer is a wrapper around `[]byte` to convert it into `io.Writer`
//
// `Write` will write into the `[]byte`.
//
// DEPRECATED: use ReadWriteSeeker instead
type Writer = ReadWriteSeeker

// NewWriter returns a new instance of `Writer`.
//
// Writer is a wrapper around `[]byte` (`storage`) to convert it into `io.Writer`.
//
// `Write` will write into the `storage`.
//
// DEPRECATED: use NewReadWriteSeeker instead
func NewWriter(storage []byte) *Writer {
	return NewReadWriteSeeker(storage)
}
