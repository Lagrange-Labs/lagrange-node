package mantle

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	evmclient.Client

	ethClient *ethclient.Client
	rpcClient *rpc.Client
}

var (
	getL2BlockNumberABI abi.ABI
	abiInput            []byte
)

func init() {
	var err error
	getL2BlockNumberABI, err = abi.JSON(strings.NewReader(`[{"inputs":[],"name":"getL2StoredBlockNumber","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		panic(err)
	}
	abiInput, err = getL2BlockNumberABI.Pack("getL2StoredBlockNumber")
	if err != nil {
		panic(err)
	}
}

// NewClient creates a new Client instance.
func NewClient(rpcURL, l1RpcURL, newRpcURL string) (*Client, error) {
	client, err := evmclient.NewClient(rpcURL)
	if err != nil {
		return nil, err
	}

	ethClient, err := ethclient.Dial(l1RpcURL)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpc.Dial(newRpcURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:    *client,
		ethClient: ethClient,
		rpcClient: rpcClient,
	}, nil
}

// GetL2FinalizedBlockNumber returns the L2 finalized block number.
func (c *Client) GetL2FinalizedBlockNumber() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := make(map[string]interface{})
	if err := c.rpcClient.CallContext(ctx, &result, "optimism_syncStatus"); err != nil {
		return 0, err
	}

	l2FinalizedInfo, ok := result["finalized_l2"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("failed to get finalized L2 Info")
	}
	blockNumber, ok := l2FinalizedInfo["number"].(float64)
	if !ok {
		return 0, fmt.Errorf("failed to get finalized L2 block number")
	}

	return uint64(blockNumber), nil
}
