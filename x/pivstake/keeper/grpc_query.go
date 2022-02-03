package keeper

import (
	"github.com/buildsr/pivstake/x/pivstake/types"
)

var _ types.QueryServer = Keeper{}
