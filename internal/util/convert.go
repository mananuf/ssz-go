package util

import (
	"encoding/binary"
)

func Uint16ToBytes(n uint16) []byte {
	twoBytesBuffer := make([]byte, 2)
	binary.LittleEndian.PutUint16(twoBytesBuffer, n)

	return twoBytesBuffer
}

func Uint32ToBytes(n uint32) []byte {
	fourBytesBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(fourBytesBuffer, n)

	return fourBytesBuffer
}

func Uint64ToBytes(n uint64) []byte {
	eightBytesBuffer := make([]byte, 8)
	binary.LittleEndian.PutUint64(eightBytesBuffer, n)

	return eightBytesBuffer
}
