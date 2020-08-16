package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var currenciesURL = "https://focusmobile-interview-materials.s3.eu-west-3.amazonaws.com/Cheap.Stocks.Internationalization.Currencies.csv"

func getCurrencies(url string) ([]string, error) {
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

func main() {

}
