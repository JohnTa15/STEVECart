package main

import (
	"database/sql"
	"go/token"
	"log"

	mqtt "github.com/eclipse-paho/paho.mqtt.golang"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var db *sql.DB

func initializeDB() *sql.DB {
	connectionURL := "uniwa_admin:adminUNIWA@tcp(mariadb:3306)/supermarket_db"
	db, err := sql.Open("mysql", connectionURL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initialize_InfluxDB() {
	bucket := "sensors"
	org := "UNIWA"
	token := "adminUNIWAsuper"
	url := "http://influxdb:8086"
	client := influxdb2.NewClient(url, token)
	write.API := client.WriteAPIBlocking(org, bucket)
	client.Close()
}

func mqtt_broker(){
	var broker = "tcp://mqtt:1883"
	opts := mqtt.NewClientOptions()
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("admin")
	opts.SetPassword("adminUNIWA")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nill {
		panic(token.Error())
	}
}

func main() {
	db = initializeDB()   //mariaDB initialization
	initialize_InfluxDB() //influxDB initialization

	defer db.Close()
}

//https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
//https://github.com/eclipse-paho/paho.mqtt.golang?tab=readme-ov-file
//https://medium.com/widle-studio/integrating-mariadb-with-golang-microservice-restful-apis-building-efficient-data-storage-4054e1588ce
//https://go.dev/doc/modules/gomod-ref
//https://docs.influxdata.com/influxdb/v2/api-guide/client-libraries/go/