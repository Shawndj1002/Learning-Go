package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	// prompt
	fmt.Print("Enter a name: ")
	fmt.Scan(&name)

	fmt.Print("Enter an address: ")
	fmt.Scan(&address)

	aMap := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonObj, err := json.Marshal(aMap)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println(string(jsonObj))

}
