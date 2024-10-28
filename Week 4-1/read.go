package main

import (
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter the name of the file: ")
	var fileName string
	fmt.Scanln(&fileName)

	// Open the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	// Create a slice to hold Name structs
	var names []Name

	// Process each line
	for _, line := range lines {
		line = strings.TrimSpace(line) // Remove any extra whitespace
		if line == "" {
			continue // Skip empty lines
		}
		parts := strings.Fields(line)
		if len(parts) == 2 {
			names = append(names, Name{fname: parts[0], lname: parts[1]})
		}
	}

	// Print the names from the slice
	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.fname, name.lname)
	}

}
