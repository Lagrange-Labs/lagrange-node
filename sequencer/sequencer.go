package sequencer

import (
	"context"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

const (
	// SyncInterval is the interval between two block syncs after fully synced.
	SyncInterval = 1 * time.Second
)

// Sequencer is the main component of the lagrange node.
// - It is responsible for fetching block headers from the blockchain.
type Sequencer struct {
	storage         storageInterface
	rpcClient       rpctypes.RpcClient
	chainID         uint32
	lastBlockNumber uint64

	ctx    context.Context
	cancel context.CancelFunc
}

// NewSequencer creates a new sequencer instance.
func NewSequencer(cfg *Config, storage storageInterface) (*Sequencer, error) {
	rpcClient, err := rpcclient.NewClient(cfg.Chain, cfg.RPCURL, cfg.EthURL, cfg.BatchStorage)
	if err != nil {
		return nil, err
	}

	chainID, err := rpcClient.GetChainID()
	if err != nil {
		return nil, err
	}

	blockNumber, err := storage.GetLastBlockNumber(context.Background(), chainID)
	if err != nil {
		return nil, err
	}

	if cfg.FromBlockNumber > blockNumber {
		blockNumber = cfg.FromBlockNumber - 1
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Sequencer{
		storage:         storage,
		rpcClient:       rpcClient,
		lastBlockNumber: blockNumber,
		chainID:         uint32(chainID),
		ctx:             ctx,
		cancel:          cancel,
	}, nil
}

// GetChainID returns the chain ID.
func (s *Sequencer) GetChainID() uint32 {
	return s.chainID
}

// Start starts the sequencer.
func (s *Sequencer) Start() error {
	finalizedBlockNumber, err := s.rpcClient.GetL2FinalizedBlockNumber()
	if err != nil {
		return err
	}

	logger.Infof("sequencer started from %d", finalizedBlockNumber)

	for {
		select {
		case <-s.ctx.Done():
			return nil
		default:
			lastBlockNumber := s.lastBlockNumber
			if s.lastBlockNumber > finalizedBlockNumber {
				time.Sleep(SyncInterval)
				finalizedBlockNumber, err = s.rpcClient.GetL2FinalizedBlockNumber()
				if err != nil {
					return err
				}
				continue
			}
			blockHash, err := s.rpcClient.GetBlockHashByNumber(lastBlockNumber)
			if err != nil {
				if err == rpctypes.ErrBlockNotFound {
					time.Sleep(SyncInterval)
					continue
				}
				return err
			}
			if err := s.storage.AddBlock(s.ctx, &types.Block{
				ChainHeader: &types.ChainHeader{
					BlockNumber: lastBlockNumber,
					BlockHash:   blockHash,
					ChainId:     s.chainID,
				},
			}); err != nil {
				return err
			}

			s.lastBlockNumber = lastBlockNumber + 1
			time.Sleep(1 * time.Millisecond)
		}
	}
}

// Stop stops the sequencer.
func (s *Sequencer) Stop() {
	s.cancel()
}
