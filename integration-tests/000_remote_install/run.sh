#!/bin/bash

export TARGET_DIR=$(cd $(dirname "$(which tfvm 2>/dev/null)") && pwd -P)
echo "TARGET_DIR is ${TARGET_DIR}"
rm ${TARGET_DIR}/tfvm ${TARGET_DIR}/terraform
export PATH=${TARGET_DIR}:/sbin:/bin:/usr/bin:/usr/sbin:/usr/local/bin
curl -sL https://raw.githubusercontent.com/cbuschka/tfvm/main/install.sh -o - | bash
tfvm info && TFVM_TERRAFORM_VERSION=latest terraform version
