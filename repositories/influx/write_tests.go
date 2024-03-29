package influx

import (
	"context"
	"math/rand"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func WriteTest() {
	client := client()
	defer client.Close()

	org := "defi_stats"
	bucket := "prod"
	// writeAPI := client.WriteAPI(org, bucket)
	// Get blocking write client
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// errorsCh := writeAPI.Errors()
	// // Create go proc for reading and logging errors
	// go logErrors(errorsCh)

	// write some points
	for i := 0; i < 100; i++ {
		// create point
		p := influxdb2.NewPoint(
			"event_logs",
			map[string]string{
				"exchange":   "AAVE3",
				"chain_name": "Polygon Main",
			},
			map[string]interface{}{
				"temperature": rand.Float64() * 80.0,
				"disk_free":   rand.Float64() * 1000.0,
				"disk_total":  (i/10 + 1) * 1000000,
				"mem_free":    rand.Uint64(),
			},
			time.Now())

		writeAPI.WritePoint(context.Background(), p)
	}
	writeAPI.Flush(context.Background())
}
