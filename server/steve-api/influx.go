package initializers

import (
  "context"
  "fmt"
  "time"

  "github.com/influxdata/influxdb-client-go/v2"
)

func ConnectINFLUX() {
  url := os.getEnv("INFLUX_DB")
  bucket := os.getEnv("INFLUX_BUCKET")
  token := os.getEnv("INFLUX_TOKEN")
  org := os.getEnv("INFLUX_ORG")
  username := os.getEnv("INFLUX_USERNAME")
  password := os.getEnv("INFLUX_PASSWORD")

  client := influxdb2.NewClient(url, token)
  write := client.WriteAPIBlocking(org, bucket)

  point := influxdb2.NewPoint("stat",
  map[string]string{"id"},
  map[string]float{"coordinates"},
  map[] //need fixes
  time.Now()) //need more

  writeAPI.WritePoint(context.Background(), p)
  client.Close()
}
