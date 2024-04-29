package fetcher

import (
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type FDFetcher struct {
}

func (f *FDFetcher) Fetch(fd holdings.Holding) (metrics.Metric, error) {
	return metrics.Metric{
		Units:        fd.UnitsHeld,
		PricePerUnit: fd.StaticPricePerUnit,
		Name:         fd.Name,
		Category:     fd.Category,
	}, nil
}
