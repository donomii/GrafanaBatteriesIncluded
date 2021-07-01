
REM curl -k https://dl.grafana.com/oss/release/grafana-8.0.3.windows-amd64.zip --output grafana.zip
unzip grafana.zip
mv grafana-8.0.3 grafana
cp -r ..\configs\conf grafana
mkdir grafana\data
cp ..\configs\grafana.db grafana\data
mv grafana ..


REM curl -k https://github.com/prometheus/prometheus/releases/download/v2.28.0/prometheus-2.28.0.windows-amd64.zip --output prometheus.zip
unzip prometheus.zip
mv prometheus-2.28.0.windows-amd64 prometheus
cp ..\configs\prometheus.yml prometheus
mv prometheus ..


REM curl -k https://github.com/prometheus-community/windows_exporter/releases/download/v0.16.0/windows_exporter-0.16.0-386.exe --output windowsexporter.zip
REM

mkdir ohmgraphite
cd ohmgraphite
cp ..\..\configs\OhmGraphite.exe.config .
cd ..
mv ohmgraphite ..
