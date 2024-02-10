package influx_repositores

import (
	"fmt"
	"math/big"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type LiquiadatedEventData interface {
	GetLiquidator() string
	GetLiquidateUser() string
	GetCollaterelAssetInUsd() float64
	GetDebtAssetInUsd() float64
	GetChainName() string
	GetCaName() string
	GetDaName() string
	GetExchangeName() string
	GetTimestamp() time.Time
	GetLiquidatedCollateralAmount() *big.Float
	GetDebtToCover() *big.Float
}

func (influxConnection *Connection) WriteEventLog(liquiadatedEventEntity LiquiadatedEventData) {
	fmt.Println("<<------------------------------------------------------")
	fmt.Println("liquiadatedEventEntity.GetExchangeName(): ", liquiadatedEventEntity.GetExchangeName())
	fmt.Println("liquiadatedEventEntity.GetChainName(): ", liquiadatedEventEntity.GetChainName())
	fmt.Println("liquiadatedEventEntity.GetLiquidator(): ", liquiadatedEventEntity.GetLiquidator())
	fmt.Println("liquiadatedEventEntity.GetLiquidateUser(): ", liquiadatedEventEntity.GetLiquidateUser())
	fmt.Println("liquiadatedEventEntity.GetCaName(): ", liquiadatedEventEntity.GetCaName())
	fmt.Println("liquiadatedEventEntity.GetDaName(): ", liquiadatedEventEntity.GetDaName())
	fmt.Println("*********")
	fmt.Println("liquiadatedEventEntity.GetCollaterelAssetInUsd(): ", liquiadatedEventEntity.GetCollaterelAssetInUsd())
	fmt.Println("liquiadatedEventEntity.GetDebtAssetInUsd(): ", liquiadatedEventEntity.GetDebtAssetInUsd())
	fmt.Println("liquiadatedEventEntity.GetLiquidatedCollateralAmount(): ", liquiadatedEventEntity.GetLiquidatedCollateralAmount())
	fmt.Println("liquiadatedEventEntity.GetDebtToCover(): ", liquiadatedEventEntity.GetDebtToCover())
	fmt.Println("------------------------------------------------------>>")

	// p := influxdb2.NewPoint(
	influxdb2.NewPoint(
		"event_logs",
		map[string]string{
			"exchange":         liquiadatedEventEntity.GetExchangeName(),
			"chain_name":       liquiadatedEventEntity.GetChainName(),
			"liquiadator":      liquiadatedEventEntity.GetLiquidator(),
			"liquiadated_user": liquiadatedEventEntity.GetLiquidateUser(),
			"ca_name":          liquiadatedEventEntity.GetCaName(),
			"da_name":          liquiadatedEventEntity.GetDaName(),
		},
		map[string]interface{}{
			"collaterel_price_usd": liquiadatedEventEntity.GetCollaterelAssetInUsd(),
			"debt_price_usd":       liquiadatedEventEntity.GetDebtAssetInUsd(),
		},
		liquiadatedEventEntity.GetTimestamp(),
	)
	// influxConnection.writeApi.WritePoint(p)
}
