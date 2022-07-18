package main

import "fmt"

type Sapiens interface {
	Walk()
	Sleep()
	Eat()
}

type Human interface {
	Sapiens
	Speak()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Walk() {
	fmt.Println(p.Name, "is walking")
}
func (p Person) Sleep() {
	fmt.Println(p.Name, "is sleeping")
}
func (p Person) Eat() {
	fmt.Println(p.Name, "is eating")
}
func (p Person) Speak() {
	fmt.Println(p.Name, "is speaking")
}

func main() {
	person1 := Person{
		Name: "Shagy",
		Age:  50,
	}

	var sapiens1 Sapiens = person1
	sapiens1.Sleep()

	var human1 Human = person1
	human1.Speak()
}
