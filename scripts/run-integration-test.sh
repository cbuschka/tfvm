#!/bin/bash

PROJECT_DIR=$(cd `dirname "$0"`/.. && pwd)

TEST_NAME=${1}
echo "Running ${TEST_NAME}..."

TEST_DIR=${PROJECT_DIR}/integration-tests/${TEST_NAME}
TEST_TEMP_DIR=$(mktemp -d)

if [ -d "${TEST_DIR}/inventory" ]; then
  cp -r ${TEST_DIR}/inventory/ ${TEST_TEMP_DIR}/
fi
if [ -d "${TEST_DIR}/workspace" ]; then
  cp -r ${TEST_DIR}/workspace/ ${TEST_TEMP_DIR}/
fi
mkdir -p ${TEST_TEMP_DIR}/tfvm ${TEST_TEMP_DIR}/workspace

if [ -f "${TEST_DIR}/env.sh" ]; then
  source ${TEST_DIR}/env.sh
fi

mkdir -p ${TEST_TEMP_DIR}/bin
ln -s ${PROJECT_DIR}/dist/tfvm-linux_amd64 ${TEST_TEMP_DIR}/bin/tfvm
ln -s ${PROJECT_DIR}/dist/tfvm-linux_amd64 ${TEST_TEMP_DIR}/bin/terraform
export PATH=${TEST_TEMP_DIR}/bin/:${PATH}

echo "tfvm: $(which tfvm)"
echo "terraform: $(which terraform)"

env | grep TFVM

export TFVM_DIR=${TEST_TEMP_DIR}/inventory/
unset TFVM_TERRAFORM_VERSION
unset TERRAFORM_VERSION
unset TFVM_TERRAFORM_ARCH
unset TERRAFORM_ARCH
unset TFVM_TERRAFORM_OS
unset TERRAFORM_OS
unset TFVM_TERRAFORM_RELEASES_BASE_URL

echo "Test temp dir..."
find ${TEST_TEMP_DIR}

cd ${TEST_TEMP_DIR}/workspace
${TEST_DIR}/run.sh 2>&1 1>${TEST_TEMP_DIR}/output
EXIT_CODE=$?

if [ -f "${TEST_DIR}/check.sh" ]; then
  ${TEST_DIR}/check.sh
  EXIT_CODE=$?
fi
echo "Exit code: ${EXIT_CODE}"

echo "Output..."
cat ${TEST_TEMP_DIR}/output

echo "Test temp dir..."
find ${TEST_TEMP_DIR}

echo "Cleaning up..."
rm -rf ${TEST_TEMP_DIR}

exit ${EXIT_CODE}
