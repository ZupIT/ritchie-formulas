package tpl

const (
	IAMK8SMaintf = `
data "aws_iam_policy_document" "aws-assume-role-policy-eks" {
  version = "2012-10-17"
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["eks.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "aws-assume-role-policy-ec2" {
  version = "2012-10-17"
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "aws-cluster-auto-scaler-policy-document" {
  version = "2012-10-17"
  statement {
    effect = "Allow"
    actions = [
      "autoscaling:DescribeAutoScalingGroups",
      "autoscaling:DescribeAutoScalingInstances",
      "autoscaling:DescribeLaunchConfigurations",
      "autoscaling:DescribeTags",
      "autoscaling:SetDesiredCapacity",
      "autoscaling:TerminateInstanceInAutoScalingGroup"
    ]
    resources = ["*"]
  }
}

data "aws_iam_policy_document" "aws-ingress-policy-document" {
  version = "2012-10-17"
  statement {
    effect = "Allow"
    actions = [
      "acm:DescribeCertificate",
      "acm:ListCertificates",
      "acm:GetCertificate",
      "ec2:AuthorizeSecurityGroupIngress",
      "ec2:CreateSecurityGroup",
      "ec2:CreateTags",
      "ec2:DeleteTags",
      "ec2:DeleteSecurityGroup",
      "ec2:DescribeAccountAttributes",
      "ec2:DescribeAddresses",
      "ec2:DescribeInstances",
      "ec2:DescribeInstanceStatus",
      "ec2:DescribeInternetGateways",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeSubnets",
      "ec2:DescribeTags",
      "ec2:DescribeVpcs",
      "ec2:ModifyInstanceAttribute",
      "ec2:ModifyNetworkInterfaceAttribute",
      "ec2:RevokeSecurityGroupIngress",
      "elasticloadbalancing:AddListenerCertificates",
      "elasticloadbalancing:AddTags",
      "elasticloadbalancing:CreateListener",
      "elasticloadbalancing:CreateLoadBalancer",
      "elasticloadbalancing:CreateRule",
      "elasticloadbalancing:CreateTargetGroup",
      "elasticloadbalancing:DeleteListener",
      "elasticloadbalancing:DeleteLoadBalancer",
      "elasticloadbalancing:DeleteRule",
      "elasticloadbalancing:DeleteTargetGroup",
      "elasticloadbalancing:DeregisterTargets",
      "elasticloadbalancing:DescribeListenerCertificates",
      "elasticloadbalancing:DescribeListeners",
      "elasticloadbalancing:DescribeLoadBalancers",
      "elasticloadbalancing:DescribeLoadBalancerAttributes",
      "elasticloadbalancing:DescribeRules",
      "elasticloadbalancing:DescribeSSLPolicies",
      "elasticloadbalancing:DescribeTags",
      "elasticloadbalancing:DescribeTargetGroups",
      "elasticloadbalancing:DescribeTargetGroupAttributes",
      "elasticloadbalancing:DescribeTargetHealth",
      "elasticloadbalancing:ModifyListener",
      "elasticloadbalancing:ModifyLoadBalancerAttributes",
      "elasticloadbalancing:ModifyRule",
      "elasticloadbalancing:ModifyTargetGroup",
      "elasticloadbalancing:ModifyTargetGroupAttributes",
      "elasticloadbalancing:RegisterTargets",
      "elasticloadbalancing:RemoveListenerCertificates",
      "elasticloadbalancing:RemoveTags",
      "elasticloadbalancing:SetIpAddressType",
      "elasticloadbalancing:SetSecurityGroups",
      "elasticloadbalancing:SetSubnets",
      "elasticloadbalancing:SetWebACL",
      "iam:CreateServiceLinkedRole",
      "iam:GetServerCertificate",
      "iam:ListServerCertificates",
      "cognito-idp:DescribeUserPoolClient",
      "waf-regional:GetWebACLForResource",
      "waf-regional:GetWebACL",
      "waf-regional:AssociateWebACL",
      "waf-regional:DisassociateWebACL",
      "waf:GetWebACL",
      "tag:GetResources",
      "tag:TagResources",
    ]
    resources = ["*"]
  }
}

data "aws_iam_policy_document" "aws-external-dns-policy-document" {
  version = "2012-10-17"
  statement {
    effect = "Allow"
    actions = [
      "route53:ChangeResourceRecordSets",
      "route53:ListHostedZones",
      "route53:ListResourceRecordSets",
    ]
    resources = ["*"]
  }
}


resource "aws_iam_policy" "aws-external-dns-policy" {
  name   = "aws-external-dns-policy-${var.kubernetes_cluster_name}"
  policy = data.aws_iam_policy_document.aws-external-dns-policy-document.json
}

resource "aws_iam_role_policy_attachment" "aws-external-dns-attachment" {
  policy_arn = aws_iam_policy.aws-external-dns-policy.arn
  role      = var.kubernetes_worker_iam_role_name
}

resource "aws_iam_policy" "aws-ingress-policy" {
  name   = "aws-ingress-policy-${var.kubernetes_cluster_name}"
  policy = data.aws_iam_policy_document.aws-ingress-policy-document.json
}

resource "aws_iam_role_policy_attachment" "aws-ingress-policy-attachment" {
  policy_arn = aws_iam_policy.aws-ingress-policy.arn
  role      = var.kubernetes_worker_iam_role_name
}

resource "aws_iam_policy" "aws-cluster-auto-scaler-policy" {
  name   = "ClusterAutoScaler-${var.kubernetes_cluster_name}"
  policy = data.aws_iam_policy_document.aws-cluster-auto-scaler-policy-document.json
}

resource "aws_iam_role_policy_attachment" "aws-cluster-auto-scaler-attachment" {
  policy_arn = aws_iam_policy.aws-cluster-auto-scaler-policy.arn
  role      = var.kubernetes_worker_iam_role_name
}
	`

	IAMK8SVariablestf = `
variable "kubernetes_cluster_name" {
  default = ""
  description = "The name of your kubernetes cluster"
}

variable "kubernetes_worker_iam_role_name" {
  default = ""
  description = "The name of the role attached to your k8s's workload runner nodes"
}
	`
)
