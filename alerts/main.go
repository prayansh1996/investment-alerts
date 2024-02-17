package main

import (
	"net/http"

	"github.com/prayansh1996/investment-alerts/fetchers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Fetch NAV periodically or on demand
	go fetchers.Start() // For simplicity, fetching once; consider using a ticker for periodic updates

	// Expose the registered metrics via HTTP
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8082", nil) // Use port 8082 for Prometheus metrics
}
