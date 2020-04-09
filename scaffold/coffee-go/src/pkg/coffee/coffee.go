package coffee

import (
	"errors"
	"log"
	"time"
)

var   (
	ErrNameIsRequired = errors.New("the customer name is required")
)

type Inputs struct {
	Name string
	CoffeeType string
	Delivery bool
	NoDelay bool
}



func GiveMeSomeCoffee(i Inputs) error {
	if i.Name == "" {
		return ErrNameIsRequired
	}

	log.Printf("Preparing your coffee %v .....\n", i.Name)
	log.Println("......")
	i.delay()
	log.Println("......")
	i.delay()
	log.Println("......")
	i.delay()
	if i.Delivery {
		log.Printf("Your %v coffee is ready, enjoy your trip\n", i.CoffeeType)
	} else {
		log.Printf("Your %v coffee is ready, have a seat and enjoy your drink\n", i.CoffeeType)
	}

	return nil
}

func (i Inputs) delay() {
	if !i.NoDelay {
		time.Sleep(time.Second * 1)
	}
}
