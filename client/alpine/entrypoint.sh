#!/bin/bash

set -e

echo "SmartCart client started. Keeping it running..."

apk update && apk upgrade
apk add --no-cache build-base pkgconf bash libpaho-mqtt-dev libnfc libnfc-dev linux-headers curl-dev mariadb-dev sqlite-dev # Adding the dependencies for the sensors
echo "The dependencies have been installed successfully!"
sleep 1
apk add --no-cache chromium mesa-gl mesa-dri-gallium xorg-server xf86-video-fbdev xf86-input-evdev xf86-input-mouse xf86-input-keyboard # Adding the display dependencies for the GUI
echo "Dependecies installed."
flag = 0

echo "Completing the setup of the client environment..."
#Compiling the sensor programs if the source files are present
if[ -f "battery_sensor.c"]; then
    gcc battery_sensor.c mqtt_handler.c -o battery_sensor -lpaho-mqtt3c
    echo "Battery sensor compiled successfully!"
else
    echo "battery_sensor.c not found. Skipping compilation."
    flag = 1;
fi
if[ -f "nfc_publisher.c"]; then
    gcc nfc_publisher.c mqtt_handler.c -o nfc_publisher -lpaho-mqtt3c -lnfc
    echo "NFC publisher compiled successfully!"
else
    echo "nfc_publisher.c not found. Skipping compilation."
    flag = 1;
fi
if[ -f "ultrasonic_publisher.c"]; then
    gcc ultrasonic_publisher.c mqtt_handler.c -o ultrasonic_publisher -lpaho-mqtt3c
    echo "Ultrasonic publisher compiled successfully!"
else
    echo "ultrasonic_publisher.c not found. Skipping compilation."
    flag = 1;
fi
if[ -f "weight_publisher.c"]; then
    gcc weight_publisher.c mqtt_handler.c -o weight_publisher -lpaho-mqtt3c -lsqlite3
    echo "Weight publisher compiled successfully!"
else
    echo "weight_publisher.c not found. Skipping compilation."
    flag = 1;
fi
if[flag == 0]; then
    echo "All sensor programs have been compiled successfully!"
    if[-f "libnfc.conf"]; then
        mkdir -p /etc/nfc
        cp libnfc.conf /etc/nfc/
        echo "libnfc.conf copied to /etc/nfc/"
else
    echo "Some sensor programs were not compiled due to missing source files."
    exit 1;
fi

echo "Starting the sensor programs in the background..."
# Check if the sensor executables exist before trying to run them
if[ -f "battery_sensor"]; then
    ./battery_sensor &
    echo "Battery sensor started."
else
    echo "Battery sensor executable not found. Skipping execution."
fi
if[ -f "nfc_publisher"]; then
    ./nfc_publisher &
    echo "NFC publisher started."
else
    echo "NFC publisher executable not found. Skipping execution."
fi
if[ -f "ultrasonic_publisher"]; then
    ./ultrasonic_publisher &
    echo "Ultrasonic publisher started."
else
    echo "Ultrasonic publisher executable not found. Skipping execution."
fi
if[ -f "weight_publisher"]; then
    ./weight_publisher &
    echo "Weight publisher started."
else
    echo "Weight publisher executable not found. Skipping execution."
fi

# Define the IP addresses of the server nodes
MASTER_IP = "192.168.136.11"
SECONDARY_IP = "192.168.136.12"
THIRD_IP = "192.168.136.13"

echo "Starting pinging the server nodes to check their availability..."

# Check the availability of the master node first, then the secondary and third nodes if the master is not reachable
if ping -c 4 $MASTER_IP &> /dev/null; then
    echo "Master node ($MASTER_IP) is reachable."
    /usr/bin/chromium-browser --no-sandbox --headless --disable-gpu --remote-debugging-port=9222 "http://192.168.136.11:8080"
else
    echo "Master node ($MASTER_IP) is not reachable. Checking the secondary node..."
    if ping -c 4 $SECONDARY_IP &> /dev/null; then
        echo "Secondary node ($SECONDARY_IP) is reachable."
        /usr/bin/chromium-browser --no-sandbox --headless --disable-gpu --remote-debugging-port=9222 "http://192.168.136.12:8080"
    else
        echo "Secondary node ($SECONDARY_IP) is not reachable. Checking the third node..."
        if ping -c 4 $THIRD_IP &> /dev/null; then
            echo "Third node ($THIRD_IP) is reachable."
            /usr/bin/chromium-browser --no-sandbox --headless --disable-gpu --remote-debugging-port=9222 "http://192.168.136.13:8080"
        else
            echo "All server nodes are unreachable."
        fi
    fi
fi

#~~~Missing the GUI part~~~

tail -f /dev/null


# mariadb -h 192.168.10.15 -uuniwa_admin -padminUNIWA supermarket_db
# use supermarket_db;
# show tables;

# gcc battery_sensor.c mqtt_handler.c -o battery_sensor -lpaho-mqtt3c
