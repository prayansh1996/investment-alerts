package fetcher

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

type MfApiFetcher struct {
	cache *cache.Cache
}

func NewMfApiFetcher() HoldingFetcher {
	return &MfApiFetcher{
		cache: cache.New(cons.HOLDING_API_CACHE_DURATION, 2*cons.HOLDING_API_CACHE_DURATION),
	}
}

func (f *MfApiFetcher) Fetch(holding holdings.Holding) (metrics.Metric, error) {
	var err error

	body, ok := f.cache.Get(holding.Api)
	if !ok {
		body, err = f.getHttpResponse(holding.Api)
		if err != nil {
			fmt.Printf("\nError reading response body for %s", holding.Api)
		}
		f.cache.Set(holding.Api, body, cons.HOLDING_API_CACHE_DURATION)
	}

	return f.convertResponseToMetric(holding, body.([]byte))
}

func (f *MfApiFetcher) getHttpResponse(url string) ([]byte, error) {
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

func (f *MfApiFetcher) convertResponseToMetric(fund holdings.Holding, body []byte) (metrics.Metric, error) {
	var apiResponse MfApiResponse
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
	if nav == 0.0 {
		return metrics.Metric{}, errors.New("nav price is 0 for " + fund.Name)
	}

	return metrics.Metric{
		Units:        fund.UnitsHeld,
		PricePerUnit: nav,
		Name:         fund.Name,
		Category:     fund.Category,
	}, nil
}
