package main

import (
	"fmt"
	"sync" // this is the package that we need to import to use go routines
	"time"
)

// this is the frist example of goroutine usage

var mutex = sync.Mutex{} // i'll place the mutex explanation in the end of the file
// note that if you use a go routines, you need to use waitgroup to wait for all the go routines to finish
// its kinda like async await in javascript
var wg = sync.WaitGroup{}

// We use *[]string so dbCall can update the same slice variable from main.
//
// If we passed []string instead, dbCall would receive a copy of the slice header
// (pointer, length, and capacity). Then append(...) would update that local copy
// of the slice header, not the original result slice variable.
func exampleOne(data []string, receptacle *[]string) {
	t0 := time.Now()
	for i := 0; i < len(data); i++ {
		wg.Add(1) // "im about to start 1 goroutine that you need to wait for"
		go dbCall(i, data, receptacle) // by using go, we can run a function in a separate thread
	} 
	// to prevent the main thread from exiting before all the go routines finish we use wg.Wait()
	wg.Wait() // "hey wait for the goroutine counter hit 0"
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0)) // and when the counter hit 0, it prints the total execution time
	fmt.Printf("The result of the db call is: %v\n", *receptacle)
}

func dbCall(i int, data []string, receptacle *[]string) {
	delay := 2000
	mutex.Lock() // lock the result slice to make sure that only one go routine can access the result slice at the same time
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is: ", data[i])
	*receptacle = append(*receptacle, data[i])
	mutex.Unlock() // unlock the mutex to enable other go routines to access the result slice
	wg.Done() // to tell the waitgroup that the go routine is done "hey, the goroutine is done"
}


/*
	the mutex is used to make sure that only one go routine can access the result slice at the same time
	if you don't use the mutex, the result slice will be changed by multiple go routines at the same time:
 
		result = append(result, dbData[i]) // on the first state, the result slice is readed as empty

	because the go routine is running concurrently, all go routines is reading the result slice as empty at the same time,
	and the result slice is changed by all the go routines at the same time without synchronization, they create a 
	data race condition

	by using the mutex, we can lock the result slice to make sure that only one go routine can access the result slice at the
	same time, and the result slice will be changed by only one go routine at a time

*/
