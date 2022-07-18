package main

import "fmt"

func main() {
	// key value secara berpasangan
	var murid map[int]string
	murid = map[int]string{
		1: "Udin",
		2: "Jaki",
		3: "Akbar",
		4: "Jaenap",
	}
	jaenap := murid[4]
	fmt.Println(jaenap)
}
