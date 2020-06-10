#!/bin/bash

CONSUL_POD=$(kubectl get pods -l="release=${CONSUL_RELEASE}" --output=jsonpath={.items[0].metadata.name})

while [ "$(kubectl exec $CONSUL_POD consul members --namespace=default | grep alive | wc -l)" != "3" ]; do
  sleep 5
done