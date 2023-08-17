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

func (event LiquidationCallEventEntityWithMeta) getTimestamp() time.Time {
	return event.timestamp
}

func (event LiquidationCallEventEntityWithMeta) getLiquidator() string {
	return event.liquidationCallEventEntity.Liquidator.Hex()
}

func (event LiquidationCallEventEntityWithMeta) getLiquidateUser() string {
	return event.liquidationCallEventEntity.User.Hex()
}

func (event LiquidationCallEventEntityWithMeta) getCollaterelAssetInUsd() float64 {
	return event.caInUsd
}

func (event LiquidationCallEventEntityWithMeta) getDebtAssetInUsd() float64 {
	return event.daInUsd
}

func (event LiquidationCallEventEntityWithMeta) getChainName() string {
	return event.chainName
}

func (event LiquidationCallEventEntityWithMeta) getCaName() string {
	return event.caName
}

func (event LiquidationCallEventEntityWithMeta) getDaName() string {
	return event.daName
}

func (event LiquidationCallEventEntityWithMeta) getExchangeName() string {
	return "aave3"
}
