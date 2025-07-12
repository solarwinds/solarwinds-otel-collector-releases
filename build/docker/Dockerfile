FROM docker.io/library/golang:1.24.5-bookworm AS base

COPY /LICENSE /LICENSE
COPY ./ /src

FROM base AS builder
ARG BUILD_TAG="playground"
WORKDIR /src

RUN make build distribution=${BUILD_TAG}
# unify binary name as solarwinds-otel-collector
RUN mv /src/_build/solarwinds-otel-collector-${BUILD_TAG} /src/_build/solarwinds-otel-collector

FROM gcr.io/distroless/static-debian12@sha256:b7b9a6953e7bed6baaf37329331051d7bdc1b99c885f6dbeb72d75b1baad54f9

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=builder /src/_build/solarwinds-otel-collector /solarwinds-otel-collector

ENTRYPOINT ["/solarwinds-otel-collector"]
CMD ["--config=/opt/default-config.yaml"]
