package main

import (
	"kubernetes/core/pkg/health"
	"os"
)

func main() {
	loadInputs().Run()
}

func loadInputs() CommandHandler {
	return health.Inputs{
		Namespace:   os.Getenv("NAMESPACE"),
		PodPartName: os.Getenv("POD_PART_NAME"),
		Kubeconfig:  os.Getenv("KUBECONFIG"),
	}
}

type CommandHandler interface {
	Run()
}
