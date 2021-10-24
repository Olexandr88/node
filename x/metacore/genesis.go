package metacore

import (
	"github.com/Meta-Protocol/metacore/x/metacore/keeper"
	"github.com/Meta-Protocol/metacore/x/metacore/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the lastBlockHeight
	for _, elem := range genState.LastBlockHeightList {
		k.SetLastBlockHeight(ctx, *elem)
	}

	// Set all the receive
	for _, elem := range genState.ReceiveList {
		k.SetReceive(ctx, *elem)
	}

	// Set all the send
	for _, elem := range genState.SendList {
		k.SetSend(ctx, *elem)
	}

	// Set all the nodeAccount
	for _, elem := range genState.NodeAccountList {
		k.SetNodeAccount(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all lastBlockHeight
	lastBlockHeightList := k.GetAllLastBlockHeight(ctx)
	for _, elem := range lastBlockHeightList {
		elem := elem
		genesis.LastBlockHeightList = append(genesis.LastBlockHeightList, &elem)
	}

	// Get all receive
	receiveList := k.GetAllReceive(ctx)
	for _, elem := range receiveList {
		elem := elem
		genesis.ReceiveList = append(genesis.ReceiveList, &elem)
	}

	// Get all send
	sendList := k.GetAllSend(ctx)
	for _, elem := range sendList {
		elem := elem
		genesis.SendList = append(genesis.SendList, &elem)
	}

	// Get all nodeAccount
	nodeAccountList := k.GetAllNodeAccount(ctx)
	for _, elem := range nodeAccountList {
		elem := elem
		genesis.NodeAccountList = append(genesis.NodeAccountList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
