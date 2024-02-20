package tracker

import (
	"fmt"
	"time"

	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/httpclient"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type FundTracker struct {
	cachedHttpClient httpclient.CachedHttpClient
}

func NewFundTracker() FundTracker {
	f := FundTracker{}
	f.cachedHttpClient = httpclient.NewCachedHttpClient()
	return f
}

func (f *FundTracker) getFundTracker(fund holdings.Fund) func(chan<- metrics.Metric) {
	duration, err := time.ParseDuration(fund.RefreshTime)
	if err != nil {
		fmt.Printf("Cannot parse duration")
	}

	return func(publish chan<- metrics.Metric) {
		ticker := time.NewTicker(duration)
		for {
			t := <-ticker.C
			fmt.Printf("\nFetching %s %s at %s", fund.Name, fund.Category, t)

			fundMetrics, _ := f.cachedHttpClient.Fetch(fund)
			publish <- fundMetrics
		}
	}
}
