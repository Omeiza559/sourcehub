package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/sourcenetwork/sourcehub/testutil/keeper"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AcpKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
