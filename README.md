# Cryptocurrency Price Tracker - with GO

## How to set up
Step 1:
Clone the project from GitHub with the url [https://github.com/aartiarora5/GroupF_Assignment_Week12]
 --- Command to clone is: git clone github url

Step 2:
Cd into the main folder of the project

Step 3:
Write the following command to run the application on localhost
 --- localhost:7000/price

The Port Number we used it 7000


## Explanation on the Project
getPrice Function: The get price function fetch the data from the external API "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,tether&vs_currencies=cad". We then extracted the bitcoin, etherium price and tether price.

In the main function, we then create a get request (local api) with port number "7000" and url path of "/price", and displayed the price of the three cyptocurrencies

