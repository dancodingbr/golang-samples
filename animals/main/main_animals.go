package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/animals"
)

func main() {

	animals.LoadDataSource()

	for {
		fmt.Printf("> ")
		in := bufio.NewReader(os.Stdin)
		request, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		reqVars := strings.Split(strings.Trim(request, "\r\n"), " ")

		if len(reqVars) != 2 {
			fmt.Println("Invalid request.")
			continue
		}

		answer := animals.RequestDataAbout(reqVars[0], reqVars[1])

		fmt.Printf("Answer: %s\n", answer)
	}

}
