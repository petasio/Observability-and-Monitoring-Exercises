package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	prom "github.com/nmdra/Observability-and-Monitoring-Exercises/Ex-1-Docker/app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
)

const (
	defaultAppPort     = "8080"
	defaultMetricsPort = "9091"
	shutdownTimeout    = 5 * time.Second
)

// --- Handlers ---
func helloHandler(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "world"
	}
	if sleep := c.Query("sleep_ms"); sleep != "" {
		if ms, err := strconv.Atoi(sleep); err == nil && ms > 0 && ms < 5000 {
			time.Sleep(time.Duration(ms) * time.Millisecond)
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("hello, %s\n", name))
}

func main() {
	// Register metrics once
	prom.PrometheusInit()

	appPort := getEnv("PORT", defaultAppPort)
	r := gin.Default()
	r.Use(prom.PrometheusMiddleware())
	r.GET("/hello", helloHandler)

	appServer := newServer(":"+appPort, r)

	// --- Metrics server ---
	metricsPort := getEnv("METRICS_PORT", defaultMetricsPort)
	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", promhttp.Handler())
	metricsServer := newServer(":"+metricsPort, metricsMux)

	// --- Run servers concurrently ---
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		log.Printf("Application server listening on :%s", appPort)
		if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("app server error: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		log.Printf("Metrics server listening on :%s", metricsPort)
		if err := metricsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("metrics server error: %w", err)
		}
		return nil
	})

	// --- Graceful shutdown ---
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
	case <-stop:
		log.Println("Shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := appServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("App server shutdown error: %v", err)
	}
	if err := metricsServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("Metrics server shutdown error: %v", err)
	}

	if err := g.Wait(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

	log.Println("Servers exited gracefully")
}

func newServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
