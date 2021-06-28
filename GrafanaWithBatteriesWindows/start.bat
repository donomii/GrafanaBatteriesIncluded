cd grafana\bin
start grafana-server.exe
cd ..
cd ..

cd prometheus
start prometheus.exe --config.file="prometheus.yml"
cd ..

cd windows_exporter
start windows_exporter.exe --config.file=config.yml --scrape.timeout-margin=3
cd ..

cd OhmGraphite
start OhmGraphite.exe
cd ..

sleep 5

start http://localhost:3000/d/_rxtlMk7k/system-monitor?orgId=1&kiosk=tv
