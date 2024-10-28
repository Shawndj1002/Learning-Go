// Coded by Shawn Marcial de Jesus
// Module 2
package main

import "fmt"

func main() {
	var fnum float64

	fmt.Printf("Please enter a floating point number: ")

	fmt.Scan(&fnum)

	trunc := int(fnum)

	fmt.Printf("Truncated Number (Integer): %d", trunc)
}
