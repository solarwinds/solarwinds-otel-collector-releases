# OpenTelemetry Changes Analyzer

This tool analyzes changes in the [OpenTelemetry Collector Contrib](https://github.com/open-telemetry/opentelemetry-collector-contrib) repository between two specified versions. It generates a Markdown summary of breaking changes, deprecations, and enhancements for components used in your project, perfect for GitHub comments or documentation.

## Purpose

- **Compare Versions**: Identify changes between two version tags (e.g., `v0.119.0` to `v0.121.0`).
- **Component-Specific**: Focuses on components extracted from a `go.mod` file or provided manually.
- **Readable Output**: Formats results in Markdown with clear sections and a compare URL.


## Running the Tool
Run the tool with the following command, adjusting paths and versions as needed:
```
go run ./main.go 
  --old v0.119.0 
  --new v0.121.0 
  --goModPath ./../../../cmd/solarwinds-otel-collector/go.mod 
  --dependencyFilter opentelemetry-collector-contrib 
  --repo opentelemetry-collector-contrib
```

--old: Starting version (e.g., v0.119.0).
--new: Ending version (e.g., v0.121.0).
--goModPath: Path to your go.mod file to detect components.
--dependencyFilter: Filters components from go.mod (e.g., opentelemetry-collector-contrib).
--encode: Flag to base64 encode the output.
--repo: OpenTelemetry repository name, as used in URL.

# Example Output

The tool produces a Markdown summary message like this:

---
# OPENTELEMETRY-COLLECTOR-CONTRIB CHANGES
**Compare URL**: [0.119.0 to 0.121.0](https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.119.0...v0.121.0)

#### cumulativetodeltaprocessor
- **Enhancements**:
  - v0.119.0: cumulativetodeltaprocessor: Add metric type filter for cumulativetodelta processor ([#33673](https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/33673))

---
#### prometheusreceiver
- **Deprecations**:
  - v0.121.0: prometheusreceiver: Deprecate metric start time adjustment in the prometheus receiver. It is being replaced by the metricstarttime processor. ([#37186](https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/37186))
- **Enhancements**:
  - v0.119.0: prometheusreceiver: Add `receiver.prometheusreceiver.UseCollectorStartTimeFallback` featuregate for the start time metric adjuster to use the collector start time as an approximation of process start time as a fallback. ([#36364](https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/36364))
