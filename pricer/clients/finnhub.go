package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type PriceResponse struct {
	Current   float64 `json:"c"`
	Change    float64 `json:"d"`
	Percent   float64 `json:"dp"`
	High      float64 `json:"h"`
	Low       float64 `json:"l"`
	Open      float64 `json:"o"`
	PrevClose float64 `json:"pc"`
	Timestamp int64   `json:"t"`
}

func GetCurrentPrice(ticker string) (PriceResponse, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return PriceResponse{}, fmt.Errorf("API_KEY not set in environment")
	}

	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", ticker, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return PriceResponse{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PriceResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return PriceResponse{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if data.Current == 0 {
		return PriceResponse{}, fmt.Errorf("no valid price returned for ticker %s", ticker)
	}

	return data, nil
}
