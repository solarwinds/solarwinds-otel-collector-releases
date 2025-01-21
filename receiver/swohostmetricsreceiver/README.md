# SWO Host Metrics Receiver

| Status                   |                      |
| ------------------------ | -------------------- |
| Stability                | [development]        |
| Supported pipeline types | metrics              |
| Distributions            | SolarWinds internal  |

SWO host metrics receiver generates metrics related to host. Receiver's purpose is to deliver functionality, which is not currently part of [opentelemetry-collector-contrib], or existing
implementation in contrib differs from our expectations.

## Receiver Configuration

Receiver supports collection interval configuration and generaly collection of scrapers.

```yaml
swohostmetrics:
  collection_interval: <duration> # default = 30s
  scrapers:
    <scraper1>:
    <scraper2>:
    ...
```

## Available Scrapers

| Scraper      | Supported OSs                | Description                                                 |
| ------------ | ---------------------------- | ----------------------------------------------------------- |
| [hostinfo]   | Linux & Windows              | Provides information about host entity itself               |
| [asset]      | Linux & Windows              | Provides information about installed software and features  |

[development]: https://github.com/open-telemetry/opentelemetry-collector#development
[opentelemetry-collector-contrib]: https://github.com/open-telemetry/opentelemetry-collector-contrib
[hostinfo]: ./internal/scraper/hostinfoscraper/documentation.md
[asset]: ./internal/scraper/assetscraper/documentation.md