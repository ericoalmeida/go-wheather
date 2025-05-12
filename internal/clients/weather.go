package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ericoalmeida/go-wheather/internal/config"
	"github.com/ericoalmeida/go-wheather/pkg"
)

type WeatherResponse struct {
	Current struct {
		Temp_c float64 `json:"temp_c"`
		Temp_f float64 `json:"temp_f"`
	} `json:"current"`
}

type CurrentWeather struct {
	Temp_c float64 `json:"temp_C"`
	Temp_f float64 `json:"temp_F"`
	Temp_k float64 `json:"temp_K"`
}

func GetCurrentWeather(lat, lon float64) (*CurrentWeather, error) {
	baseUrl := config.GetEnv("WEATHER_BASE_URL")
	apiKey := config.GetEnv("WEATHER_API_KEY")

	q := fmt.Sprintf("%f,%f", lat, lon)

	apiURL := fmt.Sprintf("%s/v1/current.json?key=%s&q=%s&aqi=no", baseUrl, apiKey, q)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error on request try: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error on api request: %s", resp.Status)
	}

	var weatherResponse WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return nil, fmt.Errorf("error on api request: %w", err)
	}

	tempC := weatherResponse.Current.Temp_c
	tempF := weatherResponse.Current.Temp_f
	tempK := pkg.CelsiusToKelvinConverter(weatherResponse.Current.Temp_c)

	return &CurrentWeather{
		Temp_c: tempC,
		Temp_f: tempF,
		Temp_k: tempK,
	}, nil
}
