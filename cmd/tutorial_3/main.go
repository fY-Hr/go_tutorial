package main

import (
	"fmt"
	"errors"
)

// this file is talking about function, destructuring, and error data type

func main() {
	printValue := "Hello World"
	printMe(printValue)

	numerator := 11
	denominator := 0

	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %v, Remainder: %v\n", result, remainder)
}

func printMe(s string) {
	fmt.Println(s)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error // the default value of err is nil, so if the err is nil, then there is no error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err //<-- because the type of the function expect 2 values of int to be returned
}
