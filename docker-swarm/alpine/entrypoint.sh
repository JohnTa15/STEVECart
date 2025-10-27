#!/bin/bash

set -e

echo "Alpine container started. Keeping it running..."

apt update && apt upgrade -y
apt install python3 python3-pip -y
apt install nodejs npm -y
pip3 install --upgrade pip
pip install flask flask-mysqldb influxdb-client

echo "The dependencies have been installed successfully!"