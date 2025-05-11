package main

import (
	"fmt"

	"github.com/ericoalmeida/go-wheather/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println(cfg.GeoapifyBaseUrl)
}
