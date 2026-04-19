package initializers

import (
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var InfluxClient influxdb2.Client

// ConnectINFLUX should be called once in your main.go when the server starts
func ConnectINFLUX() {
	// 1. Grab connection credentials from environment
	url := os.Getenv("INFLUX_DB")
	token := os.Getenv("INFLUX_TOKEN")

	// 2. Initialize the client and save it to the global variable
	InfluxClient = influxdb2.NewClient(url, token)

	log.Println("Successfully initialized global InfluxDB client")
	//telegraf handles the writing of the metrics, we just read from it
}
