#!/bin/bash

set -e # -x

echo "-----> `date`: Upload stemcell"
bosh -n upload-stemcell "https://bosh.io/d/stemcells/bosh-warden-boshlite-ubuntu-trusty-go_agent?v=3421.4" \
  --sha1 e7c440fc20bb4bea302d4bfdc2369367d1a3666e \
  --name bosh-warden-boshlite-ubuntu-trusty-go_agent \
  --version 3421.4

echo "-----> `date`: Delete previous deployment"
bosh -n -d cockroachdb delete-deployment --force

echo "-----> `date`: Deploy"
( set -e; cd ./..; bosh -n -d cockroachdb deploy ./manifests/example.yml \
  -o ./manifests/dev.yml \
  -o ./manifests/protect.yml \
  -o ./manifests/protect-allow-ssh.yml --vars-store creds.yml )

echo "-----> `data`: Use included CLI binary"
bosh -n -d cockroachdb ssh cockroachdb_root \
  -c 'sudo /var/vcap/jobs/cockroachdb_user/bin/cli sql -e "show databases;"'

echo "-----> `data`: View node debug data"
bosh -n -d cockroachdb ssh cockroachdb \
  -c '/var/vcap/jobs/cockroachdb/bin/debug' -r --column instance --column exit_code

echo "-----> `date`: Recreate all VMs"
bosh -n -d cockroachdb recreate

echo "-----> `date`: Exercise deployment"
bosh -n -d cockroachdb run-errand smoke-tests

echo "-----> `date`: Restart deployment"
bosh -n -d cockroachdb restart

echo "-----> `date`: Report any problems"
bosh -n -d cockroachdb cck --report

echo "-----> `date`: Delete random VM"
bosh -n -d cockroachdb delete-vm `bosh -d cockroachdb vms --column vm_cid|sort|head -1`

echo "-----> `date`: Fix deleted VM"
bosh -n -d cockroachdb cck --auto

echo "-----> `date`: Exercise deployment again"
bosh -n -d cockroachdb run-errand smoke-tests

echo "-----> `date`: Delete deployment"
bosh -n -d cockroachdb delete-deployment

echo "-----> `date`: Done"
