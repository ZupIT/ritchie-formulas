// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	input1 := os.Getenv("RIT_INPUT_LOGIN")
	input2 := os.Getenv("RIT_INPUT_CREDENTIAL")
	input3 := os.Getenv("RIT_INPUT_TYPE")
	input4 := os.Getenv("RIT_INPUT_HOST")
	input5 := os.Getenv("RIT_INPUT_PORT")
	input6 := os.Getenv("RIT_INPUT_TEXT")
	input7 := os.Getenv("RIT_INPUT_PASSWORD")

	formula.Formula{
		Login:        input1,
		Credential:   input2,
		DatabaseType: input3,
		Hostname:     input4,
		Port:         input5,
		Username:     input6,
		Password:     input7,
	}.Run(os.Stdout)
}
