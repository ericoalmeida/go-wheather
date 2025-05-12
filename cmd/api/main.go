package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ericoalmeida/go-wheather/internal/clients"
	"github.com/ericoalmeida/go-wheather/internal/config"
	"github.com/ericoalmeida/go-wheather/pkg"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/weather", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if !pkg.IsZipcodeValid(cep) {
		http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	coordinates, err := clients.SearchCEPCoordinates(cep)
	if err != nil {
		http.Error(w, "No zipcode info found.", http.StatusNotFound)
		return
	}

	currentWeather, err := clients.GetCurrentWeather(coordinates.Lat, coordinates.Lon)
	if err != nil {
		http.Error(w, "No matching location found.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(currentWeather)
}
