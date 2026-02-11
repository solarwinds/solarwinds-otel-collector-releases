# MongoDB Integration Example

This directory contains an example configuration for integrating MongoDB. 
This configuration is compatible with SolarWinds Observability SaaS.

- `config.yaml`: Example configuration file for MongoDB integration.

## Prerequisites

Following MongoDB versions are supported:

4.0+
5.0
6.0
7.0

It is recommended to set up a least privilege user (LPU) with a clusterMonitor role in order to collect metrics.
Example:
```
  use admin
  db.createUser(
          {
              user: "${USER}",
              pwd: "${PASS}",
              roles: [ "clusterMonitor"]
          }
  );
```