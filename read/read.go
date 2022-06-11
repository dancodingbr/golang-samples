/*
The program reads information from a file and represents it in a slice of structs.
Each line of the text file has a first name and a last name,
in that order, separated by a single space on the line.

Sample content of a text file:

John Rambo
Francis Harris
Bob Johnson
Othonospagonisdoulopoulos Konstantinomavrospolites
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAX_STRING_SIZE = 20

type Person struct {
	fname []rune
	lname []rune
}

func TruncFieldValue(strIn string) []rune {
	var strOut string
	for key, value := range strIn {
		if key < MAX_STRING_SIZE {
			strOut = strOut + string(value)
		} else {
			break
		}
	}
	return []rune(strOut)
}

func main() {

	var filename string

	fmt.Printf("Type the name of the text file (including extension): ")
	_, err := fmt.Scan(&filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	var lines []Person
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fullNames := strings.Split(scanner.Text(), " ")

		person := Person{
			fname: TruncFieldValue(fullNames[0]),
			lname: TruncFieldValue(fullNames[1]),
		}
		lines = append(lines, person)
	}

	for _, value := range lines {
		fmt.Printf("%s %s\n", string(value.fname), string(value.lname))
	}

}
