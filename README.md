# SolarWinds OpenTelemetry Collector

SolarWinds OpenTelemetry Collector is a distribution of OpenTelemetry Collector with components
bundled from [opentelemetry-collector](https://github.com/open-telemetry/opentelemetry-collector/tree/main)
and [opentelemetry-collector-contrib](https://github.com/open-telemetry/opentelemetry-collector-contrib). It also contains specific SolarWinds components for easier usage and enhanced telemetry collection.

## Getting Started

You will need to generate your ingestion token in SolarWinds Observability. See [API Tokens](https://documentation.solarwinds.com/en/success_center/observability/content/settings/api-tokens.htm).
Put it in the SOLARWINDS_INGESTION_TOKEN environment variable.

Then you can either use the provided example configurations or create your own custom configuration from scratch.

### Fully supported integrations

In the [examples folder](/examples/integrations/), there are a number of configurations for various fully supported integrations.

When these integrations are configured, you will get the same experience as with integrations set up by the Add Data wizards in SolarWinds Observability.
Just follow the inline comments within the example configurations for guidance.

### Custom configuration

You can also create your own custom configuration.

To get data correctly ingested by the SolarWinds Observability the configuration has to contain these components:

- [SolarWinds Extension](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/extension/solarwindsextension) - Provides basic collector identification and health check for SolarWinds Observability.
- [SolarWinds Processor](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/processor/solarwindsprocessor) - Enriches telemetry data with necessary attributes to be properly associated by SolarWinds Observability.
- [OTLP Exporter](https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/otlpexporter) - Exports telemetry data to SolarWinds Observability.

2. Create a `config.yaml` file that contains configuration for the SolarWinds OTel Collector.
   1. Insert the ingestion token and choose a correct data center (na-01, na-02, eu-01, ap-01).
   2. Specify the collector name.

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
    collector_name: # Required parameter
    grpc: &grpc_settings
      endpoint: # Required parameter
      tls:
        insecure: false
      headers: { "Authorization": "Bearer ${SOLARWINDS_TOKEN}" }

exporters:
  otlp:
    <<: *grpc_settings
```

3. Pull the SolarWinds OTel Collector from DockerHub.

```
docker pull solarwinds/solarwinds-otel-collector:playground
```

4. Start the container with your `config.yaml`.

```
docker run -v ./config.yaml:/opt/default-config.yaml solarwinds/solarwinds-otel-collector:playground
```

## Distributions

### Verified

The `verified` distribution contains only the components listed below. With the `verified` distribution, you will receive support with configuration, and the components have been tested by Solarwinds.

| Receivers                                                                                                                           | Processors                                                                                                                                     | Exporters                                           | Extensions                                                                                                                                             | Connectors                                                                                                                                 |
| :---------------------------------------------------------------------------------------------------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------- | :-------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------- |
| apachereceiver                                                                                                                      | memorylimiterprocessor                                                                                                                         | [solarwindsexporter](./exporter/solarwindsexporter) | [solarwindsextension](./extension/solarwindsextension)                                                                                                 | forwardconnector                                                                                                                           |
| prometheusreceiver                                                                                                                  | resourceprocessor                                                                                                                              | debugexporter                                       | memorylimiterextension                                                                                                                                 | routingconnector                                                                                                                           |
| dockerstatsreceiver                                                                                                                 | resourcedetectionprocessor                                                                                                                     | nopexporter                                         | healthcheckextension                                                                                                                                   | [solarwindsentityconnector](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/connector/solarwindsentityconnector) |
| elasticsearchreceiver                                                                                                               | metricstransformprocessor                                                                                                                      | otlpexporter                                        | k8sobserver                                                                                                                                            |                                                                                                                                            |
| iisreceiver                                                                                                                         | cumulativetodeltaprocessor                                                                                                                     | fileexporter                                        | [solarwindsapmsettingsextension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/solarwindsapmsettingsextension) |                                                                                                                                            |
| memcachedreceiver                                                                                                                   | deltatorate                                                                                                                                    |                                                     | filestorage                                                                                                                                            |                                                                                                                                            |
| nginxreceiver                                                                                                                       | metricsgenerationprocessor                                                                                                                     |                                                     |
| oracledbreceiver                                                                                                                    | transformprocessor                                                                                                                             |                                                     |
| otlpreceiver                                                                                                                        | filterprocessor                                                                                                                                |                                                     |
| rabbitmqreceiver                                                                                                                    | batchprocessor                                                                                                                                 |                                                     |
| redisreceiver                                                                                                                       | attributesprocessor                                                                                                                            |                                                     |
| snowflakereceiver                                                                                                                   | deltatocumulativeprocessor                                                                                                                     |                                                     |
| zookeeperreceiver                                                                                                                   | deltatorateprocessor                                                                                                                           |                                                     |
| nopreceiver                                                                                                                         | groupbyattrsprocessor                                                                                                                          |                                                     |
| filelogreceiver                                                                                                                     | groupbytraceprocessor                                                                                                                          |                                                     |
| haproxyreceiver                                                                                                                     | k8sattributesprocessor                                                                                                                         |                                                     |
| hostmetricsreceiver                                                                                                                 | [k8seventgenerationprocessor](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/processor/k8seventgenerationprocessor) |                                                     |
| journaldreceiver                                                                                                                    | [solarwindsprocessor](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/processor/solarwindsprocessor)                 |                                                     |
| k8seventsreceiver                                                                                                                   | [swok8sworkloadtypeprocessor](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/processor/swok8sworkloadtypeprocessor) |                                                     |                                                                                                                                                        |                                                                                                                                            |
| k8sobjectsreceiver                                                                                                                  |                                                                                                                                                |                                                     |                                                                                                                                                        |                                                                                                                                            |
| kafkareceiver                                                                                                                       |                                                                                                                                                |                                                     |                                                                                                                                                        |                                                                                                                                            |
| receivercreator                                                                                                                     |                                                                                                                                                |                                                     |                                                                                                                                                        |                                                                                                                                            |
| simpleprometheusreceiver                                                                                                            |                                                                                                                                                |                                                     |                                                                                                                                                        |                                                                                                                                            |
| statsdreceiver                                                                                                                      |                                                                                                                                                |                                                     |                                                                                                                                                        |                                                                                                                                            |
| [swohostmetricsreceiver](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/receiver/swohostmetricsreceiver) |                                                                                                                                                |                                                     |                                                                                                                                                        |                                                                                                                                            |
| [swok8sobjectsreceiver](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/receiver/swok8sobjectsreceiver)   |                                                                                                                                                |                                                     |                                                                                                                                                        |                                                                                                                                            |

### Playground

The `playground` distribution contains all components from `verified` distribution plus most of the components from `opentelemetry-collector-contrib` and `opentelemetry-collector` repositories. When using the playground distribution, we will not provide support with configuration. Also we cannot guaratee that all components from the mentioned repositories are working as expected.

### K8s

The `k8s` distribution contains only the components required for the Kubernetes monitoring in Solarwinds Obervability platform.

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).

```

```
