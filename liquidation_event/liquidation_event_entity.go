package liquidation_event

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type LiquidationCallEventEntityWithMeta struct {
	liquidationCallEventEntity LiquidationCallEventEntity
	caInUsd                    float64
	daInUsd                    float64
	chainName                  string
	caName                     string
	daName                     string
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
	return event.caName
}

func (event LiquidationCallEventEntityWithMeta) GetDaName() string {
	return event.daName
}

func (event LiquidationCallEventEntityWithMeta) GetExchangeName() string {
	return "aave3"
}
