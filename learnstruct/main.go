package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// initialization
	person1 := Person{
		FirstName: "Udin",
		LastName:  "Ganteng",
		Age:       28,
	}

	fmt.Println("Before Modification", person1)

	// modification
	person1.FirstName = "Mamung"
	person1.LastName = "Bandel"
	person1.Age = 20

	fmt.Println("After Modification", person1)
}
