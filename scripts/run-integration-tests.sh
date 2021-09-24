#!/bin/bash

PROJECT_DIR=$(cd `dirname "$0"`/.. && pwd)

TESTS_DIR=${PROJECT_DIR}/integration-tests/

ALL_EXIT_CODE=0
for TEST_DIR in $(find ${TESTS_DIR} -mindepth 1 -maxdepth 1 -type d); do
  TEST_NAME=$(basename "${TEST_DIR}")
  ${PROJECT_DIR}/scripts/run-integration-test.sh ${TEST_NAME}
  EXIT_CODE=$?
  if [ "0" = "${EXIT_CODE}" ]; then
    echo "OK ${TEST_NAME}"
  else
    echo "FAILED ${TEST_NAME}"
    ALL_EXIT_CODE=1
  fi
done

exit ${ALL_EXIT_CODE}
