package liquidation_event

import (
	"math"
	"math/big"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/ethereum/go-ethereum/common"
)

type LiquidationCallEventEntityWithMeta struct {
	liquidationCallEventEntity LiquidationCallEventEntity
	caInUsd                    float64
	daInUsd                    float64
	chainName                  string
	caMetadata                 lib.TokenMetadataEntity
	daMetadata                 lib.TokenMetadataEntity
	timestamp                  time.Time
}

type LiquidationCallEventEntity struct {
	CollateralAsset            common.Address
	DebtAsset                  common.Address
	User                       common.Address
	DebtToCover                *big.Int
	LiquidatedCollateralAmount *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
}

func (event LiquidationCallEventEntityWithMeta) GetTimestamp() time.Time {
	return event.timestamp
}

func (event LiquidationCallEventEntityWithMeta) GetLiquidator() string {
	return event.liquidationCallEventEntity.Liquidator.Hex()
}

func (event LiquidationCallEventEntityWithMeta) GetLiquidateUser() string {
	return event.liquidationCallEventEntity.User.Hex()
}

func (event LiquidationCallEventEntityWithMeta) GetCollaterelAssetInUsd() float64 {
	return event.caInUsd
}

func (event LiquidationCallEventEntityWithMeta) GetDebtAssetInUsd() float64 {
	return event.daInUsd
}

func (event LiquidationCallEventEntityWithMeta) GetChainName() string {
	return event.chainName
}

func (event LiquidationCallEventEntityWithMeta) GetCaName() string {
	return event.caMetadata.Name
}

func (event LiquidationCallEventEntityWithMeta) GetDaName() string {
	return event.daMetadata.Name
}

func (event LiquidationCallEventEntityWithMeta) GetDebtToCover() *big.Float {
	debtInFloat := new(big.Float).SetInt(event.liquidationCallEventEntity.DebtToCover)
	denominator := new(big.Float).SetFloat64(math.Pow(10, float64(event.daMetadata.Decimal)))

	return new(big.Float).Quo(debtInFloat, denominator)
}

func (event LiquidationCallEventEntityWithMeta) GetLiquidatedCollateralAmount() *big.Float {
	debtInFloat := new(big.Float).SetInt(event.liquidationCallEventEntity.LiquidatedCollateralAmount)
	denominator := new(big.Float).SetFloat64(math.Pow(10, float64(event.caMetadata.Decimal)))

	return new(big.Float).Quo(debtInFloat, denominator)
}

func (event LiquidationCallEventEntityWithMeta) GetExchangeName() string {
	return "aave3"
}
