package holdings

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

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

var holdings Holdings

func InitializeHoldings() {
	yamlFile, err := os.ReadFile("holdings.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}

	err = yaml.Unmarshal(yamlFile, &holdings)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Parsed Holdings: %+v\n", holdings)
}
