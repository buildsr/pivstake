package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/buildsr/pivstake/testutil/keeper"
	"github.com/buildsr/pivstake/x/pivstake/keeper"
	"github.com/buildsr/pivstake/x/pivstake/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PivstakeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
