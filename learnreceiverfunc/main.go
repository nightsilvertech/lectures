package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) Greetings() {
	fmt.Println("Hallo nama saya", p.FirstName, p.LastName, "umur saya", p.Age)
}

func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person) Modify() {
	fmt.Println("Hallo nama saya", p.FirstName, p.LastName, "umur saya", p.Age)
}

func main() {
	// initialization
	person1 := Person{
		FirstName: "Udin",
		LastName:  "Ganteng",
		Age:       28,
	}

	// memanggil receiver function dari person1 yang type nya Person
	person1.Greetings()

	// mengubah FirstName person1 menjadi Mark lewat receiver function dengan pointer
	person1.SetFirstName("Mark")
	person1.Greetings()
}
