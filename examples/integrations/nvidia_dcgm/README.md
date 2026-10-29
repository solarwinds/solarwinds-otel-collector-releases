# NVIDIA GPU (DCGM) Integration Example

This directory contains an example configuration for monitoring NVIDIA GPUs through the [NVIDIA DCGM Exporter](https://github.com/NVIDIA/dcgm-exporter). This configuration is compatible with SolarWinds Observability SaaS.

- `config.yaml`: Example configuration file for the NVIDIA GPU integration.

The collector scrapes the DCGM Exporter Prometheus endpoint and converts the `DCGM_FI_*` metrics to `nvidia.gpu.*` metrics with OpenTelemetry-style names, units and attributes.

## Prerequisites

#### 1. Install the NVIDIA driver and verify it works.

```sh
nvidia-smi
```

#### 2. Install the NVIDIA Container Toolkit and verify GPU access from a container.

```sh
sudo docker run --rm --gpus all ubuntu:24.04 nvidia-smi
```

#### 3. Run the DCGM Exporter. Bind it to localhost unless the collector runs on a different host.

```sh
sudo docker run -d --gpus all --restart unless-stopped -p 127.0.0.1:9400:9400 nvcr.io/nvidia/k8s/dcgm-exporter:<tag>
```

> **Warning**  
> The exporter has no authentication. Do not expose it on a public interface.

#### 4. Verify that the exporter returns metrics.

```sh
curl http://localhost:9400/metrics
```

## Configuration

Fill in the lines marked `# Required parameter` in `config.yaml`. In particular, replace the `${env:DCGM_EXPORTER_ENDPOINT}` target with the `host:port` of the DCGM Exporter (for example `localhost:9400`), or set the `DCGM_EXPORTER_ENDPOINT` environment variable for the collector.
