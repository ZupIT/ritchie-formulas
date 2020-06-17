package main

import (
	"eks/pkg/eks"
	"os"
)

func main() {
	eks.Run(loadInputs())
}

func loadInputs() eks.Inputs {
	clusterName := os.Getenv("CLUSTER_NAME")
	domainName := os.Getenv("DOMAIN_NAME")
	PWD := os.Getenv("CURRENT_PWD")

	return eks.Inputs{
		ClusterName: clusterName,
		DomainName:  domainName,
		PWD:         PWD,
	}
}
