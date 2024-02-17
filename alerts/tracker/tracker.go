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
	for _, fund := range funds {
		go getFundTracker(fund)(metrics.PublishChannel)
	}
}
