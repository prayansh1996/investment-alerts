package fetcher

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/prayansh1996/investment-alerts/holdings"
	"gopkg.in/yaml.v2"
)

type RsuHoldingFetcher struct{}

func (rsu *RsuHoldingFetcher) Fetch() []holdings.Holding {
	parsedHoldings.Equity.Rsu[0].UnitsHeld = parsedHoldings.Equity.Rsu[0].UnitsHeld + getAdditionalUnits()
	return parsedHoldings.Equity.Rsu
}

type Vesting struct {
	Google []struct {
		Date string `yaml:"date"`
		RSU  int    `yaml:"rsu"`
	} `yaml:"google"`
}

func getAdditionalUnits() float64 {
	yamlFile, err := os.ReadFile("google_vesting_schedule.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}

	vesting := Vesting{}
	err = yaml.Unmarshal(yamlFile, &vesting)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	currentDate := time.Now()
	totalAdditionalRSU := 0.0

	fmt.Printf("\nVesting: %v", vesting.Google)

	for _, entry := range vesting.Google {
		date, err := time.Parse("January 2006", entry.Date)
		if err != nil {
			log.Fatalf("Error parsing date: %v", err)
		}

		if currentDate.After(date) {
			totalAdditionalRSU += float64(entry.RSU)
		}
	}

	return totalAdditionalRSU
}
