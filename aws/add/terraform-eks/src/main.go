package main

import (
	"eks/pkg/eks"
)

func main() {
	eks.Run(loadInputs())
}

func loadInputs() eks.Inputs {
	clusterName := "name"                         //os.Getenv("CLUSTER_NAME")
	domainName := "domain.io"                     //os.Getenv("DOMAIN_NAME")
	PWD := "/Users/guillaumefalourd/test-formula" //os.Getenv("CURRENT_PWD")

	return eks.Inputs{
		ClusterName: clusterName,
		DomainName:  domainName,
		PWD:         PWD,
	}
}
