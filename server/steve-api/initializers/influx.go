package initializers

import (
	"context"
	"fmt"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var InfluxClient influxdb2.Client

func ConnectINFLUX() influxdb2.Client {
	url := os.Getenv("INFLUX_DB")
	bucket := os.Getenv("INFLUX_BUCKET")
	token := os.Getenv("INFLUX_TOKEN")
	org := os.Getenv("INFLUX_ORG")
	username := os.Getenv("INFLUX_USERNAME")
	password := os.Getenv("INFLUX_PASSWORD")

	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)


	tags := map[string]string{
		"cart_id":       cartID,
		"store_section": section,
		"status":        "active",
	}

	fields := map[string]interface{}{
		"temperature": temp,
		"weight":      weight,
		"battery":     battery,
		"coordinates": coordinates,
	}

	point := influxdb2.NewPoint(
    "cart_metrics",
		tags,
		fields,
		time.Now())

  writeAPI.WritePoint(point)
	err := writeAPI.WritePoint(context.Background(), point)
  if err != nil {
    return fmt.Errorf("Can't write to influxdb: %w", err)
  }

  fmt.Printf("Telemetry success!"); 
	client.Close()
  return client
}
