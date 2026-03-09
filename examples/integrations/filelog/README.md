# Filelog Integration Example

This directory contains an example configuration for log files observability. This configuration is compatible with SolarWinds Observability SaaS.

- `config.yaml`: Example configuration file for log files integration.

> [!NOTE]  
> Default setting of integration requires `filelog.mtimeSortType` feature gate to be enabled:
> `docker run -e SOLARWINDS_TOKEN={token} -v ./config.yaml:/opt/default-config.yaml solarwinds/solarwinds-otel-collector:latest --config=/opt/default-config.yaml --feature-gates=filelog.mtimeSortType`