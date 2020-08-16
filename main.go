package main

import (
	"io/ioutil"
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

func main() {

}
