package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// rune
	myString := []rune("hello")
	indexed := myString[1]

	fmt.Printf("%v, %T\n", indexed, indexed)
	// fyi, rune is an alias of int32

	var myRune rune = 'Y'
	fmt.Printf("%v, %T\n", myRune, myRune)
	// rune is int32 because the char is a unicode character
	// you can search about utf-8 encoding in the internet, and see how it works

	for key, value := range myString {
		fmt.Printf("%v, %v\n", key, value)
	}

	// string
	strSlice := []string{"w", "o", "r", "l", "d"}
	str, duration := concatStrWithoutBuilder(strSlice)
	fmt.Printf("%v, Total time without builder: %v\n", str, duration)
	str2, duration2 := concatStrWithBuilder(strSlice)
	fmt.Printf("%v, Total time with builder: %v\n", str2, duration2)

}

// as you can see here, the total time without builder is longer than the total time with builder
// this is because the strings.Join function is using the loop to concat the string
// and the builder is using the WriteString function to concat the string

// so the builder is faster than the strings.Join function

func concatStrWithoutBuilder(strSlice []string) (string, time.Duration) {
	start := time.Now()
	str := strings.Join(strSlice, "")
	return str, time.Since(start)
}

func concatStrWithBuilder(strSlice []string) (string, time.Duration) {
	start := time.Now()
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	return strBuilder.String(), time.Since(start)
}
