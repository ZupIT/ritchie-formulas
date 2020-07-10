package main

import (
	"github.com/ZupIT/ritchie-formulas/scaffold/spring-starter/src/pkg/dependencies/list"
)

func main() {
	loadInputs().Run()
}

func loadInputs() CommandHandler {
	return list.Inputs{}
}

type CommandHandler interface {
	Run()
}
