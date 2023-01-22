#!/bin/bash

cd `dirname $0`
instance_id=$(terraform -chdir=../terraform output -raw instance_id)
aws ec2 start-instances --instance-ids $instance_id
