#!/bin/bash
runIstioctl() {
    if [ $1 = "yes" ]; then
        echo "Istioctl installation..."
        ISTIO_VERSION=$(curl -L -s https://api.github.com/repos/istio/istio/releases | \
                        grep tag_name | sed "s/ *\"tag_name\": *\"\\(.*\\)\",*/\\1/" | \
                        grep -v -E "(alpha|beta|rc)\.[0-9]$" | sort -t"." -k 1,1 -k 2,2 -k 3,3 -k 4,4 | tail -n 1)

        curl -L https://istio.io/downloadIstio | sh -
        mv istio-$ISTIO_VERSION/bin/istioctl /usr/local/bin
        rm -rf istio-$ISTIO_VERSION
    else
        echo "Skipping Istioctl installation..."
    fi
}

runIstioctlInstall() {
    istioctl install --set profile=$1
}
