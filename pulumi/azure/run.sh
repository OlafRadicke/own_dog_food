#!/usr/bin/env bash

set -x
set -e
set -u

export PULUMI_CONFIG_PASSPHRASE_FILE=${HOME}/.ssh/pulumi-passwd
export START_PWD=$(pwd)
export WORK_DIR="pulumi/azure"

if [ -z ${INSTALL_AND_UPDATE_GO_MODS+x} ]
then
	INSTALL_AND_UPDATE_GO_MODS="FALSE"
	echo "var is set to '$INSTALL_AND_UPDATE_GO_MODS'"
else
	echo "var is set to '$INSTALL_AND_UPDATE_GO_MODS'"
	# INSTALL_AND_UPDATE_GO_MODS="TRUE"
fi

cd $WORK_DIR
go mod tidy
if [ ${INSTALL_AND_UPDATE_GO_MODS} == "TRUE" ]
then
	go get -u -t -x
else
	echo "don't update go modules."
fi

gcloud config set project pulumi-prod
pulumi login gs://pulumi-azure
# pulumi up --yes
pulumi up --debug --stack prod

cd ${START_PWD}