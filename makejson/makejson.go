/*
The program prompts the user to first enter a name and an address.
Then creates and prints a JSON object from the map.
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inName, inAddress string

	fmt.Printf("Type a name: ")
	in := bufio.NewReader(os.Stdin)
	inName, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Printf("Type an address: ")
	in = bufio.NewReader(os.Stdin)
	inAddress, err = in.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	mapObject := map[string]string{
		"name":    strings.Trim(inName, "\r\n"),
		"address": strings.Trim(inAddress, "\r\n"),
	}

	jsonObject, err := json.Marshal(mapObject)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Printf("\n%s\n", jsonObject)
}
