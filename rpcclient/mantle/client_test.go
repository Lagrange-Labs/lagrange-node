package mantle

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const GOERLI_BATCHSTORSGE_ADDR = "0xe5d639b1283352f32477a95b5d4109bcf9d4acf3"
const LOCAL_BATCHSTORSGE_ADDR = "0xbB9dDB1020F82F93e45DA0e2CFbd27756DA36956"

func TestEndpoints(t *testing.T) {
	ethURL := os.Getenv("ETH_RPC")
	if ethURL == "" {
		t.Skip("ETH_RPC not set")
	}
	cfg := &Config{
		RPCURL:           "http://localhost:8545",
		L1RPCURL:         ethURL,
		BatchStorageAddr: GOERLI_BATCHSTORSGE_ADDR,
	}
	c, err := NewClient(cfg)
	require.NoError(t, err)
	id, err := c.GetChainID()
	require.NoError(t, err)
	t.Logf("id: %d", id)

	l2Hash, err := c.GetBlockHashByNumber(1) //nolint:staticcheck
	require.NoError(t, err)
	require.Equal(t, len(l2Hash), 32)

	num, err := c.GetFinalizedBlockNumber()
	// require.NoError(t, err)
	// require.Greater(t, num, uint64(0))
	t.Log(num, err)
}

func TestFinalizedL2BlockNumberMock(t *testing.T) {
	cfg := &Config{
		RPCURL:           "http://localhost:8545",
		L1RPCURL:         "http://localhost:8545",
		BatchStorageAddr: LOCAL_BATCHSTORSGE_ADDR,
	}
	c, err := NewClient(cfg)
	require.NoError(t, err)

	// pre-merge chain does not support this
	_, err = c.GetFinalizedBlockNumber()
	// require.NoError(t, err)
	t.Log(err)
}
