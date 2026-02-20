#!/bin/bash

set -e

echo "SmartCart client started. Keeping it running..."

apk update && apk upgrade
apk add --no-cache build-base pkgconf bash paho-mqtt-c paho-mqtt-c-dev libnfc libnfc-dev linux-headers curl-dev mariadb-dev sqlite-dev # Adding the dependencies for the sensors
echo "The dependencies have been installed successfully!"
sleep 1
apk add --no-cache chromium mesa-gl mesa-dri-gallium xorg-server xf86-video-fbdev xf86-input-evdev xf86-input-libinput # Adding the display dependencies for the GUI
echo "Dependecies installed."
#making static ip address for the client
if [ -f "making_static_ip.sh" ]; then
    sudo bash -c ./making_static_ip.sh
else
    echo "making_static_ip.sh not found."
fi
flag=0
cd /STEVE-CART

echo "Completing the setup of the client environment..."
#Compiling the sensor programs if the source files are present
if [ -f "battery_publisher.c"]; then
    gcc battery_publisher.c mqtt_handler.c -o battery_publisher -lpaho-mqtt3c
    echo "Battery publisher compiled successfully!"
else
    echo "battery_publisher.c not found. Skipping compilation."
    flag=1
fi
if [ -f "nfc_publisher.c"]; then
    gcc nfc_publisher.c mqtt_handler.c -o nfc_publisher -lpaho-mqtt3c -lnfc
    echo "NFC publisher compiled successfully!"
else
    echo "nfc_publisher.c not found. Skipping compilation."
    flag=1;
fi
if [ -f "ultrasonic_publisher.c"]; then
    gcc ultrasonic_publisher.c mqtt_handler.c -o ultrasonic_publisher -lpaho-mqtt3c
    echo "Ultrasonic publisher compiled successfully!"
else
    echo "ultrasonic_publisher.c not found. Skipping compilation."
    flag=1;
fi
if [ -f "weight_publisher.c"]; then
    gcc weight_publisher.c mqtt_handler.c -o weight_publisher -lpaho-mqtt3c -lsqlite3
    echo "Weight publisher compiled successfully!"
else
    echo "weight_publisher.c not found. Skipping compilation."
    flag=1;
fi
if [flag==0]; then
    echo "All sensor programs have been compiled successfully!"
    if [-f "libnfc.conf"]; then
        mkdir -p /etc/nfc
        cp libnfc.conf /etc/nfc/
        echo "libnfc.conf copied to /etc/nfc/"
else
    echo "Some sensor programs were not compiled due to missing source files."
    exit 1;
fi

echo "Starting the sensor programs in the background..."
# Check if the sensor executables exist before trying to run them
if [ -f "battery_sensor"]; then
    ./battery_sensor &
    echo "Battery sensor started."
else
    echo "Battery sensor executable not found. Skipping execution."
fi
if [ -f "nfc_publisher"]; then
    ./nfc_publisher &
    echo "NFC publisher started."
else
    echo "NFC publisher executable not found. Skipping execution."
fi
if [ -f "ultrasonic_publisher"]; then
    ./ultrasonic_publisher &
    echo "Ultrasonic publisher started."
else
    echo "Ultrasonic publisher executable not found. Skipping execution."
fi 
if [ -f "weight_publisher"]; then
    ./weight_publisher &
    echo "Weight publisher started."
else
    echo "Weight publisher executable not found. Skipping execution."
fi

# Define the IP addresses of the server nodes
MASTER_IP="192.168.136.11"
SECONDARY_IP="192.168.136.12"
THIRD_IP="192.168.136.13"
# TELEGRAF_CONF = "/etc/telegraf/telegraf.conf"

echo "Starting pinging the server nodes to check their availability..."

#Reverting by default to the master ip in case of any previous changes
# sed -i "s/$SECONDARY_IP/$MASTER_IP/g" "$TELEGRAF_CONF"
# sed -i "s/$THIRD_IP/$MASTER_IP/g" "$TELEGRAF_CONF"

# Check the availability of the master node first, then the secondary and third nodes if the master is not reachable
# If any of ip is not reachable, the client will try to connect to the next one. If all nodes are unreachable, it will print a message and keep running without connecting to any server.
if ping -c 2 -W 2 "$MASTER_IP" &> /dev/null; then
    echo "Master node ($MASTER_IP) is reachable."
    /usr/bin/chromium-browser --no-sandbox --headless --disable-gpu --remote-debugging-port=9222 "http://$MASTER_IP:8080"
    
elif ping -c 2 -W 2 "$SECONDARY_IP" &> /dev/null; then
    echo "Master down. Secondary ($SECONDARY_IP) reachable."
    # sed -i "s/$MASTER_IP/$SECONDARY_IP/g" "$TELEGRAF_CONF"
    /usr/bin/chromium-browser --no-sandbox --headless --disable-gpu --remote-debugging-port=9222 "http://$SECONDARY_IP:8080"

elif ping -c 2 -W 2 "$THIRD_IP" &> /dev/null; then
    echo "Master/Secondary down. Third ($THIRD_IP) reachable."
    # sed -i "s/$MASTER_IP/$THIRD_IP/g" "$TELEGRAF_CONF"
    /usr/bin/chromium-browser --no-sandbox --headless --disable-gpu --remote-debugging-port=9222 "http://$THIRD_IP:8080"

else
    echo "All nodes unreachable."
fi

#~~~Missing the GUI part~~~

tail -f /dev/null


# mariadb -h 192.168.10.15 -uuniwa_admin -padminUNIWA supermarket_db
# use supermarket_db;
# show tables;

# gcc battery_sensor.c mqtt_handler.c -o battery_sensor -lpaho-mqtt3c
