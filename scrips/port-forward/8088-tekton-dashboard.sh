#!/usr/bin/env bash

set -x
set -e
set -u

MY_NAMESPACE=tekton-deshboard
MY_SERVICE=tekton-dashboard
MY_PORT_IN=9097
# MY_SERVICE=pipelines-as-code-webhook
# MY_PORT_IN=443
MY_PORT_OUT=8088

echo "go to: http://localhost:${MY_PORT_OUT}/"
# kubectx irish-sea
kubectl -n ${MY_NAMESPACE} port-forward service/${MY_SERVICE} ${MY_PORT_OUT}:${MY_PORT_IN}
# firefox http://localhost:${MY_PORT}/ &&

