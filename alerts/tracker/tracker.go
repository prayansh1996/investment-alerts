package tracker

import (
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type Tracker interface {
	Track() metrics.Metric
}

func Start() {
	fundTracker := NewHoldingTracker()

	funds := holdings.GetMutualFundHoldings()
	for _, fund := range funds {
		go fundTracker.getHoldingTracker(fund)(metrics.PublishChannel)
	}

	rsus := holdings.GetRSUHoldings()
	for _, rsu := range rsus {
		go fundTracker.getHoldingTracker(rsu)(metrics.PublishChannel)
	}
}
