# Description

Installs **latest** `istioctl`, `kustomize` and `minikube` dependencies in
order to properly run the service mesh in your local environment

## Additional info

### Minikube

In order to run Istio, your k8s cluster needs to support some extra features
that minikube delivers in a easy and configurable way.

`minikube` starts a cluster with `8GB` RAM and `4` CPUs and **3rd party service
account token** configured (see [Istio Security Best Practices](
    https://istio.io/latest/docs/ops/best-practices/security/))

#### LoadBalancer IP

Addon `metallb` is enabled by default and configured during the playbook.
It will output your `minikube ip` as reference and will ask for an IP range to
allocate. You can use the last 8 bits at will.

##### Example

If your `minikube ip`  is `192.168.64.1` you can set the IP range to
`192.168.64.100` to `192.168.64.119` so you'll be able to create
20 LoadBalancer services in your cluster

### Istioctl

`istioctl` uses the Operator pattern to install `istio`

## Command

```bash
rit istio init
```
