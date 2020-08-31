package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var currencies map[string]Currency

var languages map[string]Language

var defaultCurrency string = "USD"
var defaultLanguage string = "en"

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

func displayLanguages() {
	for code, language := range languages {
		fmt.Printf("Language\nLoaded at:\t %v \nCode:\t %s \nName:\t %s \n", language.LastFetchAt, code, language.Name)
	}
}

func displayCurrencies() {
	for code, currency := range currencies {
		fmt.Printf("Query OK\nLoaded at:\t %v \nCurrency Code:\t %s \nCountry:\t %s\nCurrency:\t %s\n", currency.LastFetchAt, code, currency.Country, currency.Name)
	}
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

func currencyCheck(key string) {
	if _, ok := currencies[key]; ok {
		defaultCurrency = key
		fmt.Printf("Default value for currency set to :  %s\n", key)
		fmt.Println("****************************************************************")
	} else {
		fmt.Printf("Invalid command flag %s\nCommand  is not supported\n", key)
		fmt.Println("****************************************************************")
	}
}

func languageCheck(key string) {
	if _, ok := languages[key]; ok {
		defaultLanguage = key
		fmt.Printf("Default value for language set to :  %s\n", key)
		fmt.Println("****************************************************************")
	} else {
		fmt.Printf("Invalid command flag %s\nCommand  is not supported\n", key)
		fmt.Println("****************************************************************")
	}
}

func convert(inputs []string) {

}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func help() {
	println("Cheap Stocks, Inc currency checker")
	println("Special Commands:")
	println("exit to exit the application")
	println("reload {flag} to reload data from  the application based on input flag\n available flags are currencies and languages\n i.e. \"reload currencies\" or \"reload languages\" ")
	println("list {flag} to list data from  the application based on input flag\n available flags are currencies and languages\n i.e. \"list currencies\" or \"list languages\" ")
	println("set {flag}=<value> ... used to setup default values for one or more functions.\n Available flags are currency and language\n i.e \"set currency=USD language=en\" ")
	println("Usage of search : input Single currency code for single search and \n space separated input for currency code multisearch i.e. JPY BGP USD")
}

func loadCurrencies() {
	reloadResponse, _err := getDataFromURL(CurrenciesURL)
	if _err != nil {
		println("An error occured while reloading currencies.")
		println("Try again or exit the application, if application had started previous currencies can be still used")
	}
	fmt.Printf("Reload successful at %s\n :", time.Now())
	processedCsv := csvLinesFromString(reloadResponse)
	currencies = bindResponseToCurrency(processedCsv)
}

func loadLanguages() {
	reloadResponse, _err := getDataFromURL(LanguagesURL)
	if _err != nil {
		println("An error occured while reloading languages.")
		println("Try again or exit the application, if application had already started previous loaded languages can be used.")
	}
	fmt.Printf("Reload successful at %s\n :", time.Now())
	processedCsv := csvLinesFromString(reloadResponse)
	languages = biindResponseToLanguage(processedCsv)
}

func list(input string) {
	switch input {
	case "currencies":
		displayCurrencies()
	case "languages":
		displayLanguages()
	default:
		fmt.Printf("Invalid command flag %s\nCommand  is not supported\n", input)
		fmt.Println("****************************************************************")
	}
}

func set(inputs []string) {
	for _, setcombo := range inputs {
		setkv := strings.Split(setcombo, "=")
		switch setkv[0] {
		case "currency":
			currencyCheck(setkv[1])
		case "language":
			languageCheck(setkv[1])
		default:
			fmt.Printf("Invalid command flag %s\nCommand  is not supported\n", setkv[0])
			fmt.Println("****************************************************************")
		}
	}
}

func reload(flag string) {
	switch flag {
	case "currencies":
		loadCurrencies()
	case "languages":
		loadLanguages()
	default:
		fmt.Printf("Invalid command flag %s\nCommand  is not supported\n", flag)
		fmt.Println("****************************************************************")
	}

}

func main() {
	println("Cheap Stocks, Inc currency checker")
	println("Special Commands:")
	println("exit to exit the application")
	println("reload to reload data from  the application")
	println("help to get extra information about the commands")
	loadCurrencies()
	loadLanguages()

Loop:
	for {
		currencyCode := strings.TrimSpace(requestInput("Enter a command or currency code(s)"))
		inputs := strings.Split(currencyCode, " ")
		switch inputs[0] {
		case "exit":
			println("bye...")
			break Loop
		case "reload":
			reload(inputs[1])
		case "help":
			help()
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
