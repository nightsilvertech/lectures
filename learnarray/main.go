package main

import "fmt"

func main() {
	emails := []string{
		"a@gmail.com", // 0
		"b@gmail.com", // 1
		"c@gmail.com", // 2
		"d@gmail.com", // 3
		"e@gmail.com", // 4
	}

	fmt.Println(emails[2:])
}
