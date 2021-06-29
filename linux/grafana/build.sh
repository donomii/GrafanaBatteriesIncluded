#!/usr/bin/env sh
GRAFANA_VERSION=8.0.3

wget https://dl.grafana.com/oss/release/grafana-${GRAFANA_VERSION}.linux-amd64.tar.gz
mkdir -p grafana
tar -zxvf grafana-${GRAFANA_VERSION}.linux-amd64.tar.gz --strip-components 1 -C grafana
cp -R conf grafana
