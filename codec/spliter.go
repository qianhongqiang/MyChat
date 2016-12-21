package codec

import (
	"io"
)

type HeadSpliter struct {
	ReadHead  func(r *Reader) int
	WriteHead func(w *Writer, l int)
}

func (s HeadSpliter) Read(r *Reader) []byte {
	n := s.ReadHead(r)
	if r.Error() != nil {
		return nil
	}
	b := make([]byte, n)
	if _, err := io.ReadFull(r, b); err != nil {
		return nil
	}
	return b
}
