package fetcher

import "github.com/prayansh1996/investment-alerts/holdings"

type FixedDepositHoldingFetcher struct {
}

func (fd *FixedDepositHoldingFetcher) Fetch() []holdings.Holding {
	return parsedHoldings.Equity.FD
}
