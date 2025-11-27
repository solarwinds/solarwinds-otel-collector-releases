# Override TOOLS_MOD_DIR to use tools from submodule
TOOLS_MOD_DIR := $(CURDIR)/submodules/solarwinds-otel-collector-core/internal/tools

include submodules/solarwinds-otel-collector-core/build/Makefile.Common
include submodules/solarwinds-otel-collector-core/build/Makefile.Release
include submodules/solarwinds-otel-collector-core/build/Makefile.Licenses

# Define compatible otel_version with the current version of the collector
otel_version := 0.140.0