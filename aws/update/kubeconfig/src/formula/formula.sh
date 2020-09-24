#!/bin/bash
runFormula() {
  if [[ $DOCKER_EXECUTION ]] ; then
    mkdir -p /root/.aws/
    echo "[default]" > /root/.aws/credentials
    echo "aws_access_key_id=$ACCESS_KEY" >> /root/.aws/credentials
    echo "aws_secret_access_key=$SECRET_ACCESS_KEY" >> /root/.aws/credentials
  fi
  echo "Running: aws eks --region $REGION update-kubeconfig --name $CLUSTER_NAME ..."
  aws eks --region "$REGION" update-kubeconfig --name "$CLUSTER_NAME"
}
