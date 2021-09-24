#!/bin/bash

TFVM_TERRAFORM_OS=darwin TFVM_TERRAFORM_ARCH=arm64 tfvm install 0.11.15
if [ ! -f "${TFVM_DIR}/v1/installed/0.11.15/darwin_amd64/terraform" ]; then
  echo "Expected ${TFVM_DIR}/v1/installed/0.11.15/darwin_amd64/terraform to be present."
  exit 1
fi

TFVM_TERRAFORM_OS=darwin TFVM_TERRAFORM_ARCH=arm64 tfvm install 1.0.7
if [ ! -f "${TFVM_DIR}/v1/installed/1.0.7/darwin_arm64/terraform" ]; then
  echo "Expected ${TFVM_DIR}/v1/installed/1.0.7/darwin_arm64/terraform to be present."
  exit 1
fi
