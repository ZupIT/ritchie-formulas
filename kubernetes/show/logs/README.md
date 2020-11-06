<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# Kubernetes show logs

## Premisses

- Set kubeconfig credentials (`rit set credentials`)

## Command

- Prompt

```bash
rit kubernetes show logs
```

- Docker

```bash
rit kubernetes show logs --docker
```

- Stdin

```bash
echo '{"namespace": "kube-system", "pod_part_name": "metrics"}' | rit kubernetes show logs --stdin
```

- Stdin + docker

```bash
echo '{"namespace": "kube-system", "pod_part_name": "metrics"}' | rit kubernetes show logs --stdin --docker
```

## Description

This _show logs_ command will continuously show the logs of a running container in the provided cluster. The user has to inform 2 different kinds of inputs:

- the namespace that the pod is running
- the name of the pod (or at least a part of it)

## Demonstration

[![asciicast](https://asciinema.org/a/0o64yCJFCsFE6OHJpLBObtyuS.svg)](https://asciinema.org/a/0o64yCJFCsFE6OHJpLBObtyuS)
