package tracker

import (
	"fmt"
	"time"

	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
	fetcher "github.com/prayansh1996/investment-alerts/metrics/holdingmetricfetcher"
)

type HoldingTracker struct {
	fetcher fetcher.HoldingMetricFetcher
}

func NewHoldingTracker() HoldingTracker {
	tracker := HoldingTracker{}
	tracker.fetcher = fetcher.NewHoldingMetricFetcher()
	return tracker
}

func (tracker *HoldingTracker) getHoldingTracker(holding holdings.Holding) func(chan<- metrics.HoldingMetric) {
	duration, err := time.ParseDuration(holding.RefreshTime)
	if err != nil {
		fmt.Printf("Cannot parse duration")
	}

	return func(publish chan<- metrics.HoldingMetric) {
		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		// Hack to make the first tick immediate
		for t := time.Now(); true; t = <-ticker.C {
			fmt.Printf("\nFetching %s %s at %s", holding.Name, holding.Category, t)

			holdingMetrics, err := tracker.fetcher.Fetch(holding)
			if err != nil {
				fmt.Println(err)
				continue
			}

			publish <- holdingMetrics
		}
	}
}
