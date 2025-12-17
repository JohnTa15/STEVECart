#!/bin/sh

set -e 
apk update && apk add bash openssl -y
mkdir -p ssl_certss
openssl genrsa 2048 > ssl_certs/ca-key.pem
openssl req -new -x509 -nodes -days 3650 -key ssl_certs/ca-key.pem -out ssl_certs/ca-cert.pem -subj "/CN=SmartCart-Internal-CA"
openssl req -newkey rsa:2048 -days 3650 -nodes -keyout ssl_certs/server-key.pem -out ssl_certs/server-req.pem -subj "/CN=percona_cluster_node"
openssl x509 -req -in ssl_certs/server-req.pem -days 3650 \
    -CA ssl_certs/ca-cert.pem -CAkey ssl_certs/ca-key.pem -set_serial 01 \
    -out ssl_certs/server-cert.pem \
    -extensions v3_req \
    -extfile <(printf "[v3_req]\nsubjectAltName=DNS:percona_cluster_1,DNS:percona_cluster_2,DNS:percona_cluster_3,DNS:localhost,IP:127.0.0.1")

chmod 644 ssl_certs/*.pem
chmod 600 ssl_certs/*-key.pem