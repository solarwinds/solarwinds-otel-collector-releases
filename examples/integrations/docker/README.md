# Docker Integration Example

This directory contains an example configuration for integrating Docker. This configuration is compatible with SolarWinds Observability SaaS.

- `config.yaml`: Example configuration file for Docker integration.

## Prerequisites

Follow the steps below to configure Docker for monitoring.

#### 1. Grant the agent access to the Docker socket file

The agent (`swagent` user) must have access to the Docker socket file.

```sh
usermod -a -G docker swagent
```

#### 2. Restart the agent

After the agent is granted access, it needs to be restarted.

```sh
service uamsclient restart
```
