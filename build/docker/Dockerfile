FROM docker.io/library/golang:1.24.2-bookworm AS base

COPY /LICENSE /LICENSE
COPY ./ /src

FROM base AS builder
ARG BUILD_TAG="playground"
ARG BUILDER_VERSION="v0.123.0"
WORKDIR /src

RUN go install go.opentelemetry.io/collector/cmd/builder@${BUILDER_VERSION}
RUN CGO_ENABLED=0 GOEXPERIMENT=boringcrypto builder --config=/src/distributions/${BUILD_TAG}/manifest.yaml
# unify binary name as solarwinds-otel-collector
RUN mv /src/_build/solarwinds-otel-collector-${BUILD_TAG} /src/_build/solarwinds-otel-collector

FROM gcr.io/distroless/static-debian12@sha256:3d0f463de06b7ddff27684ec3bfd0b54a425149d0f8685308b1fdf297b0265e9

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=builder /src/_build/solarwinds-otel-collector /solarwinds-otel-collector

ENTRYPOINT ["/solarwinds-otel-collector"]
CMD ["--config=/opt/default-config.yaml"]
