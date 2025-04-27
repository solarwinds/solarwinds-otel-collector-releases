# Contributing

Thank you for your interest in contributing to SolarWinds OpenTelemetry Collector!

Currently, this project is being developed by the SolarWinds team. 
We are actively working on stabilizing its foundations and defining contribution guidelines.
Once ready, we will update the `CONTRIBUTING.md` file with clear instructions.

## Before Adding a New Component

Before proposing or implementing a new OpenTelemetry component, please consider the following:

### Evaluate Necessity
- **Is a new component truly needed?** Many telemetry transformation requirements can be solved using existing components, especially with processors like `transform`, `filter`, or `metricstransform`.
- Review the [current components list](README.md#components) and the [OpenTelemetry Collector Contrib repository](https://github.com/open-telemetry/opentelemetry-collector-contrib) to avoid duplication.

### Contact Us First
- For SolarWinds internal contributors: Reach out in the `#swotelcol-platform` Slack channel to discuss your proposal.
- For external contributors: Open an issue describing your use case and proposed component before investing significant time in development.

Before any code is written contact us providing the following information:
* Who's the sponsor for your component. A sponsor is an approver or maintainer who will be the official reviewer of the code and a code owner for the component.
* Some information about your component, such as the reasoning behind it, use-cases, telemetry data types supported, and anything else you think is relevant for us.
* The configuration options your component will accept. This will give us a better understanding of what it does, and how it may be implemented.

## Development

For general information on the development process in this project,
please refer to the [Development Guidelines](docs/development-guidelines.md).

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
