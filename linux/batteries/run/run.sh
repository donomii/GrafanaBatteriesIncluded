#!/usr/bin/env sh
trap "exit" INT TERM
trap "kill 0" EXIT

[ -e "$SNAP_USER_DATA/grafana" ] || mkdir "$SNAP_USER_DATA"/grafana
[ -e "$SNAP_USER_DATA/prometheus" ] || mkdir "$SNAP_USER_DATA"/prometheus

"$SNAP"/exporter/bin/node_exporter --collector.systemd --collector.processes > /dev/null 2>&1 &
"$SNAP"/prometheus/prometheus  --web.listen-address="127.0.0.1:9090" --storage.tsdb.path="$SNAP_USER_DATA"/prometheus/data --config.file="$SNAP"/prometheus/prometheus.yml > /dev/null 2>&1 &
"$SNAP"/grafana/bin/grafana-server -config "$SNAP"/grafana/conf/defaults.ini -homepath "$SNAP"/grafana &

sleep 5
xdg-open http://localhost:3000/d/rYdddlPWk/node?orgId=1&refresh=1m

wait
