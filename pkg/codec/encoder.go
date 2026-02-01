package codec

import (
	"github.com/mananuf/ssz-go/internal/util"
	"github.com/mananuf/ssz-go/pkg/types"
)

func MarshalBatch(b types.Batch) []byte {
	// version = uint16/8 = 4bytes
	// offset = 4 bytes
	offsetSize := 8
	totalByteSize := offsetSize + (len(b.Data) * 4)
	buffer := make([]byte, totalByteSize)

	copy(buffer[0:4], util.Uint32ToBytes(b.Version))
	copy(buffer[4:8], util.Uint32ToBytes(uint32(offsetSize)))

	for _, data := range b.Data {
		copy(buffer[offsetSize:offsetSize+4], util.Uint32ToBytes(uint32(data)))
		offsetSize += 4
	}

	return buffer
}

func MarshalTransaction(t types.Transaction) []byte {
	offset := 6
	totalByteSize := offset + (len(t.ToIDs) * 2)
	buffer := make([]byte, totalByteSize)

	copy(buffer[0:2], util.Uint16ToBytes(t.FromID))
	copy(buffer[2:6], util.Uint16ToBytes(uint16(offset)))

	for _, id := range t.ToIDs {
		copy(buffer[offset:offset+2], util.Uint16ToBytes(id))
		offset += 2
	}

	return buffer
}

func MarshalValidator(v types.Validator) []byte {
	buffer := make([]byte, 16)
	copy(buffer[0:8], util.Uint64ToBytes(v.ID))
	copy(buffer[8:16], util.Uint64ToBytes(v.EffectiveBalance))

	return buffer
}
