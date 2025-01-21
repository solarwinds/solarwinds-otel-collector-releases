# swohostmetricsreceiver/hardwareinventory

## Scraper Configuration

Scraper metrics can be configured as follows.

```yaml
scrapers:
  hardwareinventory:
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

## swo.hardwareinventory.cpu

CPU current clock speed in MHz.

### Metric Details

| Unit | Metric Type | Value Type |
|------|-------------| ---------- |
| MHz  | Gauge       | Int        |

### Metric Attributes


| Name                   | Description                                                                                              | Values  | Example                                   |
|------------------------|----------------------------------------------------------------------------------------------------------|---------|-------------------------------------------|
| processor.name         | Processor Name                                                                                           | Any Str | Intel(R) Xeon(R) CPU E5-2686 v4 @ 2.30GHz |
| processor.caption      | Short description of the processor                                                                       | Any Str | Intel64 Family 6 Model 79 Stepping 1      |
| processor.manufacturer | Processor Manufacturer                                                                                   | Any Str | GenuineIntel                              |
| processor.model        | Processor Model                                                                                          | Any Str | 79                                        |
| processor.stepping     | Revision level of the processor in the processor family                                                  | Any Str | 1                                         |
| processor.cores        | Number of physical cores                                                                                 | Any Str | 6                                         |
| processor.threads      | Number of logical units (with hyper-threading enabled, the processor will have more threads then cores)  | Any Str | 12                                        |

### Metric Configuration

```yaml
metrics:
  swo.hardwareinventory.cpu:
    enabled: true
```
