/**
Enter a floating point number and prints the integer which is
a truncated version of the floating point number that was entered.
*/
package main

import "fmt"

func Trunc(value float32) int {
	return int(value)
}

func main() {
	var floatNumber float32

	fmt.Printf("Enter a floating number: ")
	fmt.Scan(&floatNumber)

	fmt.Printf("Truncated integer: %d\n", Trunc(floatNumber))
}
