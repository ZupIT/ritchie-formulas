package main

import (
	"kubernetes/core/pkg/forward"
	"os"
)

func main() {
	loadInputs().Run()
}

func loadInputs() CommandHandler {
	return forward.Inputs{
		Namespace:   os.Getenv("NAMESPACE"),
		PodPartName: os.Getenv("POD_PART_NAME"),
		Kubeconfig:  os.Getenv("KUBECONFIG"),
		PortMap:	 os.Getenv("PORT"),
	}
}

type CommandHandler interface {
	Run()
}