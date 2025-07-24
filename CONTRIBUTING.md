# Contributing

Thank you for your interest in contributing to SolarWinds OpenTelemetry Collector!

Currently, this project is being developed by the SolarWinds team. 
We are actively working on stabilizing its foundations and defining contribution guidelines.
Once ready, we will update the `CONTRIBUTING.md` file with clear instructions.

## Development

For general information on the development process in this project,
please refer to the [Development Guidelines](docs/development-guidelines.md).

## CI/CD
The release pipeline is triggered with designated tag publishing and consists of:
The _SolarWinds OpenTelemetry Collector_ utilizes [GitHub Actions pipeline](./.github).
The build pipeline is triggered with each PR opened to main or release branch and performs basic static checks.

The release pipeline is triggered with changes to [version file](./pkg/version.go) and consists of:
- verifying the existence of _SolarWinds OpenTelemetry Collector_ Docker images in DockerHub repository
- creating the GitHub release with DockerHub links