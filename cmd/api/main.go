package main

import (
	"log"
	"net/http"

	"github.com/ericoalmeida/go-wheather/internal/config"
	"github.com/ericoalmeida/go-wheather/internal/handlers"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/weather", handlers.GetCurrentWeatherHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
