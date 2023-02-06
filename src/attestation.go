package main

import (
	"time"
	context "context"
	json "encoding/json"
	host "github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"strconv"
	ethClient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Format of gossiped state root message
type StateRootMessage struct {
	StateRoot string
	Timestamp string
	BlockNumber string
	ShardedEdDSASignatureTuple string
	ECDSASignatureTuple string
	EthereumPublicKey string
}

func GenerateStateRootString(eth *ethClient.Client, block *types.Block) string {
	//5. ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the State root, Timestamp, Block Number and Sharded EdDSA Signature Tuple
	stateRootSeparator := GetSeparator()
	
	blockRoot := block.Root().String()
	blockTime := strconv.FormatUint(block.Time(),10)
	blockNumber := block.Number().String()
	chain,err := eth.ChainID(context.Background())
	if err != nil { panic(err) }
	chainID := chain.String()		
	salt := GenSalt32()
	
	stateRootStr := blockRoot + stateRootSeparator + blockTime + stateRootSeparator + blockNumber + stateRootSeparator + chainID + stateRootSeparator + salt
	LogMessage("State Root String: " + stateRootStr,LOG_INFO)

	return stateRootStr
}

/*
For gossiping of state roots:

1. State Root
2. Timestamp
3. Block Number
4. Sharded EdDSA Signature Tuple (TBD exact parameters)
5. ECDSA Signature Tuple (Parameters V,R,S): This signature should be done on a hash of the State root, Timestamp, Block Number and Sharded EdDSA Signature Tuple
6. Ethereum Public Key
*/
func ListenForBlocks(ethClients []*ethClient.Client, node host.Host, topic *pubsub.Topic, ps *pubsub.PubSub, nick string, subscription *pubsub.Subscription) {
	// Separator for gossip messaging
	stateRootSeparator := GetSeparator()
	// Pull ethClient from list of clients
	eth,ethClients := ethClientsShift(ethClients,true)
	// Track failures of cycled RPC endpoints in order to panic if all fail.
	clientFailures := 0
	for {
		block, err := eth.BlockByNumber(context.Background(),nil)
		
		// If RPC request fails, use the next one in the list until there are none left.  Then panic.
		if(err != nil) {
			clientFailures++
			if(len(ethClients) > 1 && clientFailures < len(ethClients)) {
				eth,ethClients = ethClientsShift(ethClients,true)
			} else {
				panic(err)
			}
		} else {
			clientFailures = 0
		}
		
		// concatenate relevant fields
		stateRootStr := GenerateStateRootString(eth, block)
		
		//ShardedEdDSASignatureTuple - TBD
		shardedSignatureTuple := ""
		
		stateRootStrWithShardedSignatureTuple := stateRootStr + stateRootSeparator + shardedSignatureTuple
		
		// generate hash from concatenated fields
		stateHash := KeccakHash(stateRootStrWithShardedSignatureTuple)
		
		// sign resultant hash
		creds := GetCredentials()
		privateKey := creds.privateKeyECDSA
		signature, err := crypto.Sign([]byte(stateHash), privateKey)
		if err != nil { panic(err) }
		ecdsaSignatureHex := hexutil.Encode(signature)

		//timestamp
		timestamp := time.Now().UTC().Unix()
		
		//public key
		publicKeyECDSA := creds.publicKeyECDSA
		
		stateRootMessage := StateRootMessage {
			StateRoot: stateRootStr,
			Timestamp: strconv.FormatInt(timestamp,10),
			BlockNumber: block.Number().String(),
			ShardedEdDSASignatureTuple: shardedSignatureTuple,
			ECDSASignatureTuple: ecdsaSignatureHex,
			EthereumPublicKey: hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))}
		
		json,err := json.Marshal(stateRootMessage)
		if err != nil { panic(err) }
		bytes := []byte(json)
		msg := string(bytes)
		
		WriteMessages(node,topic,creds.address.Hex(),msg,"StateRootMessage")
		
		time.Sleep(1 * time.Second)
	}
}
