package influx_repositores

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type LiquiadatedEventDataWoPrice interface {
	GetLiquidator() string
	GetLiquidateUser() string
	GetChainName() string
	GetCaName() string
	GetDaName() string
	GetExchangeName() string
	GetTimestamp() time.Time
}

func (influxConnection *Connection) WriteEventLogWoPrice(liquiadatedEventEntity LiquiadatedEventDataWoPrice) {
	p := influxdb2.NewPoint(
		"event_logs_wo_price",
		map[string]string{
			"exchange":         liquiadatedEventEntity.GetExchangeName(),
			"chain_name":       liquiadatedEventEntity.GetChainName(),
			"liquiadator":      liquiadatedEventEntity.GetLiquidator(),
			"liquiadated_user": liquiadatedEventEntity.GetLiquidateUser(),
			"ca_name":          liquiadatedEventEntity.GetCaName(),
			"da_name":          liquiadatedEventEntity.GetDaName(),
		},
		map[string]interface{}{},
		liquiadatedEventEntity.GetTimestamp(),
	)
	influxConnection.writeApi.WritePoint(p)
}
