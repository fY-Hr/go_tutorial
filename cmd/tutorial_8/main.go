package main

var dbData []string = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}
var result []string = []string{} // for your information, you can just only use := inside a function

func main() {
	// if you want to run one of these examples, you can uncomment the one you want to run
	// don't forget to comment the other one.
	// exampleOne(dbData, &result) // we are passing the memory address so the exampleOne will have the same refference of result
	exampleTwo(dbData, &result)
}



