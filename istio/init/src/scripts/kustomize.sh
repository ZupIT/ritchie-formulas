#!/bin/bash
runKustomize() {
    if [[ $1 == "yes" ]]; then
        echo "Installing Kustomize..."
        if [[ "$OSTYPE" == linux* ]]; then
            curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash
        elif [[ "$OSTYPE" == darwin* ]]; then
            brew install kustomize
        else
            choco install kustomize
        fi
    else
        echo "Skipping Kustomize..."
    fi
}