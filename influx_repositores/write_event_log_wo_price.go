package influx_repositores

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type LiquiadatedEventDataWoPrice interface {
	getLiquidator() string
	getLiquidateUser() string
	getChainName() string
	getCaName() string
	getDaName() string
	getExchangeName() string
	getTimestamp() time.Time
}

func WriteEventLogWoPrice(liquiadatedEventData LiquiadatedEventDataWoPrice) {
	// todo: maybe should open once and pass it
	client := client()
	defer client.Close()

	writeAPI := writeAPI(client)
	defer writeAPI.Flush()

	errorsCh := writeAPI.Errors()
	go logErrors(errorsCh)

	p := influxdb2.NewPoint(
		"event_logs_wo_price",
		map[string]string{
			"exchange":         liquiadatedEventData.getExchangeName(),
			"chain_name":       liquiadatedEventData.getChainName(),
			"liquiadator":      liquiadatedEventData.getLiquidator(),
			"liquiadated_user": liquiadatedEventData.getLiquidateUser(),
			"ca_name":          liquiadatedEventData.getCaName(),
			"da_name":          liquiadatedEventData.getDaName(),
		},
		map[string]interface{}{},
		liquiadatedEventData.getTimestamp(),
	)
	writeAPI.WritePoint(p)
}
