#!/usr/bin/env bash

set -x
set -e
set -u

MY_NAMESPACE=openbao
MY_SERVICE=openbao
MY_PORT_IN=8200
MY_PORT_OUT=8083

echo "go to: http://localhost:${MY_PORT_OUT}/ui"
# kubectx irish-sea
kubectl -n ${MY_NAMESPACE} port-forward service/${MY_SERVICE} ${MY_PORT_OUT}:${MY_PORT_IN}
# firefox http://localhost:${MY_PORT}/ &&

