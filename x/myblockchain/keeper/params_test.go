package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/Peixer/myblockchain/testutil/keeper"
	"github.com/Peixer/myblockchain/x/myblockchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MyblockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
