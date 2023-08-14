package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/umitbeniz/yasar/testutil/keeper"
	"github.com/umitbeniz/yasar/x/yasar/keeper"
	"github.com/umitbeniz/yasar/x/yasar/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.YasarKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
