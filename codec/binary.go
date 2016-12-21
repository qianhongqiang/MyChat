package codec

import (
	"encoding/binary"
)

func GetUint16LE(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func PustUint16LE(b []byte, v uint16)  {
	binary.LittleEndian.PutUint16(b,v)
}