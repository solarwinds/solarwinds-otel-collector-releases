# swohostmetricsreceiver/hostinfo

## Scraper Configuration

Scraper metrics can be configured as follows.

```yaml
scrapers:
  hostinfo:
    metrics:
      <metric1>:
      <metric2>:
      ...
```

## Default Metrics

There are no metrics enabled by default. Every metric provided by this scraper needs to be enabled explicitly. For enabling metric use following pattern.

```yaml
metrics:
  <metric_name>:
    enabled: true
```

## swo.hostinfo.uptime

Host uptime in seconds.

### Metric Details

| Unit | Metric Type | Value Type | Cumulative | Monotonic |
| ---- | ----------- | ---------- | ---------- | --------- |
| s    | Sum         | Int        | True       | True      |

### Metric Attributes

- [osdetails attributes] - Operating system details information
- [hostdetails attributes] - Host details information

### Metric Configuration

```yaml
metrics:
  swo.hostinfo.uptime:
    enabled: true
```

## swo.hostinfo.user.lastLogged

Host last logged-in user. Supported for Windows and Linux.

### Metric Details

| Unit | Metric Type | Value Type |
|------|-------------|------------|
| User | Gauge       | Int        |

### Metric Attributes
#### Windows

| Name             | Description           | Values  | Example              |
| ---------------- | --------------------- | ------- | -------------------- |
| user.name        | user name with domain | Any Str | SHORDDOMAIN\john.doe |
| user.displayname | user display name     | Any Str | John Doe             |

#### Linux

| Name             | Description           | Values  | Example |
| ---------------- | --------------------- | ------- |---------|
| user.name        | user name             | Any Str | ubuntu  |


### Metric Configuration

```yaml
metrics:
  swo.hostinfo.user.lastLogged:
    enabled: true
```

## swo.hostinfo.firewall

Metric provides firewall profiles statuses. This metric is supported only on Windows. On Linux or other platform provides empty metric slice.

### Metric Details

| Unit        | Metric Type | Value Type | Cumulative | Monotonic |
| ----------- | ----------- | ---------- | ---------- | --------- |
| {status}    | Gauge       | Int        | False      | False     |

- `status` values can be 0 or 1 where:
  - value 1 means a firewall profile is `enabled`
  - value 0 means a firewall profile is `disabled`

### Metric Attributes

| Name                  | Description             | Values  | Example              |
| --------------------- | ----------------------- | ------- | -------------------- |
| firewall.profile.name | Firewall's profile name | Any Str | domain               |

### Metric Configuration

```yaml
metrics:
  swo.hostinfo.firewall:
    enabled: true
```

[osdetails attributes]: ../../attributes/osdetails/documentation.md
[hostdetails attributes]: ../../attributes/hostdetails/documentation.md
