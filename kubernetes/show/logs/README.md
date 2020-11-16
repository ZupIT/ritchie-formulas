<!-- markdownlint-disable-file MD013 -->
<!-- markdownlint-disable-file MD033 -->

# Description

This _show logs_ command will continuously show the logs of a running container in the provided cluster. The user has to inform 2 different kinds of inputs:

- the namespace that the pod is running
- the name of the pod (or at least a part of it)

## Command

```bash
rit kubernetes show logs
```

## Requirements

- kubeconfig credentials
- [GoLang Installed](https://golang.org/doc/install)

## Demonstration

[![asciicast](https://asciinema.org/a/0o64yCJFCsFE6OHJpLBObtyuS.svg)](https://asciinema.org/a/0o64yCJFCsFE6OHJpLBObtyuS)
