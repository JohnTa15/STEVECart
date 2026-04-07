package initializers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"steve-api/models"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var payloadData map[string]interface{}

	err := json.Unmarshal(msg.Payload(), &payloadData)
	if err != nil {
		fmt.Println("Can't read JSON payload: ", err)
		return
	}

	if _, ok := payloadData["tag_id"]; ok {
		fmt.Println("NFC Scan Detected!")
	} else if _, ok := payloadData["weight"]; ok {
		fmt.Println("Weight Detected!")
	} else if _, ok := payloadData["distance"]; ok {
		fmt.Println("Distance Detected!")
	} else if _, ok := payloadData["lux"]; ok {
		fmt.Println("Lux Detected!")
	} else if _, ok := payloadData["x_coordinate"]; ok {
		fmt.Println("UWB Detected!")
	} else if _, ok := payloadData["battery_level"]; ok {
		fmt.Println("Battery Detected!")
	} else {
		fmt.Printf("Unknown Payload Format: %s\n", msg.Payload())
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func initMQTT() {
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

func ConnectMQTT() mqtt.Client {
	broker := os.Getenv("MQTT_BROKER")
	user := os.Getenv("MQTT_USER")
	pass := os.Getenv("MQTT_PASS")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(user)
	opts.SetPassword(pass)

	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var result map[string]interface{}

	json.Unmarshal(byteValue, &result)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message: %d", i)
		token := client.Publish("sensors/cart_data", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func Sub(client mqtt.Client) {
	topic := "sensors/cart_data"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
