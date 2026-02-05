package merkle

import (
	"crypto/sha256"
	"fmt"
)

var ZeroCache [33][32]byte

func Init() {
	for i := 1; i < 33; i++ {
		buffer := make([]byte, 64)
		copy(buffer[0:32], ZeroCache[i-1][:])
		copy(buffer[32:], ZeroCache[i-1][:])
		ZeroCache[i] = sha256.Sum256(buffer)
	}
}

func HashedLayer(leaves [][32]byte) [32]byte {
	length := len(leaves)

	if length == 0 {
		return [32]byte{}
	}

	currentLayer := leaves

	if length%2 != 0 {
		currentLayer = append(currentLayer, ZeroCache[0])
	}

	nextLayer := make([][32]byte, 0, len(currentLayer)/2)

	for i := 0; i < len(currentLayer); i += 2 {
		buffer := make([]byte, 0, 64)
		copy(buffer[0:32], currentLayer[i][:])
		copy(buffer[32:64], currentLayer[i+1][:])
		nextLayer = append(nextLayer, sha256.Sum256(buffer))
	}

	currentLayer = nextLayer

	return currentLayer[0]
}

func Pack(data []byte) [][32]byte {
	numChunks := (len(data) + 31) / 32

	if len(data) == 0 {
        return [][32]byte{{}}
    }

	packedData := make([][32]byte,  numChunks)

	for i := 0; i < numChunks; i++ {
		startIndex := i * 32
		stopIndex := startIndex + 32

		fmt.Println(startIndex, stopIndex, len(data), stopIndex > len(data))

		if stopIndex > len(data) {
			copy(packedData[i][:], data[startIndex:])
		} else {
			copy(packedData[i][:], data[startIndex:stopIndex])
		}
	}

	return packedData
}
