#!/bin/bash

set -e

echo "SmartCart client started on Raspbian OS Lite. Keeping it running..."

# Update and install dependencies using apt-get (Debian/Raspbian package manager)
sudo apt-get update && sudo apt-get upgrade -y
sudo apt-get install -y build-essential pkg-config bash libnfc-bin libnfc-dev libsqlite3-dev libmariadb-dev curl libcurl4-openssl-dev libi2c-dev wiringpi
echo "The dependencies have been installed successfully!"
sleep 1

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

echo "Completing the setup of the client environment..."
# Compiling the sensor programs if the source files are present
if [ -f "battery_publisher.c" ]; then
    gcc battery_publisher.c -o battery_publisher
    echo "Battery publisher compiled successfully!"
else
    echo "battery_publisher.c not found. Skipping compilation."
    flag=1
fi

if [ -f "nfc_publisher.c" ]; then
    gcc nfc_publisher.c -o nfc_publisher -lnfc
    echo "NFC publisher compiled successfully!"
else
    echo "nfc_publisher.c not found. Skipping compilation."
    flag=1
fi

if [ -f "ultrasonic_publisher.c" ]; then
    gcc ultrasonic_publisher.c -o ultrasonic_publisher -lwiringPi
    echo "Ultrasonic publisher compiled successfully!"
else
    echo "ultrasonic_publisher.c not found. Skipping compilation."
    flag=1
fi

if [ -f "weight_publisher.c" ]; then
    gcc weight_publisher.c -o weight_publisher -lsqlite3 -lwiringPi
    echo "Weight publisher compiled successfully!"
else
    echo "weight_publisher.c not found. Skipping compilation."
    flag=1
fi

if [ $flag -eq 0 ]; then
    echo "All sensor programs have been compiled successfully!"
    if [ -f "libnfc.conf" ]; then
        sudo mkdir -p /etc/nfc
        sudo cp libnfc.conf /etc/nfc/
        echo "libnfc.conf copied to /etc/nfc/"
    fi
else
    echo "Some sensor programs were not compiled due to missing source files."
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
