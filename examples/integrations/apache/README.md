# Apache Integration Example

This directory contains an example configuration for integrating Apache. This configuration is compatible with SolarWinds Observability SaaS.

- `config.yaml`: Example configuration file for Apache integration.

## Prerequisites

Follow the steps below to configure Apache for monitoring.

> **Info**  
> The instructions below use the default configuration, with HTTP on port 80. To change the default protocol or port, refer to Apache documentation.

#### 1. Make sure that the `mod_status` module is already available.

Run the command below. If the command output does not include status_module, enable the mod_status module. Refer to Apache documentation for details.

```sh
apachectl -t -D DUMP_MODULES | grep status_module
```

#### 2. Configure the `mod_status` module in the Apache configuration file (`status.conf`). By default, the file is located at `/etc/apache2/mods-enabled/`

```
<Location /server-status>
	SetHandler server-status
	Require local
</Location>

ExtendedStatus On
```

#### 3. Check the Apache configuration file `(apachectl -t)` for any errors. If there are no issues, run apachectl restart to save and restart the Apache service.

```sh
apachectl -t
apachectl restart
```

#### 4. Run the command to call the status URL. The call should return basic status data.


```sh
curl http://localhost:80/server-status
```
