package influx_repositores

import (
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
}

func (influxConnection *Connection) WriteEventLog(liquiadatedEventEntity LiquiadatedEventData) {
	p := influxdb2.NewPoint(
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
	influxConnection.writeApi.WritePoint(p)
}
