package main

import (
	"fmt"
	"time"
)

func preAllocExample() {
	n := 100000
	notAllocated := []int32{}
	preAllocated := make([]int32, 0, n)
	fmt.Printf("Total time without pre-allocation: %v\n", timeLoop(notAllocated, n))
	fmt.Printf("Total time with pre-allocation: %v\n", timeLoop(preAllocated, n)) // with pre-allocation, the time is much shorter

}

func timeLoop(slice []int32, n int) time.Duration {
	start := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(start)
}
