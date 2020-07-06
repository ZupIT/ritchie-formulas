package tpl

const (
	DnsZoneMaintf = `
resource "aws_route53_zone" "dns" {
	name = var.domain_name
}	
	`

	DnsZoneOutputstf = `
output "zone_id" {
	description = "The dns zone id"
	value       = aws_route53_zone.dns.zone_id
}	
	`

	DnsZoneVariablestf = `
variable "domain_name" {
	default = ""
}	
	`
)
