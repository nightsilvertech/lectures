package main

import "fmt"

func main() {
	emails := []string{
		"a@gmail.com",
		"b@gmail.com",
		"c@gmail.com",
		"d@gmail.com",
	}
	for index, email := range emails {
		fmt.Println(index, email)
	}

	students := map[int]string{
		1: "Zack",
		2: "Zoe",
		3: "Meg",
		4: "Stewie",
	}
	for key, value := range students {
		fmt.Println(key, value)
	}
}
