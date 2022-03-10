package config

const (
	VERSION      = "0.0.1"
	PORT         = 3444
	SERVICE_NAME = "gaia-contracts-go"

	ENV = "env"

	// envs
	ENV_EMULATOR       = "emulator"
	ENV_EMULATOR_CLOUD = "emulator-cloud"
	ENV_TESTNET        = "testnet"

	// custom env vars
	IPFS_ACCESS_NODE_URL                = "IPFS_ACCESS_NODE_URL"
	FLOW_ACCESS_NODE                    = "FLOW_ACCESS_NODE"
	FLOW_SERVICE_ACCOUNT_ADDRESS        = "FLOW_SERVICE_ACCOUNT_ADDRESS"
	FLOW_SERVICE_ACCOUNT_PRIVATE_KEY    = "FLOW_SERVICE_ACCOUNT_PRIVATE_KEY"
	FUSD_CONTRACT_ADDRESS               = "FUSD_CONTRACT_ADDRESS"
	FUNGIBLE_TOKEN_CONTRACT_ADDRESS     = "FUNGIBLE_TOKEN_CONTRACT_ADDRESS"
	NON_FUNGIBLE_TOKEN_CONTRACT_ADDRESS = "NON_FUNGIBLE_TOKEN_CONTRACT_ADDRESS"

	TEST        = "test"
	WORKSTATION = "workstation"
	DEV         = "dev"

	FLOW_HASH_ALGO_NAME          = "SHA3_256"
	FLOW_SIG_ALGO_NAME           = "ECDSA_P256"
	FLOW_SERVICE_ACCOUNT_SIG_ALG = "ECDSA_P256"
	FLOW_GAS_LIMIT               = uint64(1000)
	FLOW_NFT_NAME                = "NFT"
	FLOW_COLLECTION_NAME         = "NFTCollection002"
	IPFS_BASE_URL                = "https://ipfs.io/ipfs/"
)
