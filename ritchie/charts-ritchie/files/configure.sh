#!/bin/bash

#Script to unleash the vault
mkdir build
cd build



while [ "$(kubectl -n $NAMESPACE get po $VAULT_CONTAINER -o jsonpath={.status.phase})" != "Running" ]; do
  sleep 5
done

kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault operator init > keys

for i in `cat keys | awk '{print $4}'|head -n 3`; do kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault operator unseal $i; done

ROOT_TOKEN=$(cat keys | grep "Root Token:" | awk '{print $4}')

cat > role.yaml <<EOF
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: role-tokenreview-binding-${APPLICATION}
  namespace: $NAMESPACE
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: vault-auth
  namespace: $NAMESPACE
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: vault-auth
  namespace: $NAMESPACE
EOF


cat > ${APPLICATION}_credential_policy.hcl <<EOF
path "auth/token/*" {
  capabilities = [ "create", "read", "update", "delete", "list", "sudo" ]
}

path "auth/approle/login" {
  capabilities = [ "create", "read" ]
}

path "auth/approle/role/${APPLICATION}_credential_role/role-id" {
  capabilities = [ "read" ]
}

path "auth/approle/role/${APPLICATION}_credential_role/secret-id" {
  capabilities = ["create", "read", "update"]
}

path "secret/*" {
  capabilities = ["create", "read"]
}

path "$APPLICATION/warmup/*" {
  capabilities = ["read","create","update"]
}

path "$APPLICATION/credential/*" {
  capabilities = ["create", "update", "delete", "read", "list"]
}
EOF

kubectl apply -f role.yaml

VAULT_AUTH_TOKEN_SECRET=$(kubectl -n $NAMESPACE get secret|grep vault-auth-token| awk '{print $1}')
VAULT_AUTH_SA_JWT_TOKEN=$(kubectl -n $NAMESPACE get secret $VAULT_AUTH_TOKEN_SECRET -o jsonpath="{.data.token}" | base64 --decode)
VAULT_AUTH_SA_CA_CRT=$(kubectl -n $NAMESPACE get secret $VAULT_AUTH_TOKEN_SECRET -o jsonpath="{.data['ca\.crt']}" | base64 --decode)

CURRENT_CONTEXT=$(kubectl config current-context)
CURRENT_CLUSTER=$(kubectl config view -o jsonpath="{.contexts[?(@.name == \"${CURRENT_CONTEXT}\"})].context.cluster}")
K8S_HOST=$(kubectl config view -o jsonpath="{.clusters[?(@.name == \"${CURRENT_CLUSTER}\"})].cluster.server}")


kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault login $ROOT_TOKEN

kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault secrets enable -path secret/ generic

kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault secrets enable -path $APPLICATION/warmup/ generic

kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault secrets enable -path $APPLICATION/credential/ generic

kubectl cp ${APPLICATION}_credential_policy.hcl $NAMESPACE/$VAULT_CONTAINER:/tmp/${APPLICATION}_credential_policy.hcl

kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault policy write ${APPLICATION}_credential_policy /tmp/${APPLICATION}_credential_policy.hcl

kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault auth enable kubernetes

kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault write auth/kubernetes/config token_reviewer_jwt="$VAULT_AUTH_SA_JWT_TOKEN" kubernetes_host="$K8S_HOST:443" kubernetes_ca_cert="$VAULT_AUTH_SA_CA_CRT"

kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault write auth/kubernetes/role/${APPLICATION}_credential_role bound_service_account_names=vault-auth bound_service_account_namespaces=$NAMESPACE policies=default,${APPLICATION}_credential_policy ttl=1h

kubectl -n $NAMESPACE exec $VAULT_CONTAINER -c vault -- vault write auth/token/roles/$APPLICATION allowed_policies="default,${APPLICATION}_credential_policy" period="24h"
