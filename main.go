package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WeatherResponse struct to parse the JSON data returned by the API.
type WeatherResponse struct {
    Main struct {
        Temp float64 `json:"temp"` // Temperature
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"` // Weather conditions
    } `json:"weather"`
    Wind struct {
        Speed float64 `json:"speed"` // Wind speed
    } `json:"wind"`
    Cod     int    `json:"cod"`     // HTTP status code
    Message string `json:"message"` // Error message
}

func main() {
    fmt.Print("Enter the city name: ")
    var city string
    fmt.Scanln(&city)

    apiKey := "YOUR_API_KEY_HERE" // Use your actual OpenWeatherMap API key
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Failed to request weather data:", err)
        return
    }
    defer resp.Body.Close()

    var weather WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
        fmt.Println("Failed to decode weather data:", err)
        return
    }

    // Check if the response contains an error message
    if weather.Cod != 200 {
        fmt.Printf("Failed to fetch weather data: %s\n", weather.Message)
        return
    }

    // Ensure there is at least one weather condition before accessing it
    if len(weather.Weather) == 0 {
        fmt.Println("No weather information available for the provided city.")
        return
    }

    fmt.Printf("Weather in %s: %s with a temperature of %.1f°C\n", city, weather.Weather[0].Description, weather.Main.Temp)
    fmt.Printf("Wind speed: %.1fm/s\n", weather.Wind.Speed)
}
