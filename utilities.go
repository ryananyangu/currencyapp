package main

import (
	"bufio"
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
