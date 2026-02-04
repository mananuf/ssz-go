package main

import (
	"fmt"

	"github.com/mananuf/ssz-go/pkg/codec"
	"github.com/mananuf/ssz-go/pkg/types"
)

func main() {
	txs := types.Transaction {
		FromID: 10,
		ToIDs: []uint16{10, 20, 30, 40},
	}

	txs2 := types.Transaction {
		FromID: 50,
		ToIDs: []uint16{60, 70, 80, 90},
	}

	txBlock := types.TxBlock {
		Txs: []types.Transaction{txs, txs2},
	}

	fmt.Println(txBlock)

	codec.MarshalTxBatch(txBlock)
}