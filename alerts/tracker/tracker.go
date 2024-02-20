package tracker

import (
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type Tracker interface {
	Track() metrics.Metric
}

func Start() {
	funds := holdings.GetHoldings().Equity.MutualFunds
	fundTracker := NewFundTracker()
	for _, fund := range funds {
		go fundTracker.getFundTracker(fund)(metrics.PublishChannel)
	}
}
