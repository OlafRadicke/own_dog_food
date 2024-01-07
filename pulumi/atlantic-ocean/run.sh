#!/usr/bin/env bash

set -x
set -e
set -u

export PULUMI_CONFIG_PASSPHRASE_FILE=${HOME}/.ssh/pulumi-passwd
START_PWD=$(pwd)

cd pulumi/atlantic-ocean

kubectx atlantic-ocean
kubectl get node

go mod tidy
# go get -u


gcloud config set project pulumi-prod
pulumi login gs://pulumi-atlantic-ocean
# pulumi up --yes
pulumi up

cd ${START_PWD}