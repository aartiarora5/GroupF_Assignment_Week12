# Cryptocurrency Price Tracker - with GO

The main package in the main file is used for executable commands, and main is the entry point for the executable program.
import statements are used to include external packages necessary for the program.
encoding/json is used for encoding and decoding JSON data.
fmt is used for formatting and printing.
log is used for logging.
net/http provides HTTP client and server implementations.


The PriceResponse struct is defined to represent the structure of the JSON response from the CoinGecko API.
It includes three fields: Bitcoin, Ethereum, and Tether, each containing a map with currency codes as keys and float64 values representing prices.


The getPrice function is an HTTP handler for the "/price" endpoint.
It sends a GET request to the CoinGecko API to retrieve cryptocurrency prices in CAD.
The JSON response is decoded into the data variable, which is an instance of the PriceResponse struct.
Prices for Bitcoin, Ethereum, and Tether are extracted from the data struct.
A new JSON response is created with the extracted prices, and it is sent as the HTTP response.


The main function is the entry point of the program.
It registers the /price endpoint with the getPrice handler function.
The HTTP server is started on port 7000, and the server is run indefinitely.
If there is an error starting the server, it is logged.

The provided code defines a simple Go application that fetches cryptocurrency prices from the CoinGecko API and exposes them through a RESTful API on the "/price" endpoint. The program sets up an HTTP server, listens on port 7000, and responds to requests for cryptocurrency prices in CAD. The getPrice function handles the logic for fetching data from the CoinGecko API and formatting the response.