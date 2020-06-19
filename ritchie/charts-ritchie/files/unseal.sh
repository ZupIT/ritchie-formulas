#!/bin/sh


for vault_pod in `kubectl get pods -l app.kubernetes.io/instance=vault-test -o jsonpath='{.items[*].metadata.name}'`; do 
  if [ $(kubectl exec -it $vault_pod -- vault status -format json 2> /dev/null|jq .sealed|grep true) ]; then
	  echo "$vault_pod"
 	  for i in `cat build/keys | awk '{print $4}'|head -n 3`; do 
		  echo "kubectl -n $NAMESPACE exec $vault_pod -c vault -- vault operator unseal $i"; 
		  kubectl -n $NAMESPACE exec $vault_pod -c vault -- vault operator unseal $i;
	  done
  fi
done
