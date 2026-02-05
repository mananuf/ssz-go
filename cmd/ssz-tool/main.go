package main

import (
	"fmt"
	"github.com/mananuf/ssz-go/pkg/codec"
	"github.com/mananuf/ssz-go/pkg/merkle"
	"github.com/mananuf/ssz-go/pkg/types"
)

func main() {
	txs := types.Transaction{
		FromID: 10,
		ToIDs:  []uint16{10, 20, 30, 40},
	}

	txs2 := types.Transaction{
		FromID: 50,
		ToIDs:  []uint16{60, 70, 80, 90},
	}

	txBlock := types.TxBlock{
		Txs: []types.Transaction{txs, txs2},
	}

	fmt.Println(txBlock)

	marshalTxBatch := codec.MarshalTxBatch(txBlock)

	leaf0 := [32]byte{}
	for i := range leaf0 {
		leaf0[i] = 1
	}

	leaf1 := [32]byte{}
	for i := range leaf1 {
		leaf1[i] = 2
	}

	leaf2 := [32]byte{}
	for i := range leaf2 {
		leaf2[i] = 3
	}

	leaf3 := [32]byte{}
	for i := range leaf3 {
		leaf3[i] = 4
	}

	input := [][32]byte{leaf0, leaf1, leaf2, leaf3}

	merkleRoot := merkle.HashedLayer(input)

	fmt.Println(merkleRoot)

	packedData := merkle.Pack(marshalTxBatch)

	fmt.Println(marshalTxBatch)
	fmt.Println(packedData)

	marshalTxBatchCopy := make([]byte, 72)
	copy(marshalTxBatchCopy[0:36], marshalTxBatch)
	copy(marshalTxBatchCopy[36:72], marshalTxBatch)

	packedDataCopy := merkle.Pack(marshalTxBatchCopy)

	fmt.Println(marshalTxBatchCopy)
	fmt.Println(packedDataCopy)

	batch := types.Batch {
		Version: 10,
		Data: []uint32{10, 20, 30, 40, 50},
	}

	fmt.Println(merkle.HashTreeRootBatch(batch))
	
}
