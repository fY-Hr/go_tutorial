package main

import "fmt"

// this file is about const and variable and data type
func main(){
	var intNum uint16 = 10101
	fmt.Println(intNum)

	var floatVar float32 = 123112.9
	fmt.Println(floatVar)

	// arithmetic

	var intNum32 int32 = 3
	var floatNum32 float32 = 3.14
	var result1 float32 = float32(intNum32) + floatNum32

	fmt.Println(result1)

	var intNum1 int = 3
	var intNum2 int = 5
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	// string

	var myString string = "Hellow" + "World"
	fmt.Println(myString)

	// length check
	fmt.Println(len(myString));

	// rune

	var myRune rune = 'a'
	fmt.Println(myRune) // the result is 97 because 'a' is the 97th character in the ASCII table

	// default values
	var intDefault int
	fmt.Println(intDefault) // the result is 0
	var floatDefault float32
	fmt.Println(floatDefault) // the result is 0
	var stringDefault string
	fmt.Println(stringDefault) // the result is ""
	var boolDefault bool
	fmt.Println(boolDefault) // the result is false
	var byteDefault byte
	fmt.Println(byteDefault) // the result is 0
	var runeDefault rune
	fmt.Println(runeDefault)

	// shorthand
	shortInt := 10
	fmt.Println(shortInt) // the result is 10

	// const
	const constVar = "You can't change this"
	fmt.Println(constVar)
	// constVar = "You can change this" // this will cause an error
}
