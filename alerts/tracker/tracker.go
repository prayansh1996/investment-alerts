package tracker

import (
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type Tracker interface {
	Track() metrics.Metric
}

func Start() {
	holdingTracker := NewHoldingTracker()

	holdingsList := []holdings.Holding{}
	holdingsList = append(holdingsList, holdings.GetMutualFundHoldings()...)
	holdingsList = append(holdingsList, holdings.GetRSUHoldings()...)
	holdingsList = append(holdingsList, holdings.GetFixedDepostHoldings()...)

	for _, holding := range holdingsList {
		go holdingTracker.getHoldingTracker(holding)(metrics.PublishChannel)
	}
}
