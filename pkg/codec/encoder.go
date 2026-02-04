package codec

import (
	"github.com/mananuf/ssz-go/internal/util"
	"github.com/mananuf/ssz-go/pkg/types"
)

func MarshalBatch(b types.Batch) []byte {
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
	fixedPartSize := 6 // 2 for FromID + 4 for Offset
	variablePartSize := len(t.ToIDs) * 2
	buffer := make([]byte, fixedPartSize+variablePartSize)

	copy(buffer[0:2], util.Uint16ToBytes(t.FromID))
	copy(buffer[2:6], util.Uint32ToBytes(uint32(fixedPartSize)))

	currentPos := fixedPartSize
	for _, id := range t.ToIDs {
		copy(buffer[currentPos:currentPos+2], util.Uint16ToBytes(id))
		currentPos += 2
	}

	return buffer
}

func MarshalTxBatch(tb types.TxBlock) []byte {
	numberOfTxs := len(tb.Txs)

	fixedPartSize := numberOfTxs * 4
	var variablePartSize uint64
	serializedTxs := make([][]byte, numberOfTxs)

	for i, tx := range tb.Txs {
		data := MarshalTransaction(tx)
		serializedTxs[i] = data
		variablePartSize += uint64(len(data))
	}

	buffer := make([]byte, (fixedPartSize + int(variablePartSize)))

	copy(buffer[0:4], util.Uint32ToBytes(uint32(fixedPartSize)))

	currentOffset := fixedPartSize
	dataPointer := fixedPartSize

	for i, data := range serializedTxs {
		copy(buffer[i*4:(i*4)+4], util.Uint32ToBytes(uint32(currentOffset)))
		copy(buffer[dataPointer:dataPointer+len(data)], data)

		currentOffset += len(data)

		dataPointer += len(data)
	}

	return buffer
}

func MarshalValidator(v types.Validator) []byte {
	buffer := make([]byte, 16)
	copy(buffer[0:8], util.Uint64ToBytes(v.ID))
	copy(buffer[8:16], util.Uint64ToBytes(v.EffectiveBalance))

	return buffer
}
