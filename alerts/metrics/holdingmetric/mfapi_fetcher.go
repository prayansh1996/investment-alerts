package holdingmetric

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/patrickmn/go-cache"
	"github.com/prayansh1996/investment-alerts/cons"
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type MfApiHoldingMetricFetcher struct {
	cache   *cache.Cache
	holding holdings.Holding
}

func NewMfApiFetcher(holding holdings.Holding) HoldingMetricFetcher {
	return &MfApiHoldingMetricFetcher{
		cache:   cache.New(cons.HOLDING_API_CACHE_DURATION, 2*cons.HOLDING_API_CACHE_DURATION),
		holding: holding,
	}
}

func (f *MfApiHoldingMetricFetcher) Fetch() (metrics.HoldingMetric, error) {
	var err error

	body, ok := f.cache.Get(f.holding.Api)
	if !ok {
		body, err = f.getHttpResponse(f.holding.Api)
		if err != nil {
			fmt.Printf("\nError reading response body for %s", f.holding.Api)
		}
		f.cache.Set(f.holding.Api, body, cons.HOLDING_API_CACHE_DURATION)
	}

	return f.convertResponseToMetric(f.holding, body.([]byte))
}

// Define the structs to match the JSON response
type MfApiResponse struct {
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

func (f *MfApiHoldingMetricFetcher) getHttpResponse(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (f *MfApiHoldingMetricFetcher) convertResponseToMetric(fund holdings.Holding, body []byte) (metrics.HoldingMetric, error) {
	var apiResponse MfApiResponse
	err := json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling the response: %s\n", err)
		return metrics.HoldingMetric{}, err
	}

	// Update the gauge with the fetched NAV
	if len(apiResponse.Data) == 0 {
		fmt.Printf("No data returned in response: %v\n", apiResponse)
		return metrics.HoldingMetric{}, err
	}

	nav, _ := strconv.ParseFloat(apiResponse.Data[0].Nav, 64)
	if nav == 0.0 {
		return metrics.HoldingMetric{}, errors.New("nav price is 0 for " + fund.Name)
	}

	return metrics.HoldingMetric{
		Units:        fund.UnitsHeld,
		PricePerUnit: nav,
		Name:         fund.Name,
		Category:     fund.Category,
	}, nil
}