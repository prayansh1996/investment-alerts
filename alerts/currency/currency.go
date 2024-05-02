package currency

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func GetUsdToInrConversionRate() float64 {
	url := readCurrencyConversionApiUrl()
	fmt.Println("Url for currency conversion: " + url)

	resp, _ := getHttpResponse(url)
	var currencyResponse CurrencyResponse
	err := json.Unmarshal([]byte(resp), &currencyResponse)
	if err != nil {
		fmt.Println("Error in currency conversion:", err)
		return 0
	}

	inrRate, ok := currencyResponse.Rates["INR"]
	if !ok {
		fmt.Println("INR rate not found")
		return 0
	}

	fmt.Printf("USD to INR conversion rate: %v", inrRate)
	return inrRate
}

type CurrencyResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func getHttpResponse(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

type CurrencyConversion struct {
	Url string `yaml:"url"`
}

func readCurrencyConversionApiUrl() string {
	yamlFile, err := os.ReadFile("./config/currency_conversion.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}

	conversion := CurrencyConversion{}

	err = yaml.Unmarshal(yamlFile, &conversion)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	return conversion.Url
}
