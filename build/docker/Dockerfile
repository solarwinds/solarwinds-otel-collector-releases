FROM docker.io/library/golang:1.24.4-bookworm AS base

COPY /LICENSE /LICENSE
COPY ./ /src

FROM base AS builder
ARG BUILD_TAG="playground"
WORKDIR /src

RUN make build distribution=${BUILD_TAG}
# unify binary name as solarwinds-otel-collector
RUN mv /src/_build/solarwinds-otel-collector-${BUILD_TAG} /src/_build/solarwinds-otel-collector

FROM gcr.io/distroless/static-debian12@sha256:d9f9472a8f4541368192d714a995eb1a99bab1f7071fc8bde261d7eda3b667d8

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=builder /src/_build/solarwinds-otel-collector /solarwinds-otel-collector

ENTRYPOINT ["/solarwinds-otel-collector"]
CMD ["--config=/opt/default-config.yaml"]
