package fetchers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/prayansh1996/investment-alerts/metrics"
)

// Define the structs to match the JSON response
type ApiResponse struct {
	Meta   MetaData  `json:"meta"`
	Data   []NavData `json:"data"`
	Status string    `json:"status"`
}

type MetaData struct {
	FundHouse      string `json:"fund_house"`
	SchemeType     string `json:"scheme_type"`
	SchemeCategory string `json:"scheme_category"`
	SchemeCode     int    `json:"scheme_code"`
	SchemeName     string `json:"scheme_name"`
}

type NavData struct {
	Date string `json:"date"`
	Nav  string `json:"nav"`
}

func fetchNav() {
	url := "https://api.mfapi.in/mf/119063/latest"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the response body: %s\n", err)
		return
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling the response: %s\n", err)
		return
	}

	// Update the gauge with the fetched NAV
	if len(apiResponse.Data) > 0 {
		nav, err := strconv.ParseFloat(apiResponse.Data[0].Nav, 64)
		if err == nil {
			metrics.Publish(int(nav), 1, "HDFC", "Index Fund")
		}
	}
}

func Start() {
	for i := 0; i <= 100; i++ {
		fetchNav()
		time.Sleep(2 * time.Second)
		fmt.Println("Nav fetched")
	}
}
