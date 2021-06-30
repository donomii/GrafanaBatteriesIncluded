#!/usr/bin/env sh

LOKI_VERSION=2.2.1

mkdir -p loki
wget https://github.com/grafana/loki/releases/download/v${LOKI_VERSION}/loki-linux-amd64.zip
unzip loki-linux-amd64.zip -d loki
