package fetcher

import (
	"github.com/prayansh1996/investment-alerts/holdings/fetcher"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type FixedDepositHoldingMetricFetcher struct {
	holdingFetcher fetcher.HoldingFetcher
}

func (f *FixedDepositHoldingMetricFetcher) Fetch() (metrics.HoldingMetric, error) {
	holding := f.holdingFetcher.Fetch()
	return metrics.HoldingMetric{
		Units:        holding.UnitsHeld,
		PricePerUnit: holding.StaticPricePerUnit,
		Name:         holding.Name,
		Category:     holding.Category,
	}, nil
}
