package tpl

const (
	Variable = `
region="{{.Region}}"
vpc_name="{{.VPCName}}"
vpc_cidr="{{.VPCCIDR}}"
vpc_azs=[{{.VPCAZS}}]
customer_name="{{.CustomerName}}"

`
)
