package network

import (
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// ServerConfig is the configuration for the sequencer server.
type ServerConfig struct {
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string `mapstructure:"GRPCPort"`
	// AdminAddress is the address of the admin to verify the admin signature
	AdminAddress string `mapstructure:"AdminAddress"`
}

// ClientConfig is the configuration for the client node.
type ClientConfig struct {
	// GrpcURLs is the list of gRPC URLs
	GrpcURLs []string `mapstructure:"GrpcURLs"`
	// Chain is the chain name of the blockchain
	Chain string `mapstructure:"Chain"`
	// OperatorAddress is the address of the operator
	OperatorAddress string `mapstructure:"OperatorAddress"`
	// EthereumURL is the endpoint of the ethereum node
	EthereumURL string `mapstructure:"EthereumURL"`
	// CommitteeSCAddress is the address of the committee smart contract
	CommitteeSCAddress string `mapstructure:"CommitteeSCAddress"`
	// BLSPrivateKeyPath is the path of the BLS keystore file
	BLSKeystorePath string `mapstructure:"BLSKeystorePath"`
	// BLSKeystorePassword is the password of the BLS keystore file
	BLSKeystorePassword string `mapstructure:"BLSKeystorePassword"`
	// BLSKeystorePasswordPath is the path of the password file of the BLS keystore file
	BLSKeystorePasswordPath string `mapstructure:"BLSKeystorePasswordPath"`
	// SignerECDSAKeystorePath is the path of the ECDSA keystore file
	SignerECDSAKeystorePath string `mapstructure:"SignerECDSAKeystorePath"`
	// SignerECDSAKeystorePassword is the password of the ECDSA keystore file
	SignerECDSAKeystorePassword string `mapstructure:"SignerECDSAKeystorePassword"`
	// SignerECDSAKeystorePasswordPath is the path of the password file of the ECDSA keystore file
	SignerECDSAKeystorePasswordPath string `mapstructure:"SignerECDSAKeystorePasswordPath"`
	// PullInterval is the interval to pull the latest proof
	PullInterval utils.TimeDuration `mapstructure:"PullInterval"`
	// BLSCurve is the curve used for BLS signature
	BLSCurve string `mapstructure:"BLSCurve"`
}
