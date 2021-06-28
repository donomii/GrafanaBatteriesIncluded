.Phony: install clean

LOKI_VERSION = 2.2.1
GRAFANA_VERSION = 8.0.3
PROMETHEUS_VERSION = 2.28.0
NODE_EXPORTER_VERSION = 1.1.2

clean:
	rm -r grafana node-exporter prometheus

install:
	wget https://dl.grafana.com/oss/release/grafana-${GRAFANA_VERSION}.linux-amd64.tar.gz
	mkdir -p grafana
	tar -zxvf grafana-${GRAFANA_VERSION}.linux-amd64.tar.gz --strip-components 1 -C grafana
	mkdir -p node-exporter
	wget https://github.com/prometheus/node_exporter/releases/download/v${NODE_EXPORTER_VERSION}/node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz
	tar xvfz node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz --strip-components 1 -C node-exporter
	wget https://github.com/prometheus/prometheus/releases/download/v${PROMETHEUS_VERSION}/prometheus-${PROMETHEUS_VERSION}.linux-amd64.tar.gz
	mkdir -p prometheus
	tar xvfz prometheus-${PROMETHEUS_VERSION}.linux-amd64.tar.gz --strip-components 1 -C prometheus
	mkdir -p loki
	wget https://github.com/grafana/loki/releases/download/v${LOKI_VERSION}/loki-linux-amd64.zip
	unzip loki-linux-amd64.zip -d loki
	rm *.gz
	rm *.zip

run-loki:
	pushd loki && ./loki-linux-amd64 --config.file=../conf/loki/default.yaml

run-grafana:
	pushd grafana && bin/grafana-server --config ../conf/grafana/defaults.ini

run-prometheus:
	pushd prometheus && ./prometheus --config.file=../conf/prometheus/prometheus.yml

run-node-exporter:
	sudo node-exporter/node_exporter --collector.systemd --collector.processes
