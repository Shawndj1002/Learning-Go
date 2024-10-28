package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputStr string

	fmt.Printf("Enter a string:")
	fmt.Scan(&inputStr)

	inputStr = strings.ToLower(inputStr)

	startsWithI := strings.HasPrefix(inputStr, "i")
	containsA := strings.Contains(inputStr, "a")
	endsWithN := strings.HasSuffix(inputStr, "n")

	if startsWithI && containsA && endsWithN {
		fmt.Printf("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
