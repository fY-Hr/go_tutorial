package main

import (
	"fmt"
	"time"
	"math/rand"
)

// channel is a way to send and receive data between goroutines
// to use channel. you need to use goroutines. if you don't use goroutines, you can't use channel

var MAX_LAPTOP_PRICE float32 = 979

func main() {
	laptopChannel := make(chan string, 10)
	websites := []string{"amazon", "bestbuy", "tokopedia", "shopee", "lazada"}
	for i := range websites {
		// if you dont use goroutines, the error will be a deadlock error because there is no goroutine that is running
		go getLaptopPrice(websites[i], laptopChannel)
	}
	sendMessage(laptopChannel)
}

func getLaptopPrice(website string, laptopChannel chan string) {
	for {
		time.Sleep(time.Second*1)
		laptopPrice := rand.Float32() * 20
		if laptopPrice < MAX_LAPTOP_PRICE {
			laptopChannel <- website
			break
		}
	}
}

func sendMessage(laptopChannel chan string) {
	fmt.Println("The deal is on website: ", <- laptopChannel)
}

/*
	the problem with the code above is that the laptopChannel is not buffered, so if the getLaptopPrice function is slow,
	it will block the main thread and the program will be stuck

	to fix this, we can use a buffered channel, so the main thread will not be blocked by the getLaptopPrice function
	and the program will not be stuck

	we can use the following code to fix the problem:

		laptopChannel := make(chan string, 10)

	this will create a buffered channel with a buffer size of 10

	if we use a buffered channel, the main thread will not be blocked by the getLaptopPrice function

*/
