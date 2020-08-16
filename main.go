package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var currenciesURL = "https://focusmobile-interview-materials.s3.eu-west-3.amazonaws.com/Cheap.Stocks.Internationalization.Currencies.csv"

func getDataFromURL(url string) ([]string, error) {
	// Make a http request to ge the currencies
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		// Return an empty list and the error if we've got any
		return nil, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	// Convert the body to a string and split by new line
	csvLines := strings.Split(string(respBody), "\r\n")

	return csvLines, nil
}

func bindResponseToCurrency(response []string) map[string]Currency {
	currencies := make(map[string]Currency)
	for i, line := range response {
		if i == 0 {
			continue
		}
		rowData := strings.Split(line, ",")
		code := rowData[2]
		currencies[code] = Currency{
			Country:     rowData[0],
			Name:        rowData[1],
			Code:        code,
			LastFetchAt: time.Now(),
		}
	}
	return currencies

}

func displayResponses(inputs []string, currencies map[string]Currency) {
	for _, input := range inputs {
		if currency, ok := currencies[input]; ok {
			fmt.Printf("Query OK\nLoaded at:\t %v \nCurrency Code:\t %s \nCountry:\t %s\nCurrency:\t %s\n", currency.LastFetchAt, input, currency.Country, currency.Name)
			fmt.Println("****************************************************************")
		} else {
			fmt.Printf("Query FAILED\nCurrency %s is not supported\n", input)
			fmt.Println("****************************************************************")
		}
	}

}

func main() {
	println("Cheap Stocks, Inc currency checker")
	println("Special Commands:")
	println("exit to exit the application")
	println("reload to reload data from  the application")
	println("help to get extra information about the commands")
	response, err := getDataFromURL(currenciesURL)
	if err != nil {
		log.Fatal("Unable to load currencies from remote server")
	}
	currencies := bindResponseToCurrency(response)

Loop:
	for {
		for {
			currencyCode := strings.TrimSpace(requestInput("Enter a command or currency code(s)"))
			inputs := strings.Split(currencyCode, ",")
			switch inputs[0] {
			case "exit":
				println("bye...")
				break Loop
			case "reload":
				reloadResponse, _err := getDataFromURL(currenciesURL)
				if _err != nil {
					println("An error occured while reloading currencies.")
					println("Try again or exit the application, previous currencies can still be used")
					continue Loop
				}
				fmt.Printf("Reload successful at %s :", time.Now())
				response = reloadResponse
				currencies = bindResponseToCurrency(response)
				continue Loop
			case "help":
				println("Cheap Stocks, Inc currency checker")
				println("Special Commands:")
				println("exit to exit the application")
				println("reload to reload data from  the application")
				println("Usage of search : input Single code for single search and \n comma separated input for multisearch i.e. JPY,BGP,USD")

			case "":
				println("Cannot process empty input")
			default:
				displayResponses(inputs, currencies)
			}

		}
	}

}
