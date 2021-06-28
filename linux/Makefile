.Phony: install clean

GRAFANA_VERSION = 8.0.3
GRAFANA_ROOT = grafana

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
	rm *.gz

run-grafana:
	pushd "${GRAFANA_ROOT}" && bin/grafana-server --config ../conf/grafana/defaults.ini

run-prometheus:
	prometheus/prometheus --config.file=conf/prometheus/prometheus.yml

run-node-exporter:
	node-exporter/node_exporter
