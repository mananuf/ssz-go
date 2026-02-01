package types

type Batch struct {
    Version uint32   // Fixed (4 bytes)
    Data    []uint32 // Variable (requires an Offset!)
}

type Validator struct {
    ID             uint64
    EffectiveBalance uint64
}