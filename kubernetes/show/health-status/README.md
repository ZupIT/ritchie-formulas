# Description

This _health-status_ command checks if the container is running as expected.

How does ritchie do it? The command will fetch the port which is configured in
the livenessProbe of the container. If the container doesn't have it configured,
the command will fail. If it does have it, ritchie will make a port forward from
the container to your machine, and test if it can connect to it.

The user has to inform 2 different kinds of inputs:

- the namespace that the pod is running
- the name of the pod (or at least a part of it)

## Command

```bash
rit kubernetes show health-status
```

## Requirements

- [Kubernetes Installed](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Golang Installed](https://golang.org/doc/install)

## Demonstration

[![asciicast](https://asciinema.org/a/365442.svg)](https://asciinema.org/a/365442)
