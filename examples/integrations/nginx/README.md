# NGINX Integration Example

This directory contains an example configuration for integrating NGINX. This configuration is compatible with SolarWinds Observability Saas.

- `config.yaml`: Example configuration file for NGINX integration.

## Prerequisites

Follow the steps below to configure Nginx for monitoring.

> **Info**  
> The instructions below use the default configuration, with HTTP on port 80. To change the default protocol or port, refer to NGINX documentation.

#### 1. Make sure that the status module is already available.

Run the command below. If the command output does not include `with-http_stub_status_module`, enable the status module. Refer to NGINX documentation for details.

```sh
nginx -V 2>&1 | grep -o with-http_stub_status_module
```

#### 2. Configure the stub_status module in the NGINX configuration file. By default, the file is located at `/etc/nginx/sites-available/default`.

Review the example below and adjust the configuration file accordingly.

```nginx
location /status {
    stub_status;
    allow 127.0.0.1; #only allow requests from localhost
    deny all; #deny all other hosts
}
```

#### 3. Check the NGINX configuration file (`nginx -t`) for any errors. If there are no issues, run `nginx -s reload` to save and restart the NGINX service.

```sh
nginx -t
nginx -s reload
```

#### 4. Run the command to call the status URL. The call should return basic status data.

```sh
curl http://localhost:80/status
```
