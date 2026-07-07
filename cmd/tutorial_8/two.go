package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var mutex = sync.RWMutex{}

func exampleTwo(data []string, receptacle *[]string) {
	t0 := time.Now()
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go dbCall(i, data, receptacle)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("The result of the db call is: %v\n", *receptacle)
}

func count() {
	var res int
	for i := 0; i<100000000; i++ {
		res++
	}
	wg.Done()
}

func dbCall(i int, data []string, receptacle *[]string) {
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(data[i], receptacle)
	log(*receptacle)
	wg.Done()
}

func save(result string, receptacle *[]string) {
	mutex.Lock()
	*receptacle = append(*receptacle, result)
	mutex.Unlock()
}

func log(results []string){
	mutex.RLock()
	fmt.Printf("\nThe current result are: %v", results)
	mutex.RUnlock()
}
