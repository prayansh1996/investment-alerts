package holdingmetric

import (
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type FixedDepositHoldingMetricFetcher struct {
	holding holdings.Holding
}

func (f *FixedDepositHoldingMetricFetcher) Fetch() (metrics.HoldingMetric, error) {
	return metrics.HoldingMetric{
		Units:        f.holding.UnitsHeld,
		PricePerUnit: f.holding.StaticPricePerUnit,
		Name:         f.holding.Name,
		Category:     f.holding.Category,
	}, nil
}
