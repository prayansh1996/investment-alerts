package fetcher

import (
	"encoding/json"
	"fmt"

	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type NinjasApiResponse struct {
	Ticker   string  `json:"ticker"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Exchange string  `json:"exchange"`
	Updated  int64   `json:"updated"`
}

func apiNinjasFetcher(rsu holdings.Holding, body []byte) (metrics.Metric, error) {
	var apiResponse NinjasApiResponse
	err := json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling the response: %s\n", err)
		return metrics.Metric{}, err
	}

	nav := apiResponse.Price * 83.40

	return metrics.Metric{
		Units:        rsu.UnitsHeld,
		PricePerUnit: nav,
		Name:         rsu.Name,
		Category:     rsu.Category,
	}, nil
}
