package main

import (
	"fmt"
)

func main() {
	intSlice := []int{1, 2, 3}
	fmt.Println(sumIntSlice(intSlice))

	float32Slice := []float32{1.1, 2.2, 3.3}
	fmt.Println(sumFloat32Slice(float32Slice))

	float64Slice := []float64{1.1, 2.2, 3.3}
	fmt.Println(sumFloat64Slice(float64Slice))
}

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
