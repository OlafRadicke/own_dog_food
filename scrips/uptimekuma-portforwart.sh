#!/usr/bin/env bash

set -x
set -e
set -u

MY_NAMESPACE=pulumi-apps
MY_SERVICE=foo
MY_PORT=80

echo "go to: http://localhost:${MY_PORT}/"
kubectx atlantic-ocean
kubectl -n ${MY_NAMESPACE} port-forward service/uptime-kuma ${MY_PORT}80:${MY_PORT}
# firefox http://localhost:${MY_PORT}/ &&