package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define the structs to match the JSON response
type ApiResponse struct {
	Meta   MetaData  `json:"meta"`
	Data   []NavData `json:"data"`
	Status string    `json:"status"`
}

type MetaData struct {
	FundHouse      string `json:"fund_house"`
	SchemeType     string `json:"scheme_type"`
	SchemeCategory string `json:"scheme_category"`
	SchemeCode     int    `json:"scheme_code"`
	SchemeName     string `json:"scheme_name"`
}

type NavData struct {
	Date string `json:"date"`
	Nav  string `json:"nav"`
}

// Define a gauge outside of the main function to make it accessible in the handler
var navGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "mutual_fund_nav",
		Help: "NAV of the mutual fund.",
	},
)

func init() {
	// Register the gauge with Prometheus
	prometheus.MustRegister(navGauge)
}

func fetchNav() {
	url := "https://api.mfapi.in/mf/119063/latest"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the response body: %s\n", err)
		return
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling the response: %s\n", err)
		return
	}

	// Update the gauge with the fetched NAV
	if len(apiResponse.Data) > 0 {
		nav, err := strconv.ParseFloat(apiResponse.Data[0].Nav, 64)
		if err == nil {
			navGauge.Set(nav)
		}
	}
}

func tickNav() {
	for i := 0; i <= 100; i++ {
		fetchNav()
		time.Sleep(2 * time.Second)
		fmt.Println("Nav fetched")
	}
}

func main() {
	// Fetch NAV periodically or on demand
	go tickNav() // For simplicity, fetching once; consider using a ticker for periodic updates

	// Expose the registered metrics via HTTP
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8082", nil) // Use port 8082 for Prometheus metrics
}
