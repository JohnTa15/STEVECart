package main

import (
	"net/http"
	"steve-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.PostsCreate)
}

//https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
//https://github.com/eclipse-paho/paho.mqtt.golang?tab=readme-ov-file
//https://medium.com/widle-studio/integrating-mariaDB-with-golang-microservice-restful-apis-building-efficient-data-storage-4054e1588ce
//https://go.dev/doc/modules/gomod-ref
//https://docs.influxdata.com/influxDB/v2/api-guide/client-libraries/go/
//https://gorm.io/docs/associations.html#tags
//https://gemini.google.com/share/250c2d25448c
