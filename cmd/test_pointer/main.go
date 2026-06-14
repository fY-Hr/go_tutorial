package main

import "fmt"

func main() {
	var address1 *int32 = new(int32)
	fmt.Println(&address1)
	fmt.Println(address1)

	fmt.Println(*address1)

	var address3 **int32 = &address1
	fmt.Println(address3)

	var address2 int32 = 10
	address1 = &address2 // berubah jadi address2
	fmt.Println(*address1)
	fmt.Println(address3)
	fmt.Println(&address2)
	fmt.Println(**address3)
}
