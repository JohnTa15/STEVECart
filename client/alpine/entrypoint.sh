#!/bin/bash

set -e

echo "Alpine container started. Keeping it running..."

apk update && apk upgrade
apk add --no-cache build-base mariadb-dev python3-dev pkgconf python3 py3-pip nodejs npm mariadb mariadb-client bash libpaho-mqtt-dev
python3 -m venv venv
. venv/bin/activate
pip install --upgrade pip
pip install flask flask-mysqldb influxdb-client
pip install paho-mqtt

echo "The dependencies have been installed successfully!"

tail -f /dev/null

#Configuring PN32 (NFC MODULE)
apk add --no-cache libnfc-bin libnfc-dev
mv libnfc.conf /etc/nfc/
echo "Applying NFC changes.."
sudo reboot

# mariadb -h 192.168.10.15 -uuniwa_admin -padminUNIWA supermarket_db
# use supermarket_db;
# show tables;

# gcc battery_sensor.c mqtt_handler.c -o battery_sensor -lpaho-mqtt3c
