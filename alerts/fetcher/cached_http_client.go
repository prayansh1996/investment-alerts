package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type DomainFetcherMap map[string]func(holdings.Holding, []byte) (metrics.Metric, error)

const CACHE_DURATION = 5 * time.Minute

type CachedFetcher struct {
	domainFetcherMap DomainFetcherMap
	cache            *cache.Cache
}

func NewCachedFetcher() CachedFetcher {
	c := CachedFetcher{}
	c.domainFetcherMap = DomainFetcherMap{
		"api.mfapi.in":       mfApiFetcher,
		"api.kite.trade":     zerodhaKiteFetcher,
		"api.api-ninjas.com": apiNinjasFetcher,
	}
	c.cache = cache.New(CACHE_DURATION, 2*CACHE_DURATION)
	return c
}

func (c *CachedFetcher) Fetch(holding holdings.Holding) (metrics.Metric, error) {
	var err error

	body, ok := c.cache.Get(holding.Api)
	if !ok {
		body, err = getHttpResponseAsBytes(holding.Api)
		if err != nil {
			fmt.Printf("\nError reading response body for %s", holding.Api)
		}
		c.cache.Set(holding.Api, body, CACHE_DURATION)
	}

	url, err := url.Parse(holding.Api)
	if err != nil {
		fmt.Println("Error parsing url", err)
	}

	fetcher := c.domainFetcherMap[url.Hostname()]
	return fetcher(holding, body.([]byte))
}

func getHttpResponseAsBytes(api string) ([]byte, error) {
	resp, err := http.Get(api)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
