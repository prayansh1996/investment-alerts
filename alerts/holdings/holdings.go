package holdings

import (
	"fmt"
	"io/ioutil"
	"log"

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
	Name        string `yaml:"name"`
	Category    string `yaml:"category"`
	UnitsHeld   int    `yaml:"units_held"`
	Api         string `yaml:"api"`
	RefreshTime string `yaml:"refresh_time"`
}

type Stock struct {
	Name        string `yaml:"name"`
	Ticker      string `yaml:"ticker"`
	UnitsHeld   int    `yaml:"units_held"`
	Api         string `yaml:"api"`
	RefreshTime string `yaml:"refresh_time"`
}

var holdings Holdings

func init() {
	// Register the gauge with Prometheus
	// Assuming holdings.yaml is in the current directory
	yamlFile, err := ioutil.ReadFile("holdings.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}

	err = yaml.Unmarshal(yamlFile, &holdings)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Parsed Holdings: %+v\n", holdings)
}

func GetMutualFunds() []Fund {
	return holdings.Equity.MutualFunds
}
