package holdings

type Holdings struct {
	Equity Equity `yaml:"equity"`
}

type Equity struct {
	MutualFunds []Holding `yaml:"mutual_funds"`
	Rsu         []Holding `yaml:"rsu"`
	FD          []Holding `yaml:"fd"`
}

type Holding struct {
	Name               string  `yaml:"name"`
	Category           string  `yaml:"category"`
	UnitsHeld          float64 `yaml:"units_held"`
	StaticPricePerUnit float64 `yaml:"static_price_per_unit"`
	Api                string  `yaml:"api"`
	Symbol             string  `yaml:"symbol"`
	RefreshTime        string  `yaml:"refresh_time"`
	CacheDuration      string  `yaml:"cache_duration"`
}
