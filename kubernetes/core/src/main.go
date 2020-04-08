package main

import (
	"kubernetes/core/pkg/forward"
	"kubernetes/core/pkg/health"
	"kubernetes/core/pkg/logs"
	"log"
	"os"
)

const (
	logsCmd     = "logs"
	healthCmd     = "health"
	portForwardCmd   = "portForward"
)

func main() {
	loadInputs().Run()
}

func loadInputs() CommandHandler {
	command := os.Getenv("COMMAND")
	switch command {
	case logsCmd:
		return logs.Inputs{
			Namespace:   os.Getenv("NAMESPACE"),
			PodPartName: os.Getenv("POD_PART_NAME"),
			Kubeconfig:  os.Getenv("KUBECONFIG"),
		}
	case healthCmd:
		return health.Inputs{
			Namespace:   os.Getenv("NAMESPACE"),
			PodPartName: os.Getenv("POD_PART_NAME"),
			Kubeconfig:  os.Getenv("KUBECONFIG"),
		}
	case portForwardCmd:
		return forward.Inputs{
			Namespace:   os.Getenv("NAMESPACE"),
			PodPartName: os.Getenv("POD_PART_NAME"),
			Kubeconfig:  os.Getenv("KUBECONFIG"),
			PortMap:	 os.Getenv("PORT"),
		}
	default:
		log.Fatalln("Command not found")
	}
	return nil
}

type CommandHandler interface {
	Run()
}