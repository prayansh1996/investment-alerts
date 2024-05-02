package fetcher

import (
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type FixedDepositHoldingMetricFetcher struct {
}

func (f *FixedDepositHoldingMetricFetcher) Fetch(fd holdings.Holding) (metrics.HoldingMetric, error) {
	return metrics.HoldingMetric{
		Units:        fd.UnitsHeld,
		PricePerUnit: fd.StaticPricePerUnit,
		Name:         fd.Name,
		Category:     fd.Category,
	}, nil
}
