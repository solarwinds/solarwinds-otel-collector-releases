# Copy the journactl binary from this Debian image with
# systemd (using the copy-journal-binary.sh script).
FROM debian:bookworm AS journal
RUN apt update
RUN apt install -y systemd
COPY /copy-journal-binary.sh /script.sh
RUN chmod +x /script.sh
RUN /script.sh

FROM gcr.io/distroless/static-debian12@sha256:3d0f463de06b7ddff27684ec3bfd0b54a425149d0f8685308b1fdf297b0265e9

COPY /LICENSE /LICENSE

ARG USER_UID=10001
USER ${USER_UID}

COPY --chmod=755 solarwinds-otel-collector /solarwinds-otel-collector
COPY --from=journal /journalctl-deps/ /

ENTRYPOINT ["/solarwinds-otel-collector"]
CMD ["--config=/opt/default-config.yaml"]
