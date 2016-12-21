package codec

import (
	"encoding/binary"
	"io"
)

var zero [binary.MaxVarintLen64]byte

type Reader struct {
	R   io.Reader
	buf [1024]byte
	err error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{R: r}
}

func (reader *Reader) Reset(r io.Reader) {
	reader.R = r
	reader.err = nil
}

func (reader *Reader) Error() error {
	return reader.err
}

func (reader *Reader) Read(b []byte) (n int, err error) {
	if reader.err == nil {
		n, err = reader.R.Read(b)
		reader.err = err
	}
	return
}

func (reader *Reader) readPacket(spliter HeadSpliter) []byte {
	if reader.err != nil {
		return nil
	}
	return spliter.Read(reader)
}

func (reader *Reader) ReadBytes(l int) (b []byte) {
	b = make([]byte, l)
	_, reader.err = io.ReadFull(reader.R, b)
	return
}

func (reader *Reader) ReadString(n int) string {
	return string(reader.ReadBytes(n))
}

func (reader *Reader) ReadUvarint() (v int64) {
	if reader.err == nil {
		v, reader.err = binary.ReadUvarint(reader)
	}
	return
}

func (reader *Reader) ReadVarint() (v int64) {
	if reader.err == nil {
		v, reader.err = binary.ReadVarint(reader)
	}
	return
}
