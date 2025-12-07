package main

import (
	"net/http"
	"steve-api/initializers"
	"steve-api/models"

	"github.com/gin-gonic/gin"
)

// func initialize_InfluxDB() {
// 	bucket := "sensors"
// 	org := "UNIWA"
// 	token := "adminUNIWAsuper"
// 	url := "http://influxDB:8086"
// 	client := influxDB2.NewClient(url, token)
// 	write.API := client.WriteAPIBlocking(org, bucket)
// 	client.Close()
// }

// func mqtt_broker() {
// 	broker := "tcp://mqtt:1883"
// 	opts := mqtt.NewClientOptions()
// 	opts.SetClientID("go_mqtt_client")
// 	opts.SetUsername("admin")
// 	opts.SetPassword("adminUNIWA")
// 	opts.SetDefaultPublishHandler(messagePubHandler)
// 	opts.OnConnect = connectHandler
// 	opts.OnConnectionLost = connectLostHandler
// 	client := mqtt.NewClient(opts)
// 	if token := client.Connect(); token.Wait() && token.Error() != nill {
// 		panic(token.Error())
// 	}
// }

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "maou",
		})
	})
	r.Run(":8089")
	// DB := initializeDB() //mariaDB initialization
	// fmt.Printf("Connected to :", DB)
	// initialize_InfluxDB() //influxDB initialization
}

//https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
//https://github.com/eclipse-paho/paho.mqtt.golang?tab=readme-ov-file
//https://medium.com/widle-studio/integrating-mariaDB-with-golang-microservice-restful-apis-building-efficient-data-storage-4054e1588ce
//https://go.dev/doc/modules/gomod-ref
//https://docs.influxdata.com/influxDB/v2/api-guide/client-libraries/go/
//https://gorm.io/docs/associations.html#tags
//https://gemini.google.com/share/250c2d25448c
