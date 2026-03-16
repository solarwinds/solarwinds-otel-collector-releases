# Tail Sampling Example

This directory contains an example two-tier collector setup for tail sampling traces before exporting to SolarWinds Observability. It serves as a quick start that should be customized based on your actual deployment environment.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) with Compose

## Running

Create a `.env` file in this directory:

```
SOLARWINDS_TOKEN=<your-ingestion-token>
SOLARWINDS_ENDPOINT=<your-otel-endpoint> # e.g. otel.collector.na-01.cloud.solarwinds.com:443
```

Compose the services:
```bash
docker compose -f compose-example.yaml up
```

You can configure an instrumented application to export to the exposed ports `4317` (gRPC) or `4318` (HTTP) to try out tail sampling. The [OTEL_EXPORTER_OTLP_ENDPOINT](https://opentelemetry.io/docs/languages/sdk-configuration/otlp-exporter/#otel_exporter_otlp_endpoint) environment variable is widely supported for this. Examples:

```bash
# if the app runs directly (not containerized) on the Docker host
OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318

# if the app runs in a container on the Docker host
OTEL_EXPORTER_OTLP_ENDPOINT=http://host.docker.internal:4318
```

The example sampling policies are defined in `collector.tail-sampling.yaml`:

| Policy | Behavior |
|---|---|
| `healthcheck-policy` | Drops traces matching `/health.*` URL paths |
| `error-policy` | Always samples traces that contain an error status span |
| `high-latency-policy` | Always samples traces taking longer than 5 seconds |
| `trigger-trace-policy` | Always samples [trigger-traced](https://documentation.solarwinds.com/en/success_center/observability/content/intro/services/trigger-trace.htm) requests |
| `critical-transactions-policy` | Samples 50% of `/login` and `/payment` traces |
| `fallback-policy` | Samples 1% of all remaining traces |

## Next steps

Tail sampling requires careful configuration and monitoring. The Compose and collector YAML files in this example highlight key areas and link to relevant documentation that should be consulted in order to set up production usage in your environment.

See also the [customer documentation page](TBD) on tail sampling support in SolarWinds Observability.