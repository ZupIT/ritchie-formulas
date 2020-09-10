#!/bin/sh
runFormula() {
  aws eks --region "$REGION" update-kubeconfig --name "$CLUSTER_NAME"
}
