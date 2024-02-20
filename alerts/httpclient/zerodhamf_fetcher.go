package httpclient

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/prayansh1996/investment-alerts/holdings"
	"github.com/prayansh1996/investment-alerts/metrics"
)

func zerodhaMfFetcher(fund holdings.Fund, body []byte) (metrics.Metric, error) {
	records, err := csv.NewReader(bytes.NewReader(body)).ReadAll()
	if err != nil {
		fmt.Printf("Unable to read records from csv: %s\n", err)
		return metrics.Metric{}, err
	}

	nav := 0.0
	for _, record := range records {
		if record[0] == fund.Symbol {
			nav, _ = strconv.ParseFloat(record[14], 64)
		}
	}

	return metrics.Metric{
		Units:        fund.UnitsHeld,
		PricePerUnit: nav,
		Name:         fund.Name,
		Category:     fund.Category,
	}, nil
}
