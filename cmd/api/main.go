package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ericoalmeida/go-wheather/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg.GeoapifyBaseUrl)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"

	json.NewEncoder(w).Encode(message)
}
