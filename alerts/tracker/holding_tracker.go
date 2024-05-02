package tracker

import (
	"fmt"
	"time"

	"github.com/prayansh1996/investment-alerts/cons"
	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/holdings/fetcher"
	"github.com/prayansh1996/investment-alerts/metrics"
	"github.com/prayansh1996/investment-alerts/metrics/holdingmetric"
)

func Start() {
	ticker := time.NewTicker(cons.HOLDING_NAV_REFRESH_TIME)
	defer ticker.Stop()

	// Hack to make the first tick immediate
	for t := time.Now(); true; t = <-ticker.C {
		holdingsList := []holdings.Holding{}
		holdingsList = append(holdingsList, (&fetcher.FixedDepositHoldingFetcher{}).Fetch()...)
		holdingsList = append(holdingsList, (&fetcher.MutualFundHoldingFetcher{}).Fetch()...)
		holdingsList = append(holdingsList, (&fetcher.RsuHoldingFetcher{}).Fetch()...)

		for _, holding := range holdingsList {
			fmt.Printf("\nFetching %s %s at %s", holding.Name, holding.Category, t)

			holdingMetric, err := holdingmetric.NewHoldingMetricFetcher(holding).Fetch()
			if err != nil {
				fmt.Println(err)
				continue
			}

			metrics.PublishChannel <- holdingMetric
		}
	}
}
