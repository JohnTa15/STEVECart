#https://pypi.org/project/paho-mqtt/#client
import paho.mqtt.client as mqtt
import os
from dotenv import load_dotenv, dotenv_values
import time
import json

load_dotenv()

mqtt_broker = os.getenv("MQTT_BROKER")
mqtt_user = os.getenv("MQTT_USER")
mqtt_pass = os.getenv("MQTT_PASS")
port = os.getenv("PORT")
topic = os.getenv("TOPIC")
cart_id = os.getenv("CART_ID")
client_id = os.getenv("CLIENT_ID")


def on_connect(client, userdata, flags, rc, properties):
    if rc == 0: 
        print("Connected!")
    else: 
        print(f"Failed to connect error: ", rc)

def on_message(client, userdata, msg):
    print("Received message on topic", msg.topic)
    print("Message:", msg.payload.decode())
    print("----------------")

client = mqtt.Client()
client.username_pw_set(mqtt_user, mqtt_pass)
client.on_connect = on_connect
client.connect(mqtt_broker, port, 60)
client.loop_forever() #if it can't communicate with the server it will keep trying infinitely

print("Simulating NFC")
nfc_payload = {
    "tags" : {
        "cart_id": cart_id
    },
    "fields": {
        "NFC_data": "123451615CD" #random data from nfc tag
    }
}
client.publish(topic, json.dumps(nfc_payload))
time.sleep(2)

print("Simulating UWB")
uwb_payload = {
    "tags" : {
        "cart_id": cart_id
    },
    "fields": {
        "range": "2.1,4.6,6.3" #range of the 3 uwb anchors 
    }
}
client.publish(topic, json.dumps(uwb_payload))
time.sleep(2)

print("Simulating weight")
weight_payload = {
    "tags" : {
        "cart_id": cart_id
    },
    "fields":{
        "weight": "5.1" #random weight
    }
}
client.publish(topic, json.dumps(weight_payload))
time.sleep(2)

print("Simulating ultrasound")
ultrasound_payload = {
    "tags" : {
        "cart_id" : cart_id
    },
    "fields": {
        "distance": "15.2" #random distance
    }
}
client.publish(topic, json.dumps(ultrasound_payload))
time.sleep(2)
client.loop_stop()
client.disconnect()
print("Simulation finished!")