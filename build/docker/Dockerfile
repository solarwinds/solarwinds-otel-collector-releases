FROM docker.io/library/golang:1.23.4-bookworm AS base


# Stage 1: Dependency caching
FROM base AS deps
    # Set the working directory
WORKDIR /app
# Copy go.mod files to cache dependencies
COPY /cmd/solarwinds-otel-collector/go.mod /src/cmd/solarwinds-otel-collector/go.mod
COPY /extension/solarwindsextension/go.mod /src/extension/solarwindsextension/go.mod
COPY /exporter/solarwindsexporter/go.mod /src/exporter/solarwindsexporter/go.mod
# Download dependencies
RUN cd /src/cmd/solarwinds-otel-collector && go mod download

FROM deps AS builder

# Copy the dependency cache
COPY --from=deps /go/pkg /go/pkg
COPY --from=deps /go/bin /go/bin

COPY /LICENSE /LICENSE
COPY ./ /src

RUN cd /src/cmd/solarwinds-otel-collector && CGO_ENABLED=0 GOEXPERIMENT=boringcrypto go build -trimpath -o /src/bin/solarwinds-otel-collector "-ldflags=-s -w"

FROM builder AS tests
WORKDIR /src
# run tests for go modules of all maintained components
# image build is stopped if test failure is detected
RUN find . -name go.mod -not -path "./cmd/solarwinds-otel-collector/*" -execdir go test ./... \; | ( ! grep FAIL )

FROM gcr.io/distroless/static-debian12@sha256:41972110a1c1a5c0b6adb283e8aa092c43c31f7c5d79b8656fbffff2c3e61f05

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=tests /src/bin/solarwinds-otel-collector /solarwinds-otel-collector

ENTRYPOINT ["/solarwinds-otel-collector"]
CMD ["--config=/opt/default-config.yaml"]
