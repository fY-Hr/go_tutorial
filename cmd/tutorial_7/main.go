package main

import (
	"fmt"
)

// this file is talking bout pointer
// pointer explanation is a bit long, you can search it on internet instead

func main() {
	var p *int32   // pointer to int32 (the default value is nil)
	p = new(int32) // allocate memory for int32 (p here is just storing the location address of the real value)
	var i int32    // normal int32 variable
	fmt.Printf("The value of p is %v\n", p)
	fmt.Printf("The value of i is %v\n", i)
	// you can assign an address of a value to the memory that p is pointing to:
	p = &i // assign the address of i to p (now the address is no more the new(int32))
	fmt.Printf("The value of p is %v\n", p)
	*p = 100                                            // we assign the value of i. because *p is targetting the memory address of i
	fmt.Println(p)                                      // print the address of the memory
	fmt.Println("This is *p:", *p)                      // print the value of the memory
	fmt.Println("This is i:", i)                        // print the value of i
	fmt.Println("This is the memory address of i:", &i) // this is what *p pointing to
	// (because the mem address of the p is the address of i, then if we change the pointer of p (line 18), then the value of the mem address (which is i) will be changed too)
	// and likewise too:
	i = 200
	fmt.Println("This is i:", i)
	fmt.Println("This is *p:", *p)

	// there is a behaviour in go where a variable is referring into a variable, the new variable will make a duplicate with a different memory location
	testDupl := 2
	testDupl2 := testDupl

	fmt.Println("This is testDupl:", testDupl)
	fmt.Println("This is testDupl2:", testDupl2)
	fmt.Println("This is the memory address of testDupl:", &testDupl)
	fmt.Println("This is the memory address of testDupl2:", &testDupl2) // the memory address of testDupl2 is different from testDupl

	testDupl2 = testDupl2 * 100
	fmt.Println("This is testDupl2:", testDupl2)
	fmt.Println("This is testDupl:", testDupl) // the value of testDupl is not changed

	// but for slice it is different:
	testSlice := []int32{1, 2, 3}
	testSlice2 := testSlice
	fmt.Println("This is testSlice:", testSlice)
	fmt.Println("This is testSlice2:", testSlice2)

	testSlice2[2] = 100
	fmt.Println("This is testSlice2:", testSlice2)
	fmt.Println("This is testSlice:", testSlice) // the value of testSlice is changed

	// another example with array
	testArray := [3]int32{1, 2, 3}
	testArray2 := testArray

	testArray2[1] = 100
	fmt.Println("This is testArray2:", testArray2)
	fmt.Println("This is testArray:", testArray) // the value of testArray is not changed

	// if you want the array to follow the value and not duplicating it:
	var testArray3 *[3]int32 = &testArray // the & is to get the memory address of testArray, so the testArray3 is pointing into the memory address of testArray and get the value

	testArray3[0] = 100
	fmt.Println("This is testArray3:", testArray3)
	fmt.Println("This is testArray:", testArray) // the value of testArray is changed

	// summary:
	// * is used to point a memory address
	// & is used to get the memory address
}
