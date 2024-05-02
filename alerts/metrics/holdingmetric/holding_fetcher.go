package fetcher

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/prayansh1996/investment-alerts/cons"
	"github.com/prayansh1996/investment-alerts/holdings/fetcher"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type HoldingMetricFetcher interface {
	Fetch() (metrics.HoldingMetric, error)
}

func NewHoldingMetricFetcher(holdingFetcher fetcher.HoldingFetcher) HoldingMetricFetcher {
	if holding.StaticPricePerUnit > 0 {
		return (&FixedDepositHoldingMetricFetcher{}).Fetch(holding)
	}

	url, err := url.Parse(holding.Api)
	if err != nil {
		fmt.Println("Error parsing url", err)
	}

	switch url.Hostname() {
	case cons.API_NINJAS_HOSTNAME:
		return (&ApiNinjasHoldingMetricFetcher{}).Fetch(holding)

	case cons.MF_API_HOSTNAME:
		return (NewMfApiFetcher()).Fetch(holding)

	case cons.ZERODHA_KITE_HOSTNAME:
		return (NewZerodhaKiteFetcher()).Fetch(holding)
	}

	return metrics.HoldingMetric{}, errors.New("Unknown url encountered: " + holding.Api)
}
