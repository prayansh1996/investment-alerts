package tracker

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
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

func getFundTracker(fund holdings.Fund) func(chan<- metrics.Metric) {
	duration, err := time.ParseDuration(fund.RefreshTime)
	if err != nil {
		fmt.Printf("Cannot parse duration")
	}

	return func(publish chan<- metrics.Metric) {
		ticker := time.NewTicker(duration)
		for {
			t := <-ticker.C
			fmt.Printf("\nFetching %s %s at %s", fund.Name, fund.Category, t)

			fundMetrics, _ := getFundMetrics(fund)
			publish <- fundMetrics
		}
	}
}

func getFundMetrics(fund holdings.Fund) (metrics.Metric, error) {
	resp, err := http.Get(fund.Api)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return metrics.Metric{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the response body: %s\n", err)
		return metrics.Metric{}, err
	}

	r := csv.NewReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("Error reading csv: %s\n", err)
		return metrics.Metric{}, err
	}

	// Update the gauge with the fetched NAV
	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Unable to read records from csv: %s\n", err)
		return metrics.Metric{}, err
	}

	nav := 0.0
	for _, record := range records {
		if record[0] == fund.Symbol {
			nav, _ = strconv.ParseFloat(record[14], 64)
		}
	}

	return metrics.Metric{
		Units:        fund.UnitsHeld,
		PricePerUnit: nav,
		Name:         fund.Name,
		Category:     fund.Category,
	}, nil
}
