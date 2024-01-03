#!/usr/bin/env bash

set -x
set -e
set -u

export PULUMI_CONFIG_PASSPHRASE_FILE=${HOME}/.ssh/pulumi-passwd
START_PWD=$(pwd)

cd pulumi/baltic-sea

kubectx baltic-sea
kubectl get node

# go get -u
# go mod tidy

# pulumi up --yes
pulumi up

cd ${START_PWD}