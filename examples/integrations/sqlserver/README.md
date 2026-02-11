# SqlServer Integration Example

This directory contains an example configuration for integrating SqlServer. 
This configuration is compatible with SolarWinds Observability SaaS.

- `config.yaml`: Example configuration file for SqlServer integration.

## Prerequisites

Follow the steps below to configure SqlServer for monitoring.

The user must have the following permissions:

At least one of the following permissions:
```
CREATE DATABASE
ALTER ANY DATABASE
VIEW ANY DATABASE
```
Permission to view server state:
```
SQL Server pre-2022: VIEW SERVER STATE
SQL Server 2022 and later: VIEW SERVER PERFORMANCE STATET SELECT ON DBA_DATA_FILES TO <username>;
```
