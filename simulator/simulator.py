#https://pypi.org/project/paho-mqtt/#client
import paho.mqtt.client as mqtt
import os
from dotenv import load_dotenv
import time
import json

load_dotenv()

mqtt_broker = os.getenv("mqtt_broker")
mqtt_user = os.getenv("username")
mqtt_pass = os.getenv("password")
port = int(os.getenv("port"))
topic_nfc = os.getenv("topic_nfc")
topic_uwb = os.getenv("topic_uwb")
topic_weight = os.getenv("topic_weight")
topic_ultrasound = os.getenv("topic_ultrasound")
cart_id = os.getenv("cart_id")
client_id = os.getenv("client_id")
transport = "tcp"


def on_log(client, userdata, level, buf):
    print(buf)

def on_connect(client, userdata, flags, rc, properties):
    if rc == 0: 
        print("Connected!")
    else: 
        print(f"Failed to connect error: ", rc)
        client.reconnect_delay_set()

def on_message(client, userdata, msg):
    print("Received message on topic", msg.topic)
    print("Message:", msg.payload.decode())
    print("----------------")

client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2, client_id)
client.username_pw_set(mqtt_user, mqtt_pass)
client.on_connect = on_connect
client.on_log = on_log
i = 0 
while True:
    try:
        client.connect(mqtt_broker, port, 60)
        break
    except Exception as e:
        i+=1
        print(f"Error while trying to connect.. {e}, retries {i}")
        time.sleep(5)
client.loop_start() 

while True:
    print("Simulating NFC")
    nfc_payload = {
        "tags" : {
            "cart_id": cart_id
        },
        "fields": {
            "NFC_data": "123451615CD" #random data from nfc tag
        }
    }
    client.publish(topic_nfc, json.dumps(nfc_payload))
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
    client.publish(topic_uwb, json.dumps(uwb_payload))
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
    client.publish(topic_weight, json.dumps(weight_payload))
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
    try:
        client.publish(topic_ultrasound, json.dumps(ultrasound_payload))
    except Exception as e:
        print(f"Error while trying to publish.. {e}")
    time.sleep(5)
