package governance

import (
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// Config is the configuration for the Governance module.
type Config struct {
	// EthereumURL is the endpoint of the ethereum node.
	EthereumURL string `mapstructure:"EthereumURL"`
	// StakingSCAddress is the address of the staking smart contract.
	StakingSCAddress string `mapstructure:"StakingSCAddress"`
	// StakingCheckInterval is the interval to check the staking status.
	StakingCheckInterval utils.TimeDuration `mapstructure:"StakingCheckInterval"`
	// EvidenceUploadInterval is the interval to upload the evidence.
	EvidenceUploadInterval utils.TimeDuration `mapstructure:"EvidenceUploadInterval"`
}
