# Contributing

Thank you for your interest in contributing to SolarWinds OpenTelemetry Collector!

Currently, this project is being developed by the SolarWinds team. 
We are actively working on stabilizing its foundations and defining contribution guidelines.
Once ready, we will update the `CONTRIBUTING.md` file with clear instructions.

## Development

## CI/CD
The _SolarWinds OpenTelemetry Collector_ utilizes [GitHub Actions pipeline](./.github).
The standard build pipeline is triggered with each PR opened to main or release branch and consists of:
- build of _SolarWinds OpenTelemetry Collector_ docker image

The release pipeline is triggered with designated tag publishing and consists of:
- build of _SolarWinds OpenTelemetry Collector_ docker image and its publishing to DockerHub
- creation of GitHub release

## Distribution Build Tags
There are two distributions defined based on Go build tags differing in the number of components:
- `full` - This tag represents the default distribution with all the components described in the [README](README.md#components).
- `k8s` - This is a special distribution based on components used in 
        [solarwinds/swi-k8s-opentelemetry-collector](https://github.com/solarwinds/swi-k8s-opentelemetry-collector). 
        It is not intended for direct use by users. There is a special Docker image for this distribution described in
        [build/docker/README.md#the-k8s-dockerfile](build/docker/README.md#the-k8s-dockerfile).
