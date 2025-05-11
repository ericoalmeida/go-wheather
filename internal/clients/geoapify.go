package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ericoalmeida/go-wheather/internal/config"
)

type GeoapifyResponse struct {
	Results []struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"results"`
}

type CEPCoordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

func SearchCEPCoordinates(cep string) (*CEPCoordinates, error) {
	baseUrl := config.GetEnv("GEOAPIFY_BASE_URL")
	apiKey := config.GetEnv("GEOAPIFY_API_KEY")

	q := url.QueryEscape(cep + ", Brasil")

	apiURL := fmt.Sprintf("%s/v1/geocode/search?text=%s&format=json&apiKey=%s", baseUrl, q, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error on request try: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error on api request: %s", resp.Status)
	}

	var geoapifyResponse GeoapifyResponse
	err = json.NewDecoder(resp.Body).Decode(&geoapifyResponse)
	if err != nil {
		return nil, fmt.Errorf("error on api request: %w", err)
	}

	if len(geoapifyResponse.Results) == 0 {
		return nil, fmt.Errorf("nothing to show")
	}

	lon := geoapifyResponse.Results[0].Lon
	lat := geoapifyResponse.Results[0].Lat

	return &CEPCoordinates{
		Lon: lon,
		Lat: lat,
	}, nil
}
