ARG WIN_VERSION=2022
FROM mcr.microsoft.com/windows/nanoserver:ltsc${WIN_VERSION}

COPY /LICENSE /LICENSE

COPY solarwinds-otel-collector.exe ./solarwinds-otel-collector.exe

ENTRYPOINT ["/solarwinds-otel-collector"]
CMD ["--config=/opt/default-config.yaml"]