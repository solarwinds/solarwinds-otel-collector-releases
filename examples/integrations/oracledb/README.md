# OracleDB Integration Example

This directory contains an example configuration for integrating OracleDB. 
This configuration is compatible with SolarWinds Observability SaaS.

- `config.yaml`: Example configuration file for OracleDB integration.

## Prerequisites

Follow the steps below to configure OracleDB for monitoring.

To collect metrics from Oracle DB, the user account used for data collection needs access to specific Oracle DB tables.  

Specify the username, copy and run the code to grant the user the access.

> **Info**  
> If the user cannot access the tables, metrics related to the tables will not be collected.

```sql
GRANT SELECT ON V_$SESSION TO <username>;
GRANT SELECT ON V_$SYSSTAT TO <username>;
GRANT SELECT ON V_$RESOURCE_LIMIT TO <username>;
GRANT SELECT ON DBA_TABLESPACES TO <username>;
GRANT SELECT ON DBA_DATA_FILES TO <username>;
```
