package coffee

import (
	"testing"
)

func TestGiveMeSomeCoffee(t *testing.T) {
	tests := []struct {
		name string
		in   Inputs
		out  error
	}{
		{
			name: "espresso",
			in:   Inputs{
				Name:       "Dennis Ritchie",
				CoffeeType: "espresso",
				Delivery:   true,
				NoDelay: true,
			},
			out:  nil,
		},
		{
			name: "cappuccino",
			in:   Inputs{
				Name:       "Dennis Ritchie",
				CoffeeType: "cappuccino",
				Delivery:   false,
				NoDelay: true,
			},
			out:  nil,
		},
		{
			name: "macchiato",
			in:   Inputs{
				Name:       "Dennis Ritchie",
				CoffeeType: "macchiato",
				Delivery:   false,
				NoDelay: true,
			},
			out:  nil,
		},
		{
			name: "latte",
			in:   Inputs{
				Name:       "Dennis Ritchie",
				CoffeeType: "latte",
				Delivery:   false,
				NoDelay: true,
			},
			out:  nil,
		},
		{
			name: "name is required",
			in:   Inputs{},
			out:  ErrNameIsRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := tt.in
			out := tt.out

			got := GiveMeSomeCoffee(in)
			if got != out {
				t.Errorf("GiveMeSomeCoffee(%s) got %v, want %v", tt.name, got, out)
			}
		})
	}

}
