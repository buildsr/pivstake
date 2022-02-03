package keeper_test

import (
	"testing"

	testkeeper "github.com/buildsr/pivstake/testutil/keeper"
	"github.com/buildsr/pivstake/x/pivstake/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PivstakeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
