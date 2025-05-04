#!/usr/bin/env bash

set -x
set -e
set -u

MY_NAMESPACE=argo
MY_SERVICE=argo-server
MY_PORT_IN=2746
MY_PORT_OUT=8084

echo "go to: https://localhost:${MY_PORT_OUT}"
# kubectx irish-sea
kubectl -n ${MY_NAMESPACE} port-forward service/${MY_SERVICE} ${MY_PORT_OUT}:${MY_PORT_IN}
