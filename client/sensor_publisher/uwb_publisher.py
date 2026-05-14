#!/usr/bin/env python3
# https://github.com/Makerfabs/MaUWB_ESP32S3-with-STM32-AT-Command/blob/main/hardware/Makerfabs%20UWB%20AT%20Module%20AT%20Command%20Manual(v1.0.8).pdf

import serial
import time
import sys
import json
from datetime import datetime

PORT = '/dev/ttyUSB0'
BAUDRATE = 115200

def send_wait(command, ser, delay=0.5, quiet=False):
    if not quiet:
        print(f"Sending: {command}", file=sys.stderr)
    ser.write((command + "\r\n").encode())
    time.sleep(delay)
    response = ser.read_all().decode(errors='ignore').strip()
    if not quiet:
        print(f"Response: {response}", file=sys.stderr)
    return response

ver = "unknown"
cfg = "unknown"

try:
    ser = serial.Serial(PORT, BAUDRATE, timeout=1)
except Exception as e:
    print(f"Error opening port {PORT}: {e}", file=sys.stderr)
    sys.exit(1)

# Wake up and check device
response = send_wait("AT?", ser)

if "OK" in response:
    ver = send_wait("AT+GETVER", ser)
    print(f"VERSION:\n{ver}\n", file=sys.stderr)
        
    cfg = send_wait("AT+GETCFG", ser)
    print(f"CONFIG:\n{cfg}\n", file=sys.stderr)
        
    # Send configuration AT commands
    send_wait("AT+SETCFG=1,0,0,1", ser) # set mode to tag
    send_wait("AT+SETCAP=30,10,1", ser) # capacity 27 tags, 3 anchors, 10ms slot, 1 packet
    send_wait("AT+SETRPT=1", ser)       # automatic report mode active
    send_wait("AT+SAVE", ser)           # save to flash
    print("Configuration complete!", file=sys.stderr)
else:
    print("Failed to get OK response from AT?", file=sys.stderr)

print("Sending information to influxDB server", file=sys.stderr)

while True:
    try:
        range_data = send_wait("AT+RANGE", ser, quiet=True)
        
        # Format and print data as JSON for Telegraf parsing
        print(json.dumps({
            "client_id": "UWB_TAG",
            "version": ver,
            "role_message": cfg,
            "range": range_data,
            "timestamp_location": datetime.now().isoformat()
        }), flush=True)
        
    except Exception as e:
        print(f"Error reading UWB: {e}", file=sys.stderr)
        
    time.sleep(5)
