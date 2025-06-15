#!/usr/bin/env bash

set -x
set -e
set -u

MY_NAMESPACE=smallstep-root-ca
MY_SERVICE=smallstep-root-ca-step-certificates
MY_PORT_IN=443
MY_PORT_OUT=8083

echo "go to: http://localhost:${MY_PORT_OUT}"
# kubectx atlantic-ocean
kubectl -n ${MY_NAMESPACE} port-forward service/${MY_SERVICE} ${MY_PORT_OUT}:${MY_PORT_IN}
# firefox http://localhost:${MY_PORT}/ &&