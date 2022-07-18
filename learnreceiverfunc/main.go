package main

import "fmt"

type Car struct {
	Manufacture string
	Type        string
	Fuel        int
	CC          int
}

func (c *Car) FlushFuel() {
	c.Fuel = 0
}

func (c *Car) FillFuel() {
	c.Fuel = 1
}

func (c Car) Moving() {
	if c.Fuel == 0 {
		fmt.Println("Can't move fuel is empty")
	} else {
		fmt.Println("This car", c.Manufacture, "is moving")
	}
}

func main() {
	car1 := Car{
		Manufacture: "BMW",
		Type:        "Sport",
		Fuel:        1,
		CC:          500,
	}
	car1.FlushFuel()
	car1.FillFuel()
	car1.Moving()
}
