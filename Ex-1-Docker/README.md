
This is a Go API built with **Gin** exposing HTTP endpoints and **Prometheus metrics** for monitoring:

* Tracks total HTTP requests (`http_requests_total`), request duration (`http_request_duration_seconds`), in-flight requests (`in_flight_requests`).
* Metrics available at `/metrics`, scrapeable by Prometheus.
* Includes a Grafana dashboard for request counts, latency, in-flight requests, and system performance (CPU, memory, disk, network via node\_exporter).

### Usage

1. Run the app:

```bash
go run main.go
```

2. Access metrics:

```
http://localhost:9090/metrics
```

3. Load test with `wrk` (example: 10 connections, 30s duration):

```bash
wrk -t10 -c10 -d30s http://localhost:8080/hello
```

4. Import the provided Grafana JSON dashboard for real-time visualization.

--- 

![Dashboord](./Screenshot%202025-08-19%20at%2014-59-01%20Go%20API%20System%20Metrics%20Dashboard%20-%20Dashboards%20-%20Grafana.png)

![Drilldown](./Screenshot%202025-08-19%20at%2015-02-30%20-%20Metrics%20-%20Drilldown%20-%20Grafana.png)
