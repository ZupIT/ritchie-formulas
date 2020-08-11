#!/bin/bash

# shellcheck source=/dev/null
. "$(dirname "$0")"/scripts/minikube.sh --source-only
. "$(dirname "$0")"/scripts/istioctl.sh --source-only
#In sh for receive inputs of CLI use: $SAMPLE_TEXT, $SAMPLE_LIST and $SAMPLE_BOOL for this example

runMinikube "$MINIKUBE"
runMinikubeConfig "$MINIKUBE_CONFIG"
runIstioctl "$ISTIOCTL"
runIstioctlInstall "$PROFILE"
