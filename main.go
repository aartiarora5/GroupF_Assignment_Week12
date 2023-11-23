package main

import (
	"fmt"
	"log"
	"net/http"
)

type PriceResponse struct {
	Bitcoin  map[string]float64 `json:"bitcoin"`
	Ethereum map[string]float64 `json:"ethereum"`
	Tether   map[string]float64 `json:"tether"`
}

func main() {
	// Start the server
	port := 7000
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}