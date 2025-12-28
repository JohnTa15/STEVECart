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
  var data models.WeightData
  var data models.NFCData
  var data models.UltraSonicData
  var data models.BatteryData
  var data models.LightSensorData
  var data models.UWBData
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
