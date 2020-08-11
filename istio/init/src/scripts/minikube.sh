#!/bin/bash
runMinikube() {
  if [[ $1 == "yes" ]]; then
    echo "Installing Minikube..."

    if [[ "$OSTYPE" == linux* ]]; then
      curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
      sudo install minikube-linux-amd64 /usr/local/bin/minikube
    elif [[ "$OSTYPE" == darwin* ]]; then
      brew install minikube
    else
      choco install minikube
    fi
  else
    echo "Skipping Minikube installation..."
  fi
}

runMinikubeConfig() {
  echo "Configuring Minikube..."
  minikube start --memory=8192mb --cpus=4 \
    --extra-config=apiserver.service-account-signing-key-file=/var/lib/minikube/certs/sa.key \
    --extra-config=apiserver.service-account-issuer=kubernetes/serviceaccount \
    --extra-config=apiserver.service-account-api-audiences=api
  echo "Your minikube IP is: $(minikube ip)"
  echo "Choose a free IP range to use with metallb"
  minikube addons configure metallb
  minikube addons enable metallb
}
