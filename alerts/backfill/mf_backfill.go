package backfill

type MfBackfill struct {
	name string
}

func NewMfBackfill(name string) *MfBackfill {
	return &MfBackfill{
		name: name,
	}
}

func (m *MfBackfill) BackfillPrices() {
	CreateOpenMetricsFile(FetchHistoricalNav("INF789F01XA0"))
}
