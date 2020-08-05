package main

import (
	"os"
	"vpc/pkg/vpc"
)

func main() {
	vpc.Run(loadInputs())
}

func loadInputs() vpc.Inputs {
	region := os.Getenv("REGION")
	vpcName := os.Getenv("VPC_NAME")
	vpcCIDR := os.Getenv("VPC_CIDR")
	vpcAZS := os.Getenv("VPC_AZS")
	customerName := os.Getenv("CUSTOMER_NAME")
	PWD := os.Getenv("CURRENT_PWD")

	return vpc.Inputs{
		Region:       region,
		VPCName:      vpcName,
		VPCCIDR:      vpcCIDR,
		VPCAZS:       vpcAZS,
		CustomerName: customerName,
		PWD:          PWD,
	}
}
