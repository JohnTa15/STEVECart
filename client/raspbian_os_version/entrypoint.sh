#!/bin/bash

set -e

echo "SmartCart client started on Raspbian OS Lite. Keeping it running..."

# Update and install dependencies using apt-get (Debian/Raspbian package manager)
sudo apt-get update && sudo apt-get upgrade -y
sudo apt-get install -y build-essential pkg-config bash python3-serial libnfc-bin libnfc-dev libsqlite3-dev libmariadb-dev curl libcurl4-openssl-dev libi2c-dev wiringpi
echo "The dependencies have been installed successfully!"
sleep 1

# Installing python virtual environment
python3 -m venv /.venv
source /.venv/bin/activate

# Installing python libraries
pip3 install pn532pi
pip3 install hx711
pip3 install adafruit-circuitpython-pn532
pip3 install gpiozero
pip3 install adafruit-circuitpython-tsl2591
pip3 install pyserial


# Install Telegraf (InfluxData repository)
echo "Setting up InfluxData repository for Telegraf..."
curl -s https://repos.influxdata.com/influxdata-archive_compat.key | sudo gpg --dearmor -o /usr/share/keyrings/influxdata-archive_compat.gpg
echo "deb [signed-by=/usr/share/keyrings/influxdata-archive_compat.gpg] https://repos.influxdata.com/debian stable main" | sudo tee /etc/apt/sources.list.d/influxdata.list > /dev/null

sudo apt-get update
sudo apt-get install -y telegraf
echo "Telegraf installed successfully!"

# Install chromium and GUI display dependencies for Raspbian
sudo apt-get install -y chromium-browser xserver-xorg xserver-xorg-video-fbdev xserver-xorg-input-evdev xserver-xorg-input-libinput
echo "Display dependencies installed."

# making static ip address for the client
if [ -f "making_static_ip.sh" ]; then
    sudo bash ./making_static_ip.sh
else
    echo "making_static_ip.sh not found."
fi

flag=0
cd /STEVE-CART

echo "Verifying the setup of the client environment..."
# Checking if all Python publisher scripts are present
if [ -f "client/sensor_publisher/battery_publisher.py" ]; then
    echo "Battery publisher found!"
else
    echo "battery_publisher.py not found!"
    flag=1
fi

if [ -f "client/sensor_publisher/nfc_publisher.py" ]; then
    echo "NFC publisher found!"
else
    echo "nfc_publisher.py not found!"
    flag=1
fi

if [ -f "client/sensor_publisher/ultrasonic_publisher.py" ]; then
    echo "Ultrasonic publisher found!"
else
    echo "ultrasonic_publisher.py not found!"
    flag=1
fi

if [ -f "client/sensor_publisher/scale_publisher.py" ]; then
    echo "Scale publisher found!"
else
    echo "scale_publisher.py not found!"
    flag=1
fi

if [ -f "client/sensor_publisher/light_publisher.py" ]; then
    echo "Light publisher found!"
else
    echo "light_publisher.py not found!"
    flag=1
fi

if [ -f "client/sensor_publisher/uwb_publisher.py" ]; then
    echo "UWB publisher found!"
else
    echo "uwb_publisher.py not found!"
    flag=1
fi


if [ $flag -eq 0 ]; then
    echo "All Python sensor publishers have been verified successfully!"
else
    echo "Some sensor publishers are missing. Please check the 'client/sensor_publisher/' directory!"
    exit 1
fi

echo "Telegraf service is starting and will execute the sensor programs..."
# Make sure Telegraf uses the client configuration
if [ -f "client/telegraf/telegraf.conf" ]; then
    sudo cp client/telegraf/telegraf.conf /etc/telegraf/telegraf.conf
    sudo systemctl restart telegraf
    sudo systemctl enable telegraf
else
    echo "Warning: client/telegraf/telegraf.conf not found. Telegraf may not run the sensors automatically."
fi

# Define the IP addresses of the server nodes
MASTER_IP="192.168.136.11"
SECONDARY_IP="192.168.136.12"
THIRD_IP="192.168.136.13"

echo "Starting pinging the server nodes to check their availability..."

# Check the availability of the master node first, then the secondary and third nodes if the master is not reachable
if ping -c 2 -W 2 "$MASTER_IP" &> /dev/null; then
    echo "Master node ($MASTER_IP) is reachable."
    chromium-browser --no-sandbox --headless --disable-gpu --remote-debugging-port=9222 "http://$MASTER_IP:8080" &
    
elif ping -c 2 -W 2 "$SECONDARY_IP" &> /dev/null; then
    echo "Master down. Secondary ($SECONDARY_IP) reachable."
    chromium-browser --no-sandbox --headless --disable-gpu --remote-debugging-port=9222 "http://$SECONDARY_IP:8080" &

elif ping -c 2 -W 2 "$THIRD_IP" &> /dev/null; then
    echo "Master/Secondary down. Third ($THIRD_IP) reachable."
    chromium-browser --no-sandbox --headless --disable-gpu --remote-debugging-port=9222 "http://$THIRD_IP:8080" &

else
    echo "All nodes unreachable."
fi

tail -f /dev/null
