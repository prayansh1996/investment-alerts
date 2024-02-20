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
	MutualFunds []Fund  `yaml:"mutual_funds"`
	Stocks      []Stock `yaml:"stocks"`
}

type Fund struct {
	Name          string  `yaml:"name"`
	Category      string  `yaml:"category"`
	UnitsHeld     float64 `yaml:"units_held"`
	Api           string  `yaml:"api"`
	Symbol        string  `yaml:"symbol"`
	RefreshTime   string  `yaml:"refresh_time"`
	CacheDuration string  `yaml:"cache_duration"`
}

type Stock struct {
	Name        string `yaml:"name"`
	Ticker      string `yaml:"ticker"`
	UnitsHeld   int    `yaml:"units_held"`
	Api         string `yaml:"api"`
	Id          string `yaml:"id"`
	RefreshTime string `yaml:"refresh_time"`
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

func GetHoldings() Holdings {
	return holdings
}
