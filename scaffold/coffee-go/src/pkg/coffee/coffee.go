package coffee

import (
	"log"
	"time"
)

type Inputs struct {
	Name string
	CoffeeType string
	Delivery bool
}

func GiveMeSomeCoffee(inputs Inputs) {
	log.Printf("Preparing your coffee %v .....\n", inputs.Name)
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	if inputs.Delivery {
		log.Printf("Your %v coffee is ready, enjoy your trip\n", inputs.CoffeeType)
	} else {
		log.Printf("Your %v coffee is ready, have a seat and enjoy your drink\n", inputs.CoffeeType)
	}
}
