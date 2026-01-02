package initializers

import (
	"fmt"
	"os"
	"steve-api/models"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func initMQTT(){
	var WeightData models.WeightData
	WeightData.Value = 0.0
	WeightData.StableWeight = false
	WeightData.Timestamp = time.Now()

	var NFCData models.NFCData
	NFCData.TagID = "0000"
	NFCData.ScannerID = "sc_01"

	var UltraSonicData models.UltraSonicData
	UltraSonicData.Distance = 0.0
	UltraSonicData.Timestamp_Distance = time.Now()

	var UltraSonicData2 models.UltraSonicData
	UltraSonicData2.Distance = 0.0
	UltraSonicData2.Timestamp_Distance = time.Now()

	var BatteryData models.BatteryData
	BatteryData.BatLevel = 100.0
	BatteryData.Charging = false

	var LightSensorData models.LightSensorData
	LightSensorData.LuxLevel = 0.0
	LightSensorData.Timestamp_Lux = time.Now()

	var UWBData models.UWBData
	UWBData.X_Coordinate = 0.0
	UWBData.Y_Coordinate = 0.0
	UWBData.Z_Coordinate = 0.0
	UWBData.Timestamp_UWB = time.Now()
}
func ConnectMQTT() {
	broker := os.Getenv("MQTT_BROKER")
	user := os.Getenv("MQTT_USER")
	pass := os.Getenv("MQTT_PASS")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(user)
	opts.SetPassword(pass)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message: %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
