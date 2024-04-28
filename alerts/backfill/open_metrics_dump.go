package backfill

import (
	"fmt"
	"os"
	"time"
)

func CreateOpenMetricsFile(historicalNavs []Nav) {
	// Open or create the file
	file, err := os.Create("./backfill/open_metrics.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the header for the metric
	_, err = file.WriteString("# TYPE price_per_unit gauge\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	_, err = file.WriteString("# HELP price_per_unit The NAV price per unit for a given date.\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Iterate over historical NAV data and write each as a metric
	for _, historicalNav := range historicalNavs {
		dateStr, nav := historicalNav.date, historicalNav.nav
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}
		timestamp := date.Unix()
		metric := fmt.Sprintf(
			"price_per_unit{category=\"Index Fund\",instance=\"alerts_go:8082\", job=\"alerts\", name=\"HDFC Nifty 50 Growth Index\"} %f %d\n",
			nav,
			timestamp,
		)
		_, err = file.WriteString(metric)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Metrics successfully written to metrics.txt")
}
