package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/Peixer/myblockchain/x/myblockchain/types"
    "github.com/Peixer/myblockchain/x/myblockchain/keeper"
    keepertest "github.com/Peixer/myblockchain/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MyblockchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
