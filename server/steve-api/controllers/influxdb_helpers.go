package controllers

import (
	"context"
	"fmt"
	"os"
	"steve-api/initializers"

)

var org = os.Getenv("INFLUXDB_ORG")
var bucket = os.Getenv("INFLUXDB_BUCKET")
var measurement_name = "cart_metrics"

//making a file in order to get the data from the influxdb
//now with this I can directly use the file in the controllers..
func GetNFC(cartID string) (string, error){
	queryAPI := initializers.InfluxClient.QueryAPI(org)

	query := fmt.Sprintf(`from(bucket: "%s")
		|> range(start: -1m)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["cart_id"] == "%s")
		|> filter(fn: (r) => r["_field"] == "nfc")
		|> last()
	`, bucket, measurement_name, cartID)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return "", err
	}

	var nfc string
	for result.Next() {
		if val, ok := result.Record().Value().(string); ok {
			nfc = val
		}
	}

	return nfc, nil
}

func GetWeight(cartID string) (float64, error){
	queryAPI := initializers.InfluxClient.QueryAPI(org)

	query := fmt.Sprintf(`from(bucket: "%s")
		|> range(start: -1m)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["cart_id"] == "%s")
		|> filter(fn: (r) => r["_field"] == "weight")
		|> last()
	`, bucket, measurement_name, cartID)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return 0, err
	}

	var weight float64
	for result.Next() {
		if val, ok := result.Record().Value().(float64); ok {
			weight = val
		}
	}

	return weight, nil
}


func GetLux(cartID string) (float64, error){
	queryAPI := initializers.InfluxClient.QueryAPI(org)

	query := fmt.Sprintf(`from(bucket: "%s")
		|> range(start: -1m)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["cart_id"] == "%s")
		|> filter(fn: (r) => r["_field"] == "lux")
		|> last()
	`, bucket, measurement_name, cartID)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return 0, err
	}

	var lux float64
	for result.Next() {
		if val, ok := result.Record().Value().(float64); ok {
			lux = val
		}
	}

	return lux, nil
}
