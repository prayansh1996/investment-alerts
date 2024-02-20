package httpclient

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

type DomainFetcherMap map[string]func(holdings.Fund, []byte) (metrics.Metric, error)

const CACHE_DURATION = 5 * time.Minute

type CachedHttpClient struct {
	domainFetcherMap DomainFetcherMap
	cache            *cache.Cache
}

func NewCachedHttpClient() CachedHttpClient {
	c := CachedHttpClient{}
	c.domainFetcherMap = DomainFetcherMap{
		"api.mfapi.in":   mfApiFetcher,
		"api.kite.trade": zerodhaMfFetcher,
	}
	c.cache = cache.New(CACHE_DURATION, 2*CACHE_DURATION)
	return c
}

func (c *CachedHttpClient) Fetch(fund holdings.Fund) (metrics.Metric, error) {
	var err error

	body, ok := c.cache.Get(fund.Api)
	if !ok {
		body, err = getHttpResponseAsBytes(fund.Api)
		if err != nil {
			fmt.Printf("\nError reading response body for %s", fund.Api)
		}
		c.cache.Set(fund.Api, body, CACHE_DURATION)
	}

	url, err := url.Parse(fund.Api)
	if err != nil {
		fmt.Println("Error parsing url", err)
	}

	fetcher := c.domainFetcherMap[url.Hostname()]
	return fetcher(fund, body.([]byte))
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
