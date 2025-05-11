package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ericoalmeida/go-wheather/internal/clients"
	"github.com/ericoalmeida/go-wheather/internal/config"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "Informe o cep", http.StatusBadRequest)
		return
	}

	coordinates, err := clients.SearchCEPCoordinates(cep)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(coordinates)
}
