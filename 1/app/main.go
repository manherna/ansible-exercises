package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Coding the response gotten from OWM to a struct for easier handling
type OWMResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

// Implementing string transformation directly in the OMWResponse object
func (in OWMResponse) String() string {
	return fmt.Sprintf("city=\"%s\", description=\"%s\", temp=%.1f, humidity=%d", in.Name, in.Weather[0].Description, in.Main.Temp, in.Main.Humidity)
}

func main() {
	apiKey := os.Getenv("OWM_API_KEY")
	city := os.Getenv("OWM_CITY")
	if len(apiKey) == 0 || len(city) == 0 {
		println("no correct environment variables were given, exiting")
		os.Exit(1)
	}
	println(getWeather(apiKey, city))

}

// Basic HTTP GET method + conversion to our OWMResponse object. Returns a string with the required weather format
func getWeather(apiKey string, city string) string {
	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, apiKey))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	var result OWMResponse
	json.Unmarshal([]byte(body), &result)
	if result.Cod != 200 {
		return fmt.Sprintf("There was a problem fetching results. HTTP Satus code: %d", result.Cod)
	}
	return result.String()
}
