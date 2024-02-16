package influx

import (
	"os"
	"path/filepath"

	"github.com/ernest_k/defi_stats/lib"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Connection struct {
	client   influxdb2.Client
	writeApi api.WriteAPI
}

func CreateDB() *Connection {
	client := influxdb2.NewClientWithOptions(
		"http://localhost:8086",
		"r_g71oNQ==",
		influxdb2.DefaultOptions(),
	)

	writeAPI := writeAPI(client)
	errorsCh := writeAPI.Errors()
	go logErrors(errorsCh)

	return &Connection{
		client:   client,
		writeApi: writeAPI,
	}
}

func (influxConnection *Connection) CloseDB() {
	influxConnection.writeApi.Flush()
	influxConnection.client.Close()
}

func client() influxdb2.Client {
	return influxdb2.NewClientWithOptions(
		"http://localhost:8086",
		"r_g71oNQ==",
		influxdb2.DefaultOptions(),
	)
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

func writeAPI(client influxdb2.Client) api.WriteAPI {
	org := "defi_stats"
	bucket := "prod"
	return client.WriteAPI(org, bucket)
}
