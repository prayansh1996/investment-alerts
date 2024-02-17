package metrics

import (
	"github.com/prayansh1996/investment-alerts/cons"
	"github.com/prometheus/client_golang/prometheus"
)

// Define a gauge outside of the main function to make it accessible in the handler
var unitsGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "units",
		Help: "Units of this stock",
	},
	[]string{cons.Name, cons.Category},
)

// Define a gauge outside of the main function to make it accessible in the handler
var pricePerUnitGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "price_per_unit",
		Help: "Units of this stock",
	},
	[]string{cons.Name, cons.Category},
)

func init() {
	// Register the gauge with Prometheus
	prometheus.MustRegister(unitsGauge)
	prometheus.MustRegister(pricePerUnitGauge)
}

func Publish(units int, pricePerUnit int, name string, category string) {
	unitsGauge.With(prometheus.Labels{cons.Name: name, cons.Category: category}).Add(float64(units))
	pricePerUnitGauge.With(prometheus.Labels{cons.Name: name, cons.Category: category}).Add(float64(pricePerUnit))
}
