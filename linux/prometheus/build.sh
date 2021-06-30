#!/usr/bin/env sh

PROMETHEUS_VERSION=2.28.0

wget https://github.com/prometheus/prometheus/releases/download/v${PROMETHEUS_VERSION}/prometheus-${PROMETHEUS_VERSION}.linux-amd64.tar.gz
mkdir -p prometheus
tar xvfz prometheus-${PROMETHEUS_VERSION}.linux-amd64.tar.gz --strip-components 1 -C prometheus
cp conf/prometheus.yml prometheus
