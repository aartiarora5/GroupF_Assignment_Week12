# Cryptocurrency Price Tracker - with GO
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
The main package is used for executable commands, and main is the entry point for the executable program.
import statements are used to include external packages necessary for the program.
encoding/json is used for encoding and decoding JSON data.
fmt is used for formatting and printing.
log is used for logging.
net/http provides HTTP client and server implementations.

type PriceResponse struct {
	Bitcoin  map[string]float64 `json:"bitcoin"`
	Ethereum map[string]float64 `json:"ethereum"`
	Tether   map[string]float64 `json:"tether"`
}
The PriceResponse struct is defined to represent the structure of the JSON response from the CoinGecko API.
It includes three fields: Bitcoin, Ethereum, and Tether, each containing a map with currency codes as keys and float64 values representing prices.


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
The getPrice function is an HTTP handler for the "/price" endpoint.
It sends a GET request to the CoinGecko API to retrieve cryptocurrency prices in CAD.
The JSON response is decoded into the data variable, which is an instance of the PriceResponse struct.
Prices for Bitcoin, Ethereum, and Tether are extracted from the data struct.
A new JSON response is created with the extracted prices, and it is sent as the HTTP response.


func main() {
	// price endpoint
	http.HandleFunc("/price", getPrice)
	// Start the server
	port := 7000
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
The main function is the entry point of the program.
It registers the /price endpoint with the getPrice handler function.
The HTTP server is started on port 7000, and the server is run indefinitely.
If there is an error starting the server, it is logged.

The provided code defines a simple Go application that fetches cryptocurrency prices from the CoinGecko API and exposes them through a RESTful API on the "/price" endpoint. The program sets up an HTTP server, listens on port 7000, and responds to requests for cryptocurrency prices in CAD. The getPrice function handles the logic for fetching data from the CoinGecko API and formatting the response.