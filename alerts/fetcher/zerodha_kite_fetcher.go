package fetcher

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

type ZerodhaKiteFetcher struct {
	cache *cache.Cache
}

func NewZerodhaKiteFetcher() HoldingFetcher {
	return &ZerodhaKiteFetcher{
		cache: cache.New(cons.HOLDING_API_CACHE_DURATION, 2*cons.HOLDING_API_CACHE_DURATION),
	}
}

func (f *ZerodhaKiteFetcher) Fetch(holding holdings.Holding) (metrics.Metric, error) {
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

func (f *ZerodhaKiteFetcher) getHttpResponse(url string) ([]byte, error) {
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

func (f *ZerodhaKiteFetcher) convertResponseToMetric(fund holdings.Holding, body []byte) (metrics.Metric, error) {
	records, err := csv.NewReader(bytes.NewReader(body)).ReadAll()
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
