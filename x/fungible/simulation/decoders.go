package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/zeta-chain/node/x/fungible/types"
)

func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key, types.KeyPrefix(types.SystemContractKey)):
			var systemContractA, systemContractB types.SystemContract
			cdc.MustUnmarshal(kvA.Value, &systemContractA)
			cdc.MustUnmarshal(kvB.Value, &systemContractB)
			return fmt.Sprintf("%v\n%v", systemContractA, systemContractB)
		case bytes.Equal(kvA.Key, types.KeyPrefix(types.ForeignCoinsKeyPrefix)):
			var foreignCoinsA, foreignCoinsB types.ForeignCoins
			cdc.MustUnmarshal(kvA.Value, &foreignCoinsA)
			cdc.MustUnmarshal(kvB.Value, &foreignCoinsB)
			return fmt.Sprintf("%v\n%v", foreignCoinsA, foreignCoinsB)
		default:
			panic(fmt.Sprintf("invalid funcgible key prefix %X", kvA.Key[:1]))
		}
	}
}
