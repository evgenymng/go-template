# Go Service Template

This is a template for a Go service. It is intended to be used as a starting point for new services.

## This Branch — OpenTelemetry

This branch introduces an endpoint that sends metrics to an OpenTelemetry Collector. The endpoint is called `/send-trace`.

## Deploying

The easiest way to deploy the demo stand is to use Docker Compose. The `deploy` directory contains a `compose.yaml` file that can be used to start all the dependency services.

```shell
docker compose -f deploy/compose.yaml up -d
```

Or add `COMPOSE_FILE=deploy/compose.yaml` to your `.env` file and run `docker compose up -d`. The compose file consists of three services:
- `otel-collector` is an OpenTelemetry Collector that receives metrics (not logs or traces) and exports them to Prometheus format.
- `prometheus` is a Prometheus instance that aggregates metrics from the OpenTelemetry Collector.
- `grafana` is a Grafana instance that can be used to visualize metrics from the Prometheus.

The logs and traces are only sent to the OpenTelemetry Collector and not aggregated anywhere else. You can only observe them in the Collector logs. This is done to keep the demo as simple as possible.
