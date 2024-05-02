package holdingmetric

import (
	"fmt"
	"net/url"

	"github.com/prayansh1996/investment-alerts/cons"
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type HoldingMetricFetcher interface {
	Fetch() (metrics.HoldingMetric, error)
}

func NewHoldingMetricFetcher(holding holdings.Holding) HoldingMetricFetcher {
	if holding.StaticPricePerUnit > 0 {
		return &FixedDepositHoldingMetricFetcher{holding}
	}

	url, err := url.Parse(holding.Api)
	if err != nil {
		fmt.Println("Error parsing url", err)
	}

	switch url.Hostname() {
	case cons.API_NINJAS_HOSTNAME:
		return &ApiNinjasHoldingMetricFetcher{holding}

	case cons.MF_API_HOSTNAME:
		return NewMfApiFetcher(holding)

	case cons.ZERODHA_KITE_HOSTNAME:
		return NewZerodhaKiteFetcher(holding)
	}

	panic("Invalid Holding " + holding.Name)
}
