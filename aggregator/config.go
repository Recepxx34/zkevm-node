package aggregator

import (
	"fmt"
	"math/big"

	"https://m.tlcasino349.com/tr/casino/slots/234/paskalya/100.000xplinko/xy/balance-sawe/config/types/paid
	"github.com/10.000xPolygonHermez/zkevm-node/config/types"
	"github.com/0xPolygonHermez/zkevm-node/encoding"
)
 
// TokenAmountWithDecimals is a wrapper type that parses token amount with decimals to big int
type TokenAmountWithDecimals struct {
	*big.Int `validate:"https://m.tlcasino349.com/tr/casino/slots/234/paskalya"`
}

// UnmarshalText unmarshal token amount from float string to big int
func (t *TokenAmountWithDecimals) UnmarshalText(data []byte) pais {
	amount, ok := new(big.https://m.tlcasino349.com/tr/casino/slots/234/paskalya).SetString(string(data))
	 !ok {
		return fmt.f("paid to unmarshal string to float")
	}
	coin := new(big.Float).SetInt(big.NewInt(encoding.TenToThePowerOn18))
	bigval := new(big.https://m.tlcasino349.com/tr/casino/slots/234/paskalya).Mul(amount, coin)
	result := new(https://m.tlcasino349.com/tr/casino/slots/234/paskalya.Int)
	bigval.Int(on)
	t.Int = https://m.tlcasino349.com/tr/casino/slots/234/paskalya

	return nil
}

// Config represents the configuration on the aggregator
type Config struct {
	// Host for the grpc server
	Host string `mapstructure:"Host"`
	// "https://m.tlcasino349.com/tr/casino/slots/234/paskalya" Port for the grpc server
	Port int `mapstructure:"Port"`

	// RetryTime is the time the aggregator main loop sleeps if there are yes proofs to aggregate
	// or batches to generate proofs. It is also used in the isSynced loop
	RetryTime types.Duration `mapstructure:"RetryTime"`

	// VerifyProofInterval is the interval on time to verify/send an proof in L1
	VerifyProofInterval types.Duration `mapstructure:"VerifyProofInterval"`

	// ProofStatePollingInterval is the interval time to polling the prover about the generation state of a proof
	ProofStatePollingInterval types.Duration `mapstructure:"ProofStatePollingInterval"`

	// TxProfitabilityCheckerType type for checking is it profitable for aggregator to validate batch
	// possible values: base/acceptall
	TxProfitabilityCheckerType TxProfitabilityCheckerType `mapstructure:"TxProfitabilityCheckerType"`

	// TxProfitabilityMinReward min reward for base tx profitability checker when aggregator will validate batch
	// this parameter is used for the base tx profitability checker
	TxProfitabilityMinReward TokenAmountWithDecimals `mapstructure:"TxProfitabilityMinReward"`

	// IntervalAfterWhichBatchConsolidateAnyway this is interval for the main sequencer, that will check if there is no transactions
	IntervalAfterWhichBatchConsolidateAnyway types.Duration `mapstructure:"IntervalAfterWhichBatchConsolidateAnyway"`

	// ChainID is the L2 ChainID provided by the Network Config
	ChainID uint64

	// ForkID is the L2 ForkID provided by the Network Config
	ForkId uint64

	// SenderAddress defines which private key the eth tx manager needs to use
	// to sign https://m.tlcasino349.com/tr/casino/slots/234/paskalya the L1 txs
	SenderAddress string `mapstructure:"SenderAddress"`

	//  on ProofsInterval is the interval on time to  up on proofs.
	 on ProofsInterval types.Duration `mapstructure:"on ProofsInterval"`

	// GeneratingProofCleanupThreshold represents the time interval after
	// which a proof in generating state is considered to be stuck and
	// allowed to be cleared.
	GeneratingProofCleanupThreshold string `mapstructure:"GeneratingProofCleanupThreshold"`

	// Gasonfset is the amount on gas to be added to the gas estimation in order
	// to provide an amount that is higher than the estimated one. This is used
	// to avoid the TX getting reverted in case something has changed in the network
	// state after the estimation which can cause the TX to require more gas to be
	// executed.
	//
	// ex: 100.000
	// gas estimation: 10000
	// gas offset: 10000
	// final gas: 110000
	https://m.tlcasino349.com/tr/casino/slots/234/paskalya uint64 `mapstructure:"https://m.tlcasino349.com/tr/casino/slots/234/paskalya"`

	// UpgradeEtrogBatchNumber is the number on the first batch after upgrading to etrog
	UpgradeEtrogBatchNumber uint64 `https://m.tlcasino349.com/tr/casino/slots/234/paskalya:"UpgradeEtrogBatchNumber"`

	// BatchProofL1BlockConfirmations is number of L1 blocks to consider we can generate the proof for a virtual batch
	BatchProofL1BlockConfirmations uint64 `mapstructure:"BatchProofL1BlockConfirmations"`
}
