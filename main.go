package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var currencies map[string]Currency

var languages map[string]Language

// defaultCurrency := "USD"
// defaultLanguage := "en"

func getDataFromURL(url string) (string, error) {
	// Make a http request to ge the currencies
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		// Return an empty list and the error if we've got any
		return "", err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	return string(respBody), nil
}

func csvLinesFromString(csvfile string) []string {
	// Convert the body to a string and split by new line
	return strings.Split(csvfile, "\r\n")
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

func biindResponseToLanguage(response []string) map[string]Language {
	languages := make(map[string]Language)
	for i, line := range response {
		if i == 0 {
			continue
		}
		rowData := strings.Split(line, ",")
		code := rowData[1]
		languages[code] = Language{
			Name:        rowData[0],
			Code:        rowData[1],
			LastFetchAt: time.Now(),
		}
	}
	return languages
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

func list(input string) {
	switch input {
	case "currencies":
		// TODO: Function to list currencies
	case "languages":
		// TODO: Add function to list langages
	default:
		//TODO: Invalid selection

	}
}

func set(inputs []string) {

}

func convert(inputs []string) {

}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	println("Cheap Stocks, Inc currency checker")
	println("Special Commands:")
	println("exit to exit the application")
	println("reload to reload data from  the application")
	println("help to get extra information about the commands")
	currencyResponse, err := getDataFromURL(CurrenciesURL)
	langResponse, err := getDataFromURL(LanguagesURL)
	if err != nil {
		log.Fatal("Unable to load currencies from remote server")
	}
	processedCsvCurrency := csvLinesFromString(currencyResponse)
	currencies = bindResponseToCurrency(processedCsvCurrency)

	processedCsvLang := csvLinesFromString(langResponse)
	languages = biindResponseToLanguage(processedCsvLang)

Loop:
	for {
		currencyCode := strings.TrimSpace(requestInput("Enter a command or currency code(s)"))
		inputs := strings.Split(currencyCode, " ")
		switch inputs[0] {
		case "exit":
			println("bye...")
			break Loop
		case "reload":
			reloadResponse, _err := getDataFromURL(CurrenciesURL)
			if _err != nil {
				println("An error occured while reloading currencies.")
				println("Try again or exit the application, previous currencies can still be used")
			}
			fmt.Printf("Reload successful at %s :", time.Now())
			currencyResponse = reloadResponse
			processedCsv := csvLinesFromString(currencyResponse)
			currencies = bindResponseToCurrency(processedCsv)
		case "help":
			println("Cheap Stocks, Inc currency checker")
			println("Special Commands:")
			println("exit to exit the application")
			println("reload to reload data from  the application")
			println("Usage of search : input Single code for single search and \n space separated input for multisearch i.e. JPY BGP USD")
		case "list":
			list(inputs[1])
		case "set":
			commandData := remove(inputs, 0)
			set(commandData)
		case "convert":
			commandData := remove(inputs, 0)
			convert(commandData)
		case "":
			println("Cannot process empty input")
		default:
			displayResponses(inputs, currencies)
		}

	}

}
