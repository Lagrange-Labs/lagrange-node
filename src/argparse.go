package main

import (
	"flag"
)

type opts struct {
	port int
	nick string
	room string
	peerAddr string
	stakingEndpoint string
	stakingWS string
	attestEndpoint string
	keystore string
	address string
}

func getOpts() *opts {
	// Parse Port
	portPtr := flag.Int("port",8081,"Server listening port")
	// Parse Nickname
	nickPtr := flag.String("nick","","Nickname - CLI flag, blank by default, consider addresses or protocol TLDs later.")
	// Parse Room
	roomPtr := flag.String("room","rinkeby","Room / Network")
	// Parse Remote Peer
	peerAddrPtr := flag.String("peerAddr","","Remote Peer Address")
	// Parse ETH (Staking) URL
	stakingEndpointPtr := flag.String("stakingEndpoint","https://34.229.73.193:8545","Staking Endpoint URL:Port")
	// Parse ETH (Staking) WS
	stakingWSPtr := flag.String("stakingWS","wss://mainnet.infura.io/ws/v3/f873861ee0954155b3a560eba6151d96","Staking Listening Endpoint wss://URL/ws")
	// Parse ETH (Attestation) URL
	attestEndpointPtr := flag.String("attestEndpoint","https://eth-mainnet.gateway.pokt.network/v1/5f3453978e354ab992c4da79","Attestation Endpoint URL:Port")
	// Parse Keystore Path
	keystorePtr := flag.String("keystore","","/path/to/keystore")
	// Parse Address
	addressPtr := flag.String("address","","Staker Address")

	flag.Parse()
	
	res := opts{
		*portPtr,
		*nickPtr,
		*roomPtr,
		*peerAddrPtr,
		*stakingEndpointPtr,
		*stakingWSPtr,
		*attestEndpointPtr,
		*keystorePtr,
		*addressPtr}
	return &res
}
