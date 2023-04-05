package keeper

import (
	"github.com/Peixer/myblockchain/x/myblockchain/types"
)

var _ types.QueryServer = Keeper{}
