package main

import (
	"net/http"

	"github.com/prayansh1996/investment-alerts/metrics"
	"github.com/prayansh1996/investment-alerts/tracker"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	metrics.InitializeMetrics()
	go tracker.Start()

	// Expose the registered metrics via HTTP
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8082", nil) // Use port 8082 for Prometheus metrics
}
