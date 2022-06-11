package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/bubblesort"
)

func main() {
	var sli []int
	var intValue int
	var strValue string

	sli = make([]int, 10)
	counter := 0

	for strValue != "X" && counter < len(sli) {
		fmt.Printf("Enter an integer value (negative or positive) or X to exit: ")
		_, err := fmt.Scan(&strValue)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		if strValue == "X" {
			os.Exit(0)
		}

		intValue, err = strconv.Atoi(strValue)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		sli[counter] = intValue
		counter++
	}

	fmt.Printf("Unsorted slice: %v\n", sli)
	bubblesort.BubbleSort(sli)
	fmt.Printf("Sorted slice: %v\n", sli)
}
