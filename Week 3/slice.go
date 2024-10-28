package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	for {
		var input string

		fmt.Print("Please Enter an int to add to the slice (Enter 'X' to quit): ")
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Not a valid Integer. Please enter an Integer or 'X' ")
			continue
		}

		slice = append(slice, num)

		sort.Ints(slice)

		fmt.Println("Sorted new slice: ", slice)
	}
}
