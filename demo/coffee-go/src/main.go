package main

import (
	"coffee/pkg/coffee"
	"os"
	"strconv"
)

func main() {
	coffee.GiveMeSomeCoffee(loadInputs())
}

func loadInputs() coffee.Inputs {
	delivery, _ := strconv.ParseBool(os.Getenv("DELIVERY"))
	name := os.Getenv("NAME")
	coffeeType := os.Getenv("COFFEE_TYPE")
	return coffee.Inputs{
		Name:       name,
		CoffeeType: coffeeType,
		Delivery:   delivery,
	}
}
