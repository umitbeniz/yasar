package yasar_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/umitbeniz/yasar/testutil/keeper"
	"github.com/umitbeniz/yasar/testutil/nullify"
	"github.com/umitbeniz/yasar/x/yasar"
	"github.com/umitbeniz/yasar/x/yasar/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.YasarKeeper(t)
	yasar.InitGenesis(ctx, *k, genesisState)
	got := yasar.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
