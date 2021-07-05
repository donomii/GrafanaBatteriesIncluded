#!/bin/sh

tar -zxvf grafana-8.0.3.darwin-amd64.tar.gz
mv grafana-8.0.3 grafana-server
cp -r configs/conf grafana-server
mkdir grafana-server/data
cp configs/grafana.db grafana-server/data

tar -xzvf prometheus-2.28.0.darwin-amd64.tar.gz
mv prometheus-2.28.0.darwin-amd64 prometheus
cp configs/prometheus.yml prometheus

tar -xzvf node_exporter-1.1.2.darwin-amd64.tar.gz
mv node_exporter-1.1.2.darwin-amd64/ node_exporter

cd splash-screen && go build .
cd ..

cd process-exporter && go build .
