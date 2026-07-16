package main

import (
	"fmt"
)

// Generics
// Generics are a way to write functions that can work with multiple types
// it allows a function to have additional parameters that can be used to specify the type of the parameters

func main() {
	intSlice := []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	float32Slice := []float32{1.1, 2.2, 3.3}
	fmt.Println(sumSlice[float32](float32Slice))

	float64Slice := []float64{1.1, 2.2, 3.3}
	fmt.Println(sumSlice[float64](float64Slice))
}

// if we are using Generics, we can write a function that can work with any type
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// if we are not using Generics, we need to write the same function for each type:
/*
	func sumIntSlice(intSlice []int) int {
		var sum int
		for _, v := range intSlice {
			sum += v
		}
		return sum
	}

	func sumFloat32Slice(float32Slice []float32) float32 {
		var sum float32
		for _, v := range float32Slice {
			sum += v
		}
		return sum
	}

	func sumFloat64Slice(float64Slice []float64) float64 {
		var sum float64
		for _, v := range float64Slice {
			sum += v
		}
		return sum
	}
*/
