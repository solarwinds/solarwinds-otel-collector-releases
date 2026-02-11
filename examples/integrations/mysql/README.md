# MySql Integration Example

This directory contains an example configuration for integrating MySql. 
This configuration is compatible with SolarWinds Observability SaaS.

- `config.yaml`: Example configuration file for MySql integration.

## Prerequisites

Collecting most metrics requires the ability to execute SHOW GLOBAL STATUS.

Collecting query samples requires the performance_schema to be enabled:

```
GRANT SELECT ON performance_schema.* TO <your-user>@'%';
```
