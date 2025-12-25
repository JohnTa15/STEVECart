
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
