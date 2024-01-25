#!/usr/bin/env bash

set -x
set -e
set -u

export PULUMI_CONFIG_PASSPHRASE_FILE=${HOME}/.ssh/pulumi-passwd
START_PWD=$(pwd)

cd pulumi/azure
go mod tidy
# go get -u


gcloud config set project pulumi-prod
pulumi login gs://pulumi-azure
# pulumi up --yes
pulumi up --stack prod

cd ${START_PWD}