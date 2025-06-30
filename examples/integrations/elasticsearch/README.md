# Elasticsearch Integration Example

This directory contains an example configuration for integrating Elasticsearch. This configuration is compatible with SolarWinds Observability Saas.

- `config.yaml`: Example configuration file for Elasticsearch integration.

## Prerequisites

Follow the steps below to configure Elasticsearch for monitoring.

> **Info**  
> The instructions below use the default configuration, with HTTP on port 9200. To change the default protocol or port, refer to the Elasticsearch documentation.

#### 1. Run the command to test the Elasticsearch service. The call should return basic data.

```sh
curl -X GET 'http://localhost:9200'
```

Run the command to test the Elasticsearch service when security features are enabled and authentication is required. The call should return basic data.

```sh
curl -u <username>:<password> https://localhost:9200
```
