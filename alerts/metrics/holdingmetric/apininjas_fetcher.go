package holdingmetric

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/prayansh1996/investment-alerts/currency"
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type ApiNinjasResponse struct {
	Ticker   string  `json:"ticker"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Exchange string  `json:"exchange"`
	Updated  int64   `json:"updated"`
}

type ApiNinjasHoldingMetricFetcher struct {
	holding holdings.Holding
}

func (f *ApiNinjasHoldingMetricFetcher) Fetch() (metrics.HoldingMetric, error) {
	resp, _ := f.getHttpResponse(f.holding.Api)
	return f.convertResponseToMetric(f.holding, resp)
}

func (f *ApiNinjasHoldingMetricFetcher) getHttpResponse(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-KEY", "cJWYYvQ0e9H/oqnLaBbMbQ==CvXKctKlnwIfjmCn")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (f *ApiNinjasHoldingMetricFetcher) convertResponseToMetric(rsu holdings.Holding, body []byte) (metrics.HoldingMetric, error) {
	var apiResponse ApiNinjasResponse
	err := json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling the response: %s\n", err)
		return metrics.HoldingMetric{}, err
	}

	nav := apiResponse.Price * currency.GetUsdToInrConversionRate()

	return metrics.HoldingMetric{
		Units:        rsu.UnitsHeld,
		PricePerUnit: nav,
		Name:         rsu.Name,
		Category:     rsu.Category,
	}, nil
}
