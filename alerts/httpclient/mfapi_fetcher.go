package httpclient

import (
	"encoding/json"
	"fmt"
	"strconv"

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

func mfApiFetcher(fund holdings.Fund, body []byte) (metrics.Metric, error) {
	var apiResponse ApiResponse
	err := json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling the response: %s\n", err)
		return metrics.Metric{}, err
	}

	// Update the gauge with the fetched NAV
	if len(apiResponse.Data) == 0 {
		fmt.Printf("No data returned in response: %v\n", apiResponse)
		return metrics.Metric{}, err
	}

	nav, _ := strconv.ParseFloat(apiResponse.Data[0].Nav, 64)
	return metrics.Metric{
		Units:        fund.UnitsHeld,
		PricePerUnit: nav,
		Name:         fund.Name,
		Category:     fund.Category,
	}, nil
}
