# REPOSITORY INSTRUCTIONS

## Repository Overview

This is a public releases repository for SolarWinds OTel Collector.
It has a private counterpart repository responsible for building of all the artifacts.
This repository only exposes the artifacts, documentation and release notes to be publicly available.
Repository also contains example configuration files for various SolarWinds integrations.


## Example Configurations
Folders ./examples/integrations/{integration_name}/ contain example configuration files for various SolarWinds integrations.

They can be used as reference when creating configuration for a specific integration.
Customers only need to fill in the required fields (lines containing `# Required parameter` comments) to get started with monitoring a specific integration.

Optional parameters or sections are commented out with `# Optional` comments.

Customers can set one of the following endpoints to send data to SolarWinds OTel Collector:
```
otel.collector.na-01.cloud.solarwinds.com:443
otel.collector.na-02.cloud.solarwinds.com:443
otel.collector.eu-01.cloud.solarwinds.com:443
otel.collector.ap-01.cloud.solarwinds.com:443
```

The API key used in the configuration must match the environment of the endpoint being used.
Config needs to be used with compatible version of SolarWinds OTel Collector binary or docker image, otherwise incompatibilities may arise.

Each integration also has a README.md file explaining the integration and any special instructions that may be needed.
