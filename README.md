# SolarWinds OpenTelemetry Collector

SolarWinds OpenTelemetry Collector (swotelcol) is a distribution of OpenTelemetry Collector with components
bundled from [opentelemetry-collector], [opentelemetry-collector-contrib] and [solarwinds-otel-collector-contrib].
It contains SolarWinds-specific components for better integration with SolarWinds Observability (SWO) and enhances telemetry collection.

[opentelemetry-collector]: https://github.com/open-telemetry/opentelemetry-collector
[opentelemetry-collector-contrib]: https://github.com/open-telemetry/opentelemetry-collector-contrib
[solarwinds-otel-collector-contrib]: https://github.com/solarwinds/solarwinds-otel-collector-contrib

## Getting Started

You will need to generate your ingestion token in SWO. See [API Tokens](https://documentation.solarwinds.com/en/success_center/observability/content/settings/api-tokens.htm).
Put it in the SOLARWINDS_TOKEN environment variable.

```sh
# Unix shell
export SOLARWINDS_TOKEN="<your-ingestion-token>"
```

```ps1
# PowerShell
$env:SOLARWINDS_TOKEN="<your-ingestion-token>"
```

Then you can either use the provided example configurations or create your own custom configuration from scratch.

### Fully supported integrations

In the [examples folder](/examples/integrations/), there are a number of configurations for various fully supported integrations.

When these integrations are configured, you will get the same experience as with integrations set up by the Add Data wizards in SWO.

Just follow the inline comments within the example configurations for guidance.

### Custom configuration

You can also create your own custom configuration.

To get data correctly ingested by the SWO the configuration has to contain these components:

- [SolarWinds Extension](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/extension/solarwindsextension) - Provides basic collector identification and health check for SWO.
- [SolarWinds Processor](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/processor/solarwindsprocessor) - Enriches telemetry data with attributes to be properly associated by SWO.
- [OTLP Exporter](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/otlpexporter) - Exports telemetry data to SWO.

Create a `config.yaml` file that contains configuration for the swotelcol.

1.  Set ingestion endpoint. To get correct endpoint, search for OTLP in [Data centers and endpoint URIs](https://documentation.solarwinds.com/en/success_center/observability/content/system_requirements/endpoints.htm) docs.
2.  Specify the collector name.

```yaml
service:
  extensions: [solarwinds]
  pipelines:
    metrics:
      receivers: [redis]
      exporters: [otlp]

receivers:
  redis:
    endpoint: "<redis-url>:6379"
    collection_interval: 10s
    password: ${env:REDIS_PASSWORD}

extensions:
  solarwinds:
    collector_name: "<collector-name>" # Required parameter
    grpc: &grpc_settings
      endpoint: "<endpoint>" # Required parameter
      tls:
        insecure: false
      headers: { "Authorization": "Bearer ${env:SOLARWINDS_TOKEN}" }

exporters:
  otlp:
    <<: *grpc_settings
```

### Running the Collector

1. Pull the swotelcol image from DockerHub (Verified [distribution](#distributions) in this case).

```sh
docker pull solarwinds/solarwinds-otel-collector:verified
```

2. Start the container with your `config.yaml`.

```sh
docker run -v ./config.yaml:/opt/default-config.yaml solarwinds/solarwinds-otel-collector:verified
```

## Distributions

### Verified
The `verified` distribution contains only the components listed below. With the `verified` distribution, you will receive support with configuration, and the components have been tested by Solarwinds.

Full set of components is available in the [verified-components](./docs/verified-components.md).

### Playground
The `playground` distribution contains all components from `verified` distribution plus most of the components from `opentelemetry-collector-contrib` and `opentelemetry-collector` repositories. When using the playground distribution, we will not provide support with configuration. Also we cannot guaratee that all components from the mentioned repositories are working as expected.

Full set of components is available in the [playground-components](./docs/playground-components.md).

### K8s
The `k8s` distribution contains only the components required for the Kubernetes monitoring in Solarwinds Obervability platform.

Full set of components is available in the [k8s-components](./docs/k8s-components.md)

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).
