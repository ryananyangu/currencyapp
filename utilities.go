package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// RequestInput function to read user input from terminal using bufio
func RequestInput(prompt string) string {
	fmt.Printf("\n%s\n> ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func DecodingResponse(data string) (map[string]interface{}, error) {
	var response map[string]interface{}
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
