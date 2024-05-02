package metrics

import (
	"fmt"

	"github.com/prayansh1996/investment-alerts/cons"
	"github.com/prometheus/client_golang/prometheus"
)

type HoldingMetric struct {
	Units        float64
	PricePerUnit float64
	Name         string
	Category     string
}

var PublishChannel = make(chan HoldingMetric)

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

func InitializeMetrics() {
	// Register the gauge with Prometheus
	prometheus.MustRegister(unitsGauge)
	prometheus.MustRegister(pricePerUnitGauge)

	go initializePublisher()
}

func initializePublisher() {
	for {
		metric := <-PublishChannel
		publish(metric)
	}
}

func publish(metric HoldingMetric) {
	if metric.Name == "" {
		fmt.Print("Empty metric received")
		return
	}
	fmt.Printf("\nFlushing Metric: %v", metric)

	unitsGauge.With(prometheus.Labels{cons.Name: metric.Name, cons.Category: metric.Category}).Set(metric.Units)
	pricePerUnitGauge.With(prometheus.Labels{cons.Name: metric.Name, cons.Category: metric.Category}).Set(metric.PricePerUnit)
}
