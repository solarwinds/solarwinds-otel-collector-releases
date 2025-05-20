# SolarWinds OTel Collector Docker Images

Docker images for _SolarWinds OpenTelemetry Collector_.

There are several available images, differentiated by tags. They differ in the scope of available OpenTelemetry components.

`playground` 
- contains the full set of components provided by OpenTelemetry community, and component created for use in __SolarWinds Observability SaaS__.
- list of included components can be found in [components_playground](/cmd/solarwinds-otel-collector/components_playground.go)
`verified`   
- contains limited set of components used by __SolarWinds Observability SaaS__, including components that are part of **k8s**.
- list of included components can be found in  [components_verified](/cmd/solarwinds-otel-collector/components_verified.go)
`k8s`
- contains limited set of components __SolarWinds Observability SaaS__ for purposes of kubernetes monitoring feature.
- list of included components can be found in [components_k8s](/cmd/solarwinds-otel-collector/components_k8s.go)

There are also Windows images available. For details see [dockerhub](https://hub.docker.com/r/solarwinds/solarwinds-otel-collector/tags).

## Getting the Image
Pull the image from DockerHub. You can use other available tags `k8s`, `verified`, `playground` and their windows variants.

`docker pull solarwinds/solarwinds-otel-collector:latest`

Optionally you can build the image yourself, simply run docker build command, i.e.

`docker build . -f build/docker/Dockerfile --build-arg BUILD_TAG=playground --build-arg BUILDER_VERSION="v0.123.0" -t solarwinds-otel-collector:local`

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

## The "k8s" Dockerfile

The [Dockerfile.k8s](Dockerfile.k8s) builds a distribution of the `solarwinds-otel-collector` defined with the `k8s` 
Go build tag (see [CONTRIBUTING.md](../../CONTRIBUTING.md#distribution-build-tags)).
