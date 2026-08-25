# Windows Event Logs Integration Example

This directory contains an example configuration for Windows Event Log observability. This configuration is compatible with SolarWinds Observability SaaS.

- `config.yaml`: Example configuration file for Windows Event Logs integration.

> [!NOTE]
> The `windows_event_log` receiver is Windows-only; the published image tags are multiarch manifests that include Windows nanoserver builds. Run the collector on the Windows host whose channel you want to collect — a container reads its own Event Log channels, not the host's. To collect from another machine, use the receiver's [`remote`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/windowseventlogreceiver#remote-configuration) configuration instead of `channel`.

## Filtering by log level

`channel` collects a whole channel. To collect only selected levels, replace it with the commented-out [XML `query`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/windowseventlogreceiver#xml-queries) in `config.yaml`, which carries the channel in its `Path` attribute. `channel` and `query` are mutually exclusive. The `Level=` predicate is an exact match, so list every level you want to collect.
