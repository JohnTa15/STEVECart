#!/bin/bash

set -e

echo "Alpine container started. Keeping it running..."

apk update && apk upgrade
apk add --no-cache build-base mariadb-dev python3-dev pkgconf python3 py3-pip nodejs npm mariadb mariadb-client bash
python3 -m venv venv
. venv/bin/activate
pip install --upgrade pip
pip install flask flask-mysqldb influxdb-client

echo "The dependencies have been installed successfully!"

tail -f /dev/null

# mariadb -h 192.168.10.15 -uuniwa_admin -padminUNIWA supermarket_db
# use supermarket_db;
# show tables;