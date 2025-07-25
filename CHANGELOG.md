# Changelog

## vNext

## v0.127.7
Reintroduce components missing in `v0.127.6`:
- [pprofextension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/pprofextension) (k8s distribution)
- [logdedupprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/logdedupprocessor) (k8s distribution)
- [zpagesextension](https://github.com/open-telemetry/opentelemetry-collector/tree/main/extension/zpagesextension) (verified distribution)
- [solarwindsprocessor](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/processor/solarwindsprocessor) (all distributions)

## v0.127.6
- Updated GHA deploy workflow to verify existing Docker images instead of building new ones.
- Removed Docker image building and pushing steps from GHA deploy workflow.

## v0.127.5
- Added `logdedupprocessor` to K8s distribution.
- Added README files to [integration examples](./examples/integrations).
- Updated [README file](README.md) with information about supported integrations.
- Updated E2E tests to use new OTLP exporter and SolarWinds processor.
- Consumes [solarwinds-otel-collector-contrib](https://github.com/solarwinds/solarwinds-otel-collector-contrib) `v0.127.5` dependencies - [full changelog](https://github.com/solarwinds/solarwinds-otel-collector-contrib/blob/main/CHANGELOG.md#v01275)
- Fixed GHSA-fv92-fjc5-jj9h: mapstructure may leak sensitive information in logs when processing malformed data.

## v0.127.4
- Add [pprofextension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/pprofextension) to the `k8s` distribution.
- Add [logdedupprocessor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/logdedupprocessor) to the `k8s` distribution.
- Add [solarwindsprocessor](https://github.com/solarwinds/solarwinds-otel-collector-contrib/tree/main/processor) to all distributions.
- Consumes [solarwinds-otel-collector-contrib](https://github.com/solarwinds/solarwinds-otel-collector-contrib) `v0.127.4` dependencies - [full changelog](https://github.com/solarwinds/solarwinds-otel-collector-contrib/blob/main/CHANGELOG.md#v01274)

## v0.127.3
- Consumes [solarwinds-otel-collector-contrib](https://github.com/solarwinds/solarwinds-otel-collector-contrib) `v0.127.3` dependencies - [full changelog](https://github.com/solarwinds/solarwinds-otel-collector-contrib/blob/main/CHANGELOG.md#v01273)
- Fixed CVE-2025-22874: Calling Verify with a VerifyOptions.KeyUsages that contains ExtKeyUsageAny unintentionally disabledpolicy validation. This only affected certificate chains which contain policy graphs, which are rather uncommon.

## v0.127.2
- Consumes [solarwinds-otel-collector-contrib](https://github.com/solarwinds/solarwinds-otel-collector-contrib) `v0.127.2` dependencies - [full changelog](https://github.com/solarwinds/solarwinds-otel-collector-contrib/blob/main/CHANGELOG.md#v01272)
- Added examples of [integration templates](./examples/integrations) compatible with SolarWinds Observability SaaS.

## v0.127.1
- Consumes [solarwinds-otel-collector-contrib](https://github.com/solarwinds/solarwinds-otel-collector-contrib) `v0.127.1` dependencies - [full changelog](https://github.com/solarwinds/solarwinds-otel-collector-contrib/blob/main/CHANGELOG.md#v01271)

## v0.127.0
- Consumes OpenTelemetry Collector dependencies v0.127.0.
- Release process utilizes builder to build the various distributions.
- Dependencies bumped to latest available versions.

## v0.123.7
- Add `routingconnector` to the k8s distribution
- Add `solarwindsentityconnector` to the k8s distribution
- Consumed components from [solarwinds-otel-collector-contrib](https://github.com/solarwinds/solarwinds-otel-collector-contrib) repository. Moved components were removed from this repository but there should be no functional changes.

## v0.123.6
- No changes, previous release failed to finish properly.

## v0.123.5
- No changes, previous release failed to finish properly.

## v0.123.4
- Fixing issues with release of windows docker images.

## v0.123.3
- Publishing `verified` and `playground` docker images.

## v0.123.2
- Moved connection-check code to separate binary. Binary is added to k8s docker images.
- Adds [SolarWinds Kubernetes Workload Type Processor](./processor/swok8sworkloadtypeprocessor/README.md) for annotating metrics with a k8s workload type based on their attributes.

## v0.123.1
- Fix CVE-2025-22872: golang.org/x/net vulnerable to Cross-site Scripting

## v0.123.0
- Consumes OpenTelemetry Collector dependencies v0.123.0.
- SolarWinds exporter is now reported as otlp/solarwinds-<name> in collector's telemetry.

## v0.119.12
- Updates non-opentelemetry dependencies to latest possible version
- Sets metrics scope name to `github.com/solarwinds/solarwinds-otel-collector-releases`

## v0.119.11
- Fix CVE-2025-27144: Uncontrolled Resource Consumption

## v0.119.10
- Ignores any timestamps in all the Kubernetes manifests
- Fix CVE-2025-30204

## v0.119.9
- Fix CVE-2025-29786

## v0.119.8
- Adds ap-01 datacell support

## v0.119.7
- Fix CVE-2025-22866

## v0.119.6
- Updates Go build toolchain to 1.23.6

## v0.119.5
- Updating `swok8sobjectsreceiver` to remove `managedFields` for both PULL and WATCH objects

## v0.119.4
- Updating `swok8sobjectsreceiver` to report changes in other than `status`, `spec`, and `metadata` sections

## v0.119.3
- Adds custom `k8sobjectsreceiver` to notify about what sections in manifest were changed 

## v0.119.2
- Fix CVE-2025-22869
- Fix CVE-2025-22868

## v0.119.1
- Utilizes `pdatatest` for specific E2E tests.
- SolarWinds-specific packages are tagged and can be referenced from other repositories.
- Adds custom `k8seventgenerationprocessor` to transform K8S entities change events to logs.
- Removes opentelemetry-collector wrapper used for corruption prevention as newly introduced `fsync@v0.96.0` solves the issue.

## v0.119.0
- Consumes OpenTelemetry Collector dependencies v0.119.0.

## v0.113.8
- Updates Go build toolchain to 1.23.5.
- Adds [SWO Host Metrics Receiver](./receiver/swohostmetricsreceiver/README.md) for additional Host metrics monitoring.
- Adds connection check functionality to K8s distribution startup.
- Adds Windows architecture for Docker builds.

## v0.113.7
- Adds `without_entity` to [SolarWinds Extension](./extension/solarwindsextension/README.md#getting-started) configuration, so users can opt out of collector entity creation.
- Tags all signals with `entity_creation` attribute, except when without_entity is set on [SolarWinds Extension](./extension/solarwindsextension/README.md#getting-started).

## v0.113.6
- Marks all outgoing telemetry from the [SolarWinds Exporter](./exporter/solarwindsexporter) with
an attribute storing the collector name (`sw.otelcol.collector.name`) as it is configured in the
[SolarWinds Extension](./extension/solarwindsextension/README.md#getting-started).
- The uptime metric used to signal heartbeat is now decorated with `sw.otelcol.collector.version` which contains collector version.

## v0.113.5
Tags released docker images with `latest` tag.

## v0.113.4
Adds optional `resource` configuration parameter for [SolarWinds Extension](./extension/solarwindsextension).

## v0.113.3
Removes `insecure` testing configuration parameter for [SolarWinds Extension](./extension/solarwindsextension).

## v0.113.2
Fixes OTLP port number used for exporting telemetry.

## v0.113.1
Adds [SolarWinds Extension](./extension/solarwindsextension). The [SolarWinds Exporter](./exporter/solarwindsexporter) is now dependent on the extension.

## v0.113.0
Initial version of SolarWinds OpenTelemetry Collector.
The collector provides all available components (receivers, processors, exporters, connectors, providers)
from [opentelemetry-collector](https://github.com/open-telemetry/opentelemetry-collector/tree/v0.113.0) (version `v0.113.0`) and [opentelemetry-collector-contrib](https://github.com/open-telemetry/opentelemetry-collector/tree/v0.113.0) (version `v0.113.0`).

### Additional details:
- `solarwindsexporter` has been added to easily integrate with **SolarWinds Observability SaaS**. Please read its [documentation](exporter/solarwindsexporter/README.md) to learn more.
