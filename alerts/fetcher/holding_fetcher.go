package fetcher

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/prayansh1996/investment-alerts/cons"
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type HoldingFetcher interface {
	Fetch(holdings.Holding) (metrics.Metric, error)
}

func NewHoldingFetcher() HoldingFetcher {
	return &HoldingFetchOrchestrator{}
}

type HoldingFetchOrchestrator struct {
}

func (f *HoldingFetchOrchestrator) Fetch(holding holdings.Holding) (metrics.Metric, error) {
	url, err := url.Parse(holding.Api)
	if err != nil {
		fmt.Println("Error parsing url", err)
	}

	switch url.Hostname() {
	case cons.API_NINJAS_HOSTNAME:
		return (&ApiNinjasFetcher{}).Fetch(holding)

	case cons.MF_API_HOSTNAME:
		return (NewMfApiFetcher()).Fetch(holding)

	case cons.ZERODHA_KITE_HOSTNAME:
		return (NewZerodhaKiteFetcher()).Fetch(holding)
	}

	return metrics.Metric{}, errors.New("Unknown url encountered: " + holding.Api)
}
