#!/usr/bin/env bash

set -o errexit

REPO_ROOT=$(git rev-parse --show-toplevel)

logs() {
  kubectl -n openfaas get all
  kubectl -n openfaas describe deployment/gateway
  kubectl -n openfaas logs deployment/gateway -c operator
}
trap "logs" EXIT SIGINT

echo ">>> Load OpenFaaS operator local image onto the cluster"
docker tag openfaas/openfaas-operator:latest test/openfaas-operator:latest
kind load docker-image test/openfaas-operator:latest

echo ">>> Create OpenFaaS namespaces"
kubectl apply -f https://raw.githubusercontent.com/openfaas/faas-netes/master/namespaces.yml

echo ">>> Create OpenFaaS CRD"
kubectl apply -f ${REPO_ROOT}/artifacts/operator-crd.yaml

echo ">>> Install OpenFaaS with Helm"
helm upgrade -i openfaas openfaas/openfaas \
--namespace openfaas \
--set openfaasImagePullPolicy=IfNotPresent \
--set functionNamespace=openfaas-fn \
--set generateBasicAuth=true \
--set async=false \
--set faasIdler.create=false \
--set operator.create=true \
--set operator.createCRD=false \
--set operator.image=test/openfaas-operator:latest

echo ">>> Patch operator deployment"
TEMP_DIR=$(mktemp -d)
cat > ${TEMP_DIR}/patch.yaml << EOL
spec:
 template:
  spec:
   containers:
   - name: operator
     command:
     - ./openfaas-operator
EOL
kubectl -n openfaas patch deployment gateway --patch "$(cat ${TEMP_DIR}/patch.yaml)"

echo ">>> Wait for operator deployment to be ready"
kubectl -n openfaas rollout status deployment/gateway --timeout=60s
