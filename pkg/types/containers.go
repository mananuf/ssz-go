package types

type Batch struct {
    Version uint32
    Data    []uint32
}

type Transaction struct {
    FromID uint16
    ToIDs  []uint16
}

type TxBlock struct {
    Txs []Transaction
}

type Validator struct {
    ID             uint64
    EffectiveBalance uint64
}