package influx_repo

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/ernest_k/defi_stats/lib"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func WriteTest() {
	client := influxdb2.NewClientWithOptions(
		"http://localhost:8086",
		"r_g71oNQ==",
		influxdb2.DefaultOptions(),
	)
	defer client.Close()

	org := "defi_stats"
	bucket := "prod"
	writeAPI := client.WriteAPI(org, bucket)
	defer writeAPI.Flush()

	errorsCh := writeAPI.Errors()
	// Create go proc for reading and logging errors
	go logErrors(errorsCh)

	// write some points
	for i := 0; i < 100; i++ {
		// create point
		p := influxdb2.NewPoint(
			"aave3_polygon",
			map[string]string{
				"id":       fmt.Sprintf("rack_%v", i%10),
				"vendor":   "AWS",
				"hostname": fmt.Sprintf("host_%v", i%100),
			},
			map[string]interface{}{
				"temperature": rand.Float64() * 80.0,
				"disk_free":   rand.Float64() * 1000.0,
				"disk_total":  (i/10 + 1) * 1000000,
				"mem_total":   (i/100 + 1) * 10000000,
				"mem_free":    rand.Uint64(),
			},
			time.Now())
		// write asynchronously
		writeAPI.WritePoint(p)
	}
}

func logErrors(errorsCh <-chan error) {
	for influx_err := range errorsCh {
		absPath, err := filepath.Abs("logs/influxdb_errors.txt")
		lib.Check(err)
		f, err := os.Create(absPath)
		lib.Check(err)

		_, err = f.WriteString(influx_err.Error())
		lib.Check(err)
		f.Close()
	}
}
