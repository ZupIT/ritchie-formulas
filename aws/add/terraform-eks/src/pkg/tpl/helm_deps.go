package tpl

const (
	HelmMaintf = `
data "helm_repository" "stable" {
  name = "stable"
  url  = "https://kubernetes-charts.storage.googleapis.com"
}

data "helm_repository" "incubator" {
  name = "incubator"
  url  = "http://storage.googleapis.com/kubernetes-charts-incubator"
}

data "helm_repository" "codecentric" {
  name = "codecentric"
  url  = "https://codecentric.github.io/helm-charts"
}

data "helm_repository" "loki" {
  name = "loki"
  url  = "https://grafana.github.io/loki/charts"
}

# ------------------------------------------- kubernetes service accounts required

resource "kubernetes_service_account" "tiller" {
  metadata {
    name      = "tiller"
    namespace = "kube-system"
  }
  depends_on = [
    var.kubernetes_cluster,
  ]
}

resource "kubernetes_cluster_role_binding" "tiller" {
  metadata {
    name = "tiller"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster-admin"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "tiller"
    namespace = "kube-system"
  }

  depends_on = [
    kubernetes_service_account.tiller
  ]
}

# -------------------------------------------- misc things required to run, expose and auto scale things on k8s



resource "helm_release" "cluster-autoscaler" {
  name       = "cluster-autoscaler"
  repository = "stable"
  chart      = "stable/cluster-autoscaler"
  version    = "6.0.0"
  namespace  = "cluster-autoscaler"

  set {
    name  = "autoDiscovery.clusterName"
    value = var.kubernetes_cluster_name
  }

  set {
    name  = "cloud-provider"
    value = "aws"
  }

  set {
    name  = "awsRegion"
    value = var.region
  }

  set {
    name  = "rbac.create"
    value = true
  }

  depends_on = [
    kubernetes_cluster_role_binding.tiller,
    kubernetes_service_account.tiller
  ]
}


resource "helm_release" "external-dns" {
  name       = "external-dns"
  repository = "stable"
  chart      = "stable/external-dns"
  version    = "2.9.4"
  namespace  = "external-dns"

  set {
    name  = "provider"
    value = "aws"
  }

  set {
    name  = "aws.zoneType"
    value = "public"
  }

  set {
    name  = "txtOwnerId"
    value = var.dns_zone_id
  }

  set {
    name  = "rbac.create"
    value = true
  }

  set {
    name  = "policy"
    value = "sync"
  }

  depends_on = [
    kubernetes_cluster_role_binding.tiller,
    kubernetes_service_account.tiller,
    helm_release.aws-alb-ingress-controller
  ]
}


resource "helm_release" "aws-alb-ingress-controller" {
  name       = "aws-alb-ingress-controller"
  repository = "incubator"
  chart      = "incubator/aws-alb-ingress-controller"
  version    = "1.0.0"
  namespace  = "aws-alb-ingress-controller"

  set {
    name  = "clusterName"
    value = var.kubernetes_cluster_name
  }

  set {
    name  = "awsRegion"
    value = var.region
  }

  set {
    name  = "awsVpcID"
    value = var.vpc_id
  }

  depends_on = [
    kubernetes_cluster_role_binding.tiller,
    kubernetes_service_account.tiller
  ]
}

	`

	HelmVariablestf = `
variable "kubernetes_cluster" {
  default = ""
}

variable "kubernetes_cluster_name" {
  default = ""
}

variable "region" {
  default = ""
}

variable "dns_zone_id" {
  default = ""
}

variable "vpc_id" {
  default = ""
}
	`
)
