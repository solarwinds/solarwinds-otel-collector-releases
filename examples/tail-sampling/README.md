# Tail Sampling Example

This directory contains an example two-tier collector setup for tail sampling traces before exporting to SolarWinds Observability. It serves as a quick start that should be customized based on your actual deployment environment.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) with Compose

## Running

Create a `.env` file in this directory:

```
SOLARWINDS_TOKEN=<your-ingestion-token>
SOLARWINDS_ENDPOINT=<your-otel-endpoint>
```

Compose the services:
```bash
docker compose -f compose-example.yaml up
```

You can configure an instrumented application to export to the exposed ports `4317` (gRPC) or `4318` (HTTP) to try out tail sampling. The example sampling policies are defined in `collector.tail-sampling.yaml`:

| Policy | Behavior |
|---|---|
| `healthcheck-policy` | Drops traces matching `/health.*` URL paths |
| `error-policy` | Always samples traces containing an error span |
| `high-latency-policy` | Always samples traces taking longer than 5 seconds |
| `trigger-trace-policy` | Always samples trigger-traced requests |
| `critical-transactions-policy` | Samples 50% of `/login` and `/payment` traces |
| `fallback-policy` | Samples 1% of all remaining traces |

The compose and collector configuration YAML files highlight key areas and link to relevant documentation.
