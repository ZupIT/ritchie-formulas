terraform {
	required_version = "0.12.13"
	required_providers {
		kubernetes = "~> 1.11.0"
		local      = "1.4.0"
		template   = "2.1.2"
		helm       = "0.10.4"
		external   = "1.2.0"
		tls        = "2.1.1"
		archive    = "1.3.0"
		random     = "2.2.1"
	}
}