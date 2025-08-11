include submodules/solarwinds-otel-collector-core/build/Makefile.Common
include submodules/solarwinds-otel-collector-core/build/Makefile.Release
include submodules/solarwinds-otel-collector-core/build/Makefile.Licenses

# Define compatible otel_version with the current version of the collector
otel_version := 0.131.0

govulncheck-version = v1.1.4
goversioninfo-version = v1.5.0