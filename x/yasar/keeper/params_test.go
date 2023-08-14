package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/umitbeniz/yasar/testutil/keeper"
	"github.com/umitbeniz/yasar/x/yasar/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.YasarKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
