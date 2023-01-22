#!/bin/bash

cd `dirname $0`
instance_id=$(terraform -chdir=../terraform output -raw instance_id)
aws ec2 describe-instances --filters Name=instance-id,Values=$instance_id \
  | jq -r '.Reservations[].Instances[].NetworkInterfaces[].Association.PublicDnsName'
