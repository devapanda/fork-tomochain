package common

import (
	"math/big"
)

const (
	RewardMasterPercent        = 51
	RewardVoterPercent         = 19
	RewardFoundationPercent    = 30
	HexSignMethod              = "e341eaa4"
	HexSetSecret               = "34d38600"
	HexSetOpening              = "e11f5ba2"
	EpocBlockSecret            = 800
	EpocBlockOpening           = 850
	EpocBlockRandomize         = 900
	MaxMasternodes             = 1500
	LimitPenaltyEpoch          = 4
	LimitThresholdNonceInQueue = 10
	DefaultMinGasPrice         = 25000
	MergeSignRange             = 15
	RangeReturnSigner          = 150
	MinimunMinerBlockPerEpoch  = 1
)

var TIP2019Block = big.NewInt(0)
var TIPSigning = big.NewInt(0)
var TIPRandomize = big.NewInt(0)
var IsTestnet bool = false
var StoreRewardFolder string
var RollbackHash Hash
var MinGasPrice = big.NewInt(DefaultMinGasPrice)
var TRC21IssuerSMCTestNet = HexToAddress("0x7081C72c9DC44686C7B7EAB1d338EA137Fa9f0D3")
var TRC21IssuerSMC = HexToAddress("0x8c0faeb5C6bEd2129b8674F262Fd45c4e9468bee")
var TRC21GasPriceBefore = big.NewInt(0)
var TRC21GasPrice = big.NewInt(0)
var TIPTRC21Fee = big.NewInt(0)
