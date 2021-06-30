#!/usr/bin/env sh
trap "exit" INT TERM
trap "kill 0" EXIT

"$SNAP"/exporter/bin/node_exporter > /dev/null 2>&1 &
"$SNAP"/prometheus/prometheus --storage.tsdb.path=/tmp/prometheus/data --config.file="$SNAP"/prometheus/prometheus.yml > /dev/null 2>&1 &
"$SNAP"/grafana/bin/grafana-server -config "$SNAP"/grafana/conf/defaults.ini -homepath "$SNAP"/grafana &

sleep 2
xdg-open http://localhost:3000/d/rYdddlPWk/node?orgId=1&refresh=1m

wait
