package main

import (
	"fmt"
)

func main() {
	// array
	intArr := [3]int32{1, 2, 3} //<-- the [] is used as the type so the value of the array stored inside {}
	fmt.Println(intArr)

	fmt.Printf("%T\n", intArr)

	// slice (slice is like an array but the size is dynamic. It is build on top of the array)
	intSlice := []int32{4, 5, 6} //<-- slice has no fixed size
	fmt.Printf("%T\n", intSlice)
	fmt.Println(intSlice)
	fmt.Printf("The length of the slice is %v, and the capacity of the slice is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) // append new value into the slice
	fmt.Printf("The length of the slice is %v, and the capacity of the slice is %v\n", len(intSlice), cap(intSlice))

	// so the slice is dynamic and the capacity is the size of the array
	// when the slice is full, the capacity will be doubled

	// you can also append a slice into the slice
	intSlice = append(intSlice, []int32{8, 9, 10}...)
	// but you need to use spread operator to unpack the 8, 9, 10 as a value and then append it into the slice
	fmt.Printf("The length of the slice is %v, and the capacity of the slice is %v\n", len(intSlice), cap(intSlice))

	// there is another way to make a slice
	intSlice2 := make([]int32, 3, 8)
	// by using make function, we can make a slice with the length of 0 and the capacity of 5
	fmt.Printf("The length of the slice is %v, and the capacity of the slice is %v\n", len(intSlice2), cap(intSlice2))
	fmt.Println(intSlice2) // and because there is no value in the slice, the default value is 0

	// map
	myMap := make(map[string]int32)
	fmt.Println(myMap) // this is how you declare a map without any value

	// you can also declare a map with value immediately
	myMap2 := map[string]uint8{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	// the structure of the map:
	// map[keyType]valueType{key: value, ...}

	fmt.Println(myMap2["a"])
	fmt.Println(myMap2["b"])

	// you can also change the value of the map
	myMap2["a"] = 8
	fmt.Println(myMap2["a"])

	// becareful with using map in go because even if there is no key inside the map, the default value is 0
	// so you need to check if the key is inside the map before using it
	// by destructuring the map, you can check if the key is there or not:
	check, ok := myMap2["g"]
	fmt.Println(check, ok) // the result will be 0, and false
	// if the key is there, the result will be the value of the key, and true
	check, ok = myMap2["a"]
	fmt.Println(check, ok) // the result will be 1, and true

	// to delete a key from the map, you can use delet function
	delete(myMap2, "a")
	fmt.Println(myMap2)

	// for iteration
	// this is a basic usage of for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// and the fact is there is no other way to do iteration in go except for for loop
	// you can use range keyword to iterate over the map
	// for key, value := range myMap2 {
	// 	fmt.Println(key, value)
	// }

	// if you want to just iterate the value of the map, you can use range keyword
	for _, value := range myMap2 {
		fmt.Println(value)
	}

	// another example:
	anotherMap := map[string]string{
		"name":   "antei",
		"age":    "20",
		"gender": "male",
	}
	// if you want to just iterate the key of the map, you can use range keyword
	for name := range anotherMap {
		// the name here is getting the key of the map. Even if the variable name is changed
		// the value that will be inserted into the name variable is still the key
		// that's why people commonly use key and value name.
		fmt.Println(name)
	}

	// you can also loop through array and slice

	// even if there is no other loop in go,
	// you can still type for loop in a similiar way with the while loop style
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// pre allocation example
	preAllocExample()
}
