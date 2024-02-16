package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func main() {
	// API endpoint
	url := "https://api.mfapi.in/mf/119063/latest"

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making the request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading the response body: %s\n", err)
		return
	}

	// Unmarshal the JSON response into the ApiResponse struct
	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling the response: %s\n", err)
		return
	}

	// Print the extracted data
	fmt.Printf("Fund House: %s\n", apiResponse.Meta.FundHouse)
	fmt.Printf("Scheme Type: %s\n", apiResponse.Meta.SchemeType)
	fmt.Printf("Scheme Category: %s\n", apiResponse.Meta.SchemeCategory)
	fmt.Printf("Scheme Code: %d\n", apiResponse.Meta.SchemeCode)
	fmt.Printf("Scheme Name: %s\n", apiResponse.Meta.SchemeName)
	for _, data := range apiResponse.Data {
		fmt.Printf("Date: %s, NAV: %s\n", data.Date, data.Nav)
	}
}
