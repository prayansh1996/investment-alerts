package fetcher

import (
	"fmt"
	"log"
	"os"

	"github.com/prayansh1996/investment-alerts/holdings"
	"gopkg.in/yaml.v2"
)

type HoldingFetcher interface {
	Fetch() holdings.Holding
}

var parsedHoldings holdings.Holdings

func InitializeHoldings() {
	yamlFile, err := os.ReadFile("holdings.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}

	err = yaml.Unmarshal(yamlFile, &parsedHoldings)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Parsed Holdings: %+v\n", parsedHoldings)
}
