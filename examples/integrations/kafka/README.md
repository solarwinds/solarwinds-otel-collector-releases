# Kafka Integration Example

This directory contains an example configuration for integrating Kafka. This configuration is compatible with SolarWinds Observability Saas.

- `config.yaml`: Example configuration file for Kafka integration.

## Prerequisities

Follow the steps below to configure Kafka for monitoring.

> **Info**  
> The instructions below use the default configuration, with connection to the local JVM without authentication.  
On the machine where Kafka is running:

#### 1. Download the Prometheus JMX Exporter from the [official github repository](https://github.com/prometheus/jmx_exporter).

#### 2. Download the Prometheus JMX Exporter configuration from our site 

Edit the configuration if the JMX connection is protected and/or should be done over SSL (see the [JMX Exporter documentation](https://github.com/prometheus/jmx_exporter#configuration) for more details)

```yaml
username: <username>
password: <password>
ssl: true

lowercaseOutputName: true
rules:
...
```

#### 3. Set the EXTRA_ARGS environment variable to configure the JMX Exporter

The example below uses kafka endpoint `localhost:1234`, which might need to be adjusted accordingly. The paths can also differ.


```sh
EXTRA_ARGS=-javaagent:/jmx-exporter/jmx_prometheus_javaagent-0.18.0.jar=localhost:1234:/jmx-exporter/jmx-kafka-config.yml
```

The JMX Exporter can be configured to bind to all interfaces. This is not recommended, but can be achieved by specifying just the port number in the environment variable.

Example:

```sh
EXTRA_ARGS=-javaagent:/jmx-exporter/jmx_prometheus_javaagent-0.18.0.jar=1234:/jmx-exporter/jmx-kafka-config.yml
```

#### 4. Restart the Kafka server.

#### 5. Run the command to test that metrics are exposed. The call should return list of metrics.

```sh
curl http://localhost:1234
```
