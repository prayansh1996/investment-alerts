package fetcher

import "github.com/prayansh1996/investment-alerts/holdings"

type MutualFundHoldingFetcher struct {
}

func (mf *MutualFundHoldingFetcher) Fetch() []holdings.Holding {
	return parsedHoldings.Equity.MutualFunds
}
