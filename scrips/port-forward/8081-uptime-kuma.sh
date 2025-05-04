#!/usr/bin/env bash

set -x
set -e
set -u

MY_NAMESPACE=uptime-kuma
MY_SERVICE=uptime-kuma
MY_PORT_IN=80
MY_PORT_OUT=8081

echo "go to: http://localhost:${MY_PORT_OUT}/"
kubectl -n ${MY_NAMESPACE} port-forward service/${MY_SERVICE} ${MY_PORT_OUT}:${MY_PORT_IN}
# firefox http://localhost:${MY_PORT}/ &&