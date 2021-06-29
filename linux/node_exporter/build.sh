#!/usr/bin/env sh

NODE_EXPORTER_VERSION=1.1.2

mkdir -p node_exporter
wget https://github.com/prometheus/node_exporter/releases/download/v${NODE_EXPORTER_VERSION}/node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz
tar xvfz node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz --strip-components 1 -C node_exporter
