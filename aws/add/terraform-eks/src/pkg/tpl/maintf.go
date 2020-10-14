package tpl

const (
	Maintf = `
# --------------------------------------

data "aws_eks_cluster" "cluster" {
	name = module.kubernetes_cluster.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
	name = module.kubernetes_cluster.cluster_id
}

provider "kubernetes" {
	host                   = data.aws_eks_cluster.cluster.endpoint
	cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
	token                  = data.aws_eks_cluster_auth.cluster.token
	load_config_file       = false
	version                = "~> 1.9"
}

variable "kubernetes_cluster_name" {
	default = ""
}

module "kubernetes_cluster" {
	source          = "terraform-aws-modules/eks/aws"
	cluster_name    = var.kubernetes_cluster_name
	cluster_version = "1.16"
	subnets         = module.vpc.private_subnets
	vpc_id          = module.vpc.vpc_id

	worker_groups = [
		{
			instance_type = "t2.small"
			asg_max_size  = 5
		}
	]
}

# --------------------------------------- dns zone to expose your applications
variable "domain_name" {
	default = ""
}

module "dns" {
	source = "./modules/dns_zone"
	domain_name = var.domain_name
}

# --------------------------------------- iam to do things on k8s
module "iam_k8s" {
	source = "./modules/iam_k8s"

	kubernetes_cluster_name = var.kubernetes_cluster_name
	kubernetes_worker_iam_role_name = module.kubernetes_cluster.worker_iam_role_name

}

# --------------------------------------- helm
provider "helm" {
	kubernetes {
		host                   = data.aws_eks_cluster.cluster.endpoint
		cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
		token                  = data.aws_eks_cluster_auth.cluster.token
		load_config_file       = false
	}

	debug           = true
}

# --------------------------------------- helm repositories
module "helm_deps" {

	source = "./modules/helm_deps"
	kubernetes_cluster = module.kubernetes_cluster
	kubernetes_cluster_name = var.kubernetes_cluster_name
	region = var.region
	dns_zone_id = module.dns.zone_id
	vpc_id = module.vpc.vpc_id

}
# -------------------------------- helm test exposure

variable "namespace" {
	default = ""
}
`
)
