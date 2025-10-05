#!/bin/sh
set -e

mkdir -p /etc/grafana/provisioning/datasources

cat > /etc/grafana/provisioning/datasources/datasource.yml <<'EOF'
apiVersion: 1
datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    url: http://prometheus:9090
    isDefault: true

  - name: InfluxDB
    type: influxdb
    access: proxy
    url: http://influxdb3:8086
    user: admin
    secureJsonData:
      password: adminUNIWA
    jsonData:
      dbName: sensors
EOF

/run.sh
