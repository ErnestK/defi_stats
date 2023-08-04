package influx_repositores

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type EventLogEntry struct {
	exchangeName string
	chainName    string
	timestamps   []time.Time
}

func WriteEventLog(eventLogEntry EventLogEntry) {
	client := client()
	defer client.Close()

	writeAPI := writeAPI(client)
	defer writeAPI.Flush()

	errorsCh := writeAPI.Errors()
	// Create go proc for reading and logging errors
	go logErrors(errorsCh)

	// write some points
	for _, timestamp := range eventLogEntry.timestamps {
		p := influxdb2.NewPoint(
			"event_logs",
			map[string]string{
				"exchange":   eventLogEntry.exchangeName,
				"chain_name": eventLogEntry.chainName,
			},
			map[string]interface{}{},
			timestamp)
		// write asynchronously
		writeAPI.WritePoint(p)
	}
}
