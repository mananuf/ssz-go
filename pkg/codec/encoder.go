package codec

import (
	"github.com/mananuf/ssz-go/internal/util"
	"github.com/mananuf/ssz-go/pkg/types"
)

func MarshalValidator(v types.Validator) []byte {
	buffer := make([]byte, 16)
	copy(buffer[0:8], util.Uint64ToBytes(v.ID))
	copy(buffer[8:16], util.Uint64ToBytes(v.EffectiveBalance))

	return buffer
}