package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PriceResponse struct {
	Bitcoin  map[string]float64 `json:"bitcoin"`
	Ethereum map[string]float64 `json:"ethereum"`
	Tether   map[string]float64 `json:"tether"`
}

func getPrice(w http.ResponseWriter, r *http.Request) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,tether&vs_currencies=cad"

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data from CoinGecko API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var data PriceResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error decoding JSON response", http.StatusInternalServerError)
		return
	}

	// Extract prices
	bitcoinPrice := data.Bitcoin["cad"]
	ethereumPrice := data.Ethereum["cad"]
	tetherPrice := data.Tether["cad"]

	// Create a JSON response
	response := map[string]float64{
		"bitcoin":  bitcoinPrice,
		"ethereum": ethereumPrice,
		"tether":   tetherPrice,
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// price endpoint
	http.HandleFunc("/price", getPrice)
	// Start the server
	port := 7000
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
