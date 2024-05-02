package holdings

type FDHolding struct {
}

func GetFixedDepostHoldings() []Holding {
	return holdings.Equity.FD
}
