package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const apiKey = "YOUR_API_KEY_HERE"

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather(city string) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("city not found or API error (Status: %s)", resp.Status)
	}

	var weatherData WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &weatherData, nil
}

func main() {
	var city string
	fmt.Print("Enter a city name: ")
	fmt.Scanln(&city)

	weather, err := getWeather(city)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nWeather in %s:\n", weather.Name)
	fmt.Printf("Temperature: %.1f°C\n", weather.Main.Temp)
	fmt.Printf("Condition: %s\n", weather.Weather[0].Description)
}
