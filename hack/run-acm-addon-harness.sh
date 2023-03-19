#!/bin/bash

oc --kubeconfig=$TEST_KUBECONFIG delete ns osde2e-ci-secrets || true

if [ -d /tmp/report ]; then
  echo "backing up old report folder ..."
  mv /tmp/report /tmp/report.$$
fi

sleep 10

ADDON_POLLING_TIMEOUT=300 \
DESTROY_CLUSTER=false \
MUST_GATHER=false \
osde2e test \
--configs aws,stage,addon-suite \
--label-filter "!Scale && !Upgrade" \
--skip-destroy-cluster true \
--focus-tests "ACM" \
--skip-health-check true

if [ -f /tmp/report/install/junit-acm-addon.xml ]; then
  cat /tmp/report/install/junit-acm-addon.xml
fi