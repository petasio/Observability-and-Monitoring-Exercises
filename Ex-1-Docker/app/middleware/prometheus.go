package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics
var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "code"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of latencies for HTTP request",
			Buckets: []float64{0.01, 0.05, 0.1, 0.3, 0.5, 1, 2, 5}, // custom buckets
		},
		[]string{"path", "method", "code"},
	)

	inFlightRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "in_flight_requests",
			Help: "Current number of in-flight requests",
		},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(requestCounter, requestDuration, inFlightRequests)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		inFlightRequests.Inc()
		defer inFlightRequests.Dec()

		start := time.Now()
		c.Next()

		statusCode := c.Writer.Status()
		duration := time.Since(start).Seconds()

		// handle empty path
		path := c.FullPath()
		if path == "" {
			path = "unknown"
		}

		labels := prometheus.Labels{
			"path":   path,
			"method": c.Request.Method,
			"code":   strconv.Itoa(statusCode),
		}

		requestCounter.With(labels).Inc()
		requestDuration.With(labels).Observe(duration)
	}
}
