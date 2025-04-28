package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ExchangeRateResponse struct {
	Result  string            `json:"result"`
	Rate    float64           `json:"conversion_rate"`
	Error   string            `json:"error"`
}

func main() {
	var amount float64
	var fromCurrency, toCurrency string

	fmt.Print("Enter the amount: ")
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Invalid input for amount.")
		return
	}

	fmt.Print("Enter the currency code you want to convert from (e.g., USD): ")
	_, err = fmt.Scan(&fromCurrency)
	if err != nil {
		fmt.Println("Invalid input for currency code.")
		return
	}

	fmt.Print("Enter the currency code you want to convert to (e.g., EUR): ")
	_, err = fmt.Scan(&toCurrency)
	if err != nil {
		fmt.Println("Invalid input for currency code.")
		return
	}

	fromCurrency = strings.ToUpper(fromCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	apiKey := "YOUR_API_KEY" 
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s/%f", apiKey, fromCurrency, toCurrency, amount)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching exchange rate:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var exchangeRateResp ExchangeRateResponse
	err = json.Unmarshal(body, &exchangeRateResp)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	if exchangeRateResp.Result != "success" {
		fmt.Println("Error: ", exchangeRateResp.Error)
		return
	}

	convertedAmount := amount * exchangeRateResp.Rate

	fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
