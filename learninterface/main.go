package main

import "fmt"

type Primate interface {
	Sleep()
	Eat()
	Walk()
}

type Human interface {
	Primate
	Laugh()
	Speak()
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) Sleep() {
	fmt.Println(fmt.Sprintf("%s is sleeping", p.FirstName))
}

func (p Person) Eat() {
	fmt.Println(fmt.Sprintf("%s is eating", p.FirstName))
}

func (p Person) Walk() {
	fmt.Println(fmt.Sprintf("%s is walking", p.FirstName))
}

func (p Person) Laugh() {
	fmt.Println(fmt.Sprintf("%s is laughing", p.FirstName))
}

func (p Person) Speak() {
	fmt.Println(fmt.Sprintf("%s is speaking", p.FirstName))
}

func main() {
	person1 := Person{
		FirstName: "Udin",
		LastName:  "Ganteng",
		Age:       28,
	}

	var primate Primate = person1
	primate.Sleep()
	primate.Eat()
	primate.Walk()

	var human Human = person1
	human.Sleep()
	human.Eat()
	human.Walk()
	human.Speak()
	human.Laugh()
}
