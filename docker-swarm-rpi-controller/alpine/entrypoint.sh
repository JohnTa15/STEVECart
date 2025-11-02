#!/bin/bash

set -e

echo "Alpine container started. Keeping it running..."

apk update && apk upgrade
apk add python3 py3-pip nodejs npm
apk add --no-cache build-base mariadb-dev python3-dev pkgconf
python3 -m venv venv
. venv/bin/activate
pip install --upgrade pip
pip install flask flask-mysqldb influxdb-client

echo "The dependencies have been installed successfully!"