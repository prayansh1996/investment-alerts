package fetcher

import (
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type FDHoldingMetricFetcher struct {
}

func (f *FDHoldingMetricFetcher) Fetch(fd holdings.Holding) (metrics.HoldingMetric, error) {
	return metrics.HoldingMetric{
		Units:        fd.UnitsHeld,
		PricePerUnit: fd.StaticPricePerUnit,
		Name:         fd.Name,
		Category:     fd.Category,
	}, nil
}
