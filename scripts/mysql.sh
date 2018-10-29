#!/bin/bash

# Connects to Aurora via the Bastion
# https://us-east-2.console.aws.amazon.com/rds/home?region=us-east-2#cluster:ids=moss

BASTION_HOSTNAME=ubuntu@ec2-13-58-201-2.us-east-2.compute.amazonaws.com
BASTION_KEY=~/.ssh/mecolinkingcobastion.pem
AURORA_HOSTNAME=moss.cluster-ccyz1djlth1k.us-east-2.rds.amazonaws.com
AURORA_PORT=3306
AURORA_USER=root
AURORA_PASSWORD=${1:?"Pass a password as the first parameter"}

ssh -tt ${BASTION_HOSTNAME} -i ${BASTION_KEY} "mysql --host='${AURORA_HOSTNAME}' --port=${AURORA_PORT} --user=${AURORA_USER} --password='${AURORA_PASSWORD}'"
