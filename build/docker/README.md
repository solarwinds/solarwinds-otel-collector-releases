# SolarWinds OTel Collector Docker Image
Docker image for _SolarWinds OpenTelemetry Collector_:
 - contains full build of the Collector with full set of components registered
 - runs set of tests validating correct functionality of components used
 - starts the Collector at entrypoint, expecting input config located by default as `/opt/default-config.yaml`

## Getting the Image
Pull the image from DockerHub.

`docker pull solarwinds/solarwinds-otel-collector:0.113.2`

Optionally you can build the image yourself, simply run docker build command, i.e.

`docker build . -f build/docker/Dockerfile -t solarwinds-otel-collector:local`

## How to Run
To run the image utilize following command:

`docker run  -v ./your_config_file.yaml:/opt/default-config.yaml solarwinds-otel-collector:local`

> [!WARNING]  
> Note that the volume mounting **is required**. The image expects config at designated location at `/opt/default-config.yaml` by default to start properly.

## Layers
The image contains following layers:
- **base** - handles preparation of source code
- **builder** - builds the source code
- **tests** - runs the unit tests of the components defined in this repository used in the Collector

Final layer runs the binary built in **builder** layer, starting with `--config` parameter pointed at `/opt/default-config.yaml` location.
