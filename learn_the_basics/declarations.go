package main

import "fmt"

func main() {
	// Type is specified using keyword "string"
	var student1 string = "John"

	// Type is inferred
	var student2 = "Jane"
	x := 2

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
