#!/usr/bin/env bash

set -x
set -e
set -u

MY_NAMESPACE=pipelines-as-code
MY_SERVICE=pipelines-as-code-controller
MY_PORT_IN=8080
# MY_SERVICE=pipelines-as-code-webhook
# MY_PORT_IN=443
MY_PORT_OUT=8087

echo "go to: http://localhost:${MY_PORT_OUT}/"
# kubectx irish-sea
kubectl -n ${MY_NAMESPACE} port-forward service/${MY_SERVICE} ${MY_PORT_OUT}:${MY_PORT_IN}
# firefox http://localhost:${MY_PORT}/ &&

