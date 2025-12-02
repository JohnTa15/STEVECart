#include <ESP8266WiFi.h>
#include <PubSubClient.h>

// ====== WiFi Settings ======
const char* ssid = "Cudy-Home";
const char* password = "@EoNpQXQV4chE$L";

// ====== MQTT Broker Settings ======
const char* mqtt_server = "192.168.10.15";
const int mqtt_port = 1883;
const char* mqtt_user = "admin";
const char* mqtt_pass = "adminUNIWA";

// ====== Pins ======
int pirPin = 5;   // D1 = GPIO5
int ledPin = 4;   // D2 = GPIO4

// ====== Global Objects ======
WiFiClient espClient;
PubSubClient client(espClient);

// ====== WiFi Connection ======
void setup_wifi() {
  delay(10);
  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);

  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("\nWiFi connected!");
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());
}

// ====== MQTT Reconnect ======
void reconnect() {
  while (!client.connected()) {
    Serial.print("Connecting to MQTT broker...");
    if (client.connect("ESP8266-PIR", mqtt_user, mqtt_pass)) {
      Serial.println("connected!");
    } else {
      Serial.print("failed, rc=");
      Serial.print(client.state());
      Serial.println(" — retrying in 5s");
      delay(5000);
    }
  }
}

void setup() {
  Serial.begin(115200);
  pinMode(pirPin, INPUT);
  pinMode(ledPin, OUTPUT);
  
  Serial.println("PIR Sensor Active...");
  setup_wifi();
  client.setServer(mqtt_server, mqtt_port);
}

void loop() {
  if (!client.connected()) {
    reconnect();
  }
  client.loop();

  int motion = digitalRead(pirPin);

  if (motion == HIGH) {
    Serial.println("Κίνηση εντοπίστηκε!");

    digitalWrite(ledPin, HIGH);
    client.publish("home/pir/status", "ACTIVE"); //ACTIVE
  } else {
    digitalWrite(ledPin, LOW);
    client.publish("home/pir/status", "CLEAR"); //CLEAR
  }

  delay(1000);
}
