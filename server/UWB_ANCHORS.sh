
# https://github.com/Makerfabs/MaUWB_ESP32S3-with-STM32-AT-Command/blob/main/hardware/Makerfabs%20UWB%20AT%20Module%20AT%20Command%20Manual(v1.0.8).pdf

# Install python library for connecting to the UWB sensor
sudo apt-get update
sudo apt-get install -y python3-serial

# Setting up the UWB sensor TAG programmatically using an inline Python script
python3 - << 'EOF'
import serial
import time
import sys

PORT = '/dev/ttyUSB0'
BAUDRATE = 115200

def send_wait(command, ser, delay=0.5):
    print(f"Sending: {command}")
    ser.write((command + "\r\n").encode())
    time.sleep(delay)
    response = ser.read_all().decode(errors='ignore')
    print(f"Response: {response}")
    return response

try:
    ser = serial.Serial(PORT, BAUDRATE, timeout=1)
except Exception as e:
    print(f"Error opening port {PORT}: {e}")
    sys.exit(1)

# Wake up and check device
response = send_wait("AT?", ser)

if "OK" in response:
    # Save version and configuration to files
    ver = send_wait("AT+GETVER", ser)
    with open("current_ver.txt", "a") as f:
        f.write(f"VERSION:\n{ver}\n")
        
    cfg = send_wait("AT+GETCFG", ser)
    with open("current_cfg.txt", "a") as f:
        f.write(f"CONFIG:\n{cfg}\n")
        
    # Send configuration commands
    # Note: Standard AT commands typically use = rather than parentheses
    send_wait("AT+SETCFG=1,0,0,1", ser) # set mode to tag
    send_wait("AT+SETCAP=30,10,1", ser) # capacity 27 tags, 3 anchors, 10ms slot, 1 packet
    send_wait("AT+SETRPT=1", ser)       # automatic report mode active
    send_wait("AT+SAVE", ser)           # save to flash
    print("Configuration complete!")
else:
    print("Failed to get OK response from AT?")

ser.close()
EOF