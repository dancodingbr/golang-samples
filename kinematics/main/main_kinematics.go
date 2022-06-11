package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/kinematics"
)

func main() {

	var strValue string
	var acceleration, initial_velocity, initial_displacement, time float64

	fmt.Printf("Type an acceleration value: ")
	_, err := fmt.Scan(&strValue)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	acceleration, err = strconv.ParseFloat(strValue, 64)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Printf("Type an initial velocity value: ")
	_, err = fmt.Scan(&strValue)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	initial_velocity, err = strconv.ParseFloat(strValue, 64)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Printf("Type an initial displacement value: ")
	_, err = fmt.Scan(&strValue)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	initial_displacement, err = strconv.ParseFloat(strValue, 64)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Printf("Type a time value: ")
	_, err = fmt.Scan(&strValue)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	time, err = strconv.ParseFloat(strValue, 64)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fn := kinematics.GenDisplaceFn(acceleration, initial_velocity, initial_displacement)
	fmt.Printf("Displacement after 3 seconds: %f\n", fn(time))

}
