# RabbitMQ Integration Example

This directory contains an example configuration for integrating RabbitMQ. This configuration is compatible with SolarWinds Observability Saas.

- `config.yaml`: Example configuration file for RabbitMQ integration.

## Prerequisities

Follow the steps below to configure RabbitMQ for monitoring.

> **Info**  
> The instructions below use the default configuration, with HTTP on port 15672, and Prometheus metrics on port 15692. To change the default protocol or port, refer to the RabbitMQ documentation.

#### 1. Make sure the RabbitMQ version running on your server is 3.8 or 3.9.

#### 2. Enable the RabbitMQ management plugin by running the command below.

```sh
rabbitmq-plugins enable rabbitmq_management
```

#### 3. In a browser, verify that RabbitMQ management is running on `http://localhost:15672`.

> **Tip**  
> Consider changing the default user and password to improve security.

#### 4. Enable the RabbitMQ Prometheus plugin. Refer to RabbitMQ documentation for details.

```sh
rabbitmq-plugins enable rabbitmq_prometheus
```

#### 5. Run the following command to test that Prometheus metrics are exposed. The call should return a list of metrics.

```sh
curl http://localhost:15692/metrics
```
