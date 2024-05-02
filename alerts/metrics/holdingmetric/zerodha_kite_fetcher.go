package holdingmetric

import (
	"bytes"
	"encoding/csv"
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

type ZerodhaKiteHoldingMetricFetcher struct {
	cache   *cache.Cache
	holding holdings.Holding
}

func NewZerodhaKiteFetcher(holding holdings.Holding) HoldingMetricFetcher {
	return &ZerodhaKiteHoldingMetricFetcher{
		cache:   cache.New(cons.HOLDING_API_CACHE_DURATION, 2*cons.HOLDING_API_CACHE_DURATION),
		holding: holding,
	}
}

func (f *ZerodhaKiteHoldingMetricFetcher) Fetch() (metrics.HoldingMetric, error) {
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

func (f *ZerodhaKiteHoldingMetricFetcher) getHttpResponse(url string) ([]byte, error) {
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

func (f *ZerodhaKiteHoldingMetricFetcher) convertResponseToMetric(fund holdings.Holding, body []byte) (metrics.HoldingMetric, error) {
	records, err := csv.NewReader(bytes.NewReader(body)).ReadAll()
	if err != nil {
		fmt.Printf("Unable to read records from csv: %s\n", err)
		return metrics.HoldingMetric{}, err
	}

	nav := 0.0
	for _, record := range records {
		if record[0] == fund.Symbol {
			nav, _ = strconv.ParseFloat(record[14], 64)
		}
	}
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
