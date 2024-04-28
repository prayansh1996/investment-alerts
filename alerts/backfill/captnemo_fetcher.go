package backfill

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define a struct to match the JSON response structure
type ApiResponse struct {
	ISIN          string          `json:"ISIN"`
	Name          string          `json:"name"`
	Nav           float64         `json:"nav"`
	Date          string          `json:"date"`
	HistoricalNav [][]interface{} `json:"historical_nav"`
}

type Nav struct {
	date string
	nav  float64
}

func FetchHistoricalNav(isin string) []Nav {
	url := fmt.Sprintf("https://mf.captnemo.in/nav/%s", isin)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return nil
	}

	var historicalNav []Nav
	for _, nav := range apiResponse.HistoricalNav {
		date, okDate := nav[0].(string)
		nav, okNav := nav[1].(float64)
		if !okDate || !okNav {
			fmt.Printf("\nDate %v or Nav %v incorrect", date, nav)
			continue
		}

		fmt.Printf("\nDate: %s, NAV: %f", date, nav)
		historicalNav = append(historicalNav, Nav{date: date, nav: nav})
	}

	return historicalNav
}
