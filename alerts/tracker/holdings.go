package tracker

import (
	"fmt"
	"time"

	"github.com/prayansh1996/investment-alerts/fetcher"
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

type HoldingTracker struct {
	cachedFetcher fetcher.CachedFetcher
}

func NewHoldingTracker() HoldingTracker {
	f := HoldingTracker{}
	f.cachedFetcher = fetcher.NewCachedFetcher()
	return f
}

func (f *HoldingTracker) getHoldingTracker(holding holdings.Holding) func(chan<- metrics.Metric) {
	duration, err := time.ParseDuration(holding.RefreshTime)
	if err != nil {
		fmt.Printf("Cannot parse duration")
	}

	return func(publish chan<- metrics.Metric) {
		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		// Hack to make the first tick immediate
		for t := time.Now(); true; t = <-ticker.C {
			fmt.Printf("\nFetching %s %s at %s", holding.Name, holding.Category, t)

			holdingMetrics, err := f.cachedFetcher.Fetch(holding)
			if err != nil {
				fmt.Println(err)
				continue
			}

			publish <- holdingMetrics
		}
	}
}
