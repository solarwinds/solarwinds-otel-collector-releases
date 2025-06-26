# ZooKeeper Integration Example

This directory contains an example configuration for integrating ZooKeeper. This configuration is compatible with SolarWinds Observability Saas.

- `config.yaml`: Example configuration file for ZooKeeper integration.

## Prerequisities

Follow the steps below to configure ZooKeeper for monitoring.

> **Info**  
> The instructions below use the default configuration, on port 2181. To change the default port, refer to ZooKeeper documentation.

#### 1. Configure the mntr - '4 letter word command' in the ZooKeeper configuration file. By default, the file is located at `/opt/zookeeper/conf/zoo.cfg`.

Review the example below and adjust the configuration file accordingly.

```
4lw.commands.whitelist=mntr,ruok
```

> **Warning**  
> The ZooKeeper New Metric System feature is not supported in this integration. Instead, consider using the [Prometheus](../prometheus/README.md) integration.

#### 2. Run the command to get status

The output should have status data.

```sh
echo mntr | nc localhost 2181
```
