package main

import (
    "fmt"
    "log"
    "os"
    "github.com/go-resty/resty/v2"
    "github.com/joho/godotenv"
    "encoding/json"
)

const apiURL = "https://openexchangerates.org/api/latest.json"

func getExchangeRates(apiKey string) (map[string]float64, error) {
    client := resty.New()
    resp, err := client.R().
        SetQueryParam("app_id", apiKey).
        Get(apiURL)
    if err != nil {
        return nil, err
    }

    var result map[string]interface{}
    err = json.Unmarshal(resp.Body(), &result)
    if err != nil {
        return nil, err
    }

    rates := result["rates"].(map[string]interface{})
    exchangeRates := make(map[string]float64)
    for _, currency := range []string{"USD", "EUR", "JPY", "GBP", "AUD", "CAD", "CHF", "CNY", "SEK", "NZD"} {
        exchangeRates[currency] = rates[currency].(float64)
    }

    return exchangeRates, nil
}

func findProfitableExchange(rates map[string]float64) {
    for base, baseRate := range rates {
        for quote1, rate1 := range rates {
            if base == quote1 {
                continue
            }
            for quote2, rate2 := range rates {
                if quote1 == quote2 || base == quote2 {
                    continue
                }
                // Calculate potential profit using triangular arbitrage
                potentialProfit := (1 / rate1) * (1 / rate2) * baseRate
                if potentialProfit > 1 {
                    fmt.Printf("Profitable exchange found: %s -> %s -> %s -> %s with profit %.2f%%\n",
                        base, quote1, quote2, base, (potentialProfit-1)*100)
                }
            }
        }
    }
}

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    apiKey := os.Getenv("OPENEXCHANGERATES_API_KEY")
    if apiKey == "" {
        log.Fatalf("API key not found. Set the OPENEXCHANGERATES_API_KEY environment variable.")
    }

    rates, err := getExchangeRates(apiKey)
    if err != nil {
        log.Fatalf("Error fetching exchange rates: %v", err)
    }

    findProfitableExchange(rates)
}

