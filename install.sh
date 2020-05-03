#!/bin/bash
set -e

TARGET_DIR=/usr/local/bin/
if [ "${UID}" != "0" ]; then
  TARGET_DIR=${HOME}/bin
  mkdir -p ${TARGET_DIR}
fi

if [ ! -d "${TARGET_DIR}" ]; then
  echo "${TARGET_DIR} is no directory. Aborted."
  exit 1
fi

if [ -f "${TARGET_DIR}/tfvm" ]; then
  echo "${TARGET_DIR}/tfvm already exists. Aborted."
  exit 1
fi

if [ -f "${TARGET_DIR}/terraform" ]; then
  echo "${TARGET_DIR}/terraform already exists. Aborted."
  exit 1
fi

echo "Will install tfvm as ${TARGET_DIR}/tfvm and ${TARGET_DIR}/terraform."

VERSION=$(curl --silent "https://api.github.com/repos/cbuschka/tfvm/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")')

function get_os() {
  local unameOut="$(uname -s)"
  case "${unameOut}" in
    Linux*) OS="linux"; EXT="";;
    Darwin*) OS="darwin"; EXT="";;
    CYGWIN*|MINGW*) OS="windows"; EXT=".exe";;
    *) echo "${unameOut} is not supported. Aborted"; exit 1;;
  esac
}

function get_arch() {
  local unameOut="$(uname -m)"
  case "${unameOut}" in
    x86_64*) ARCH="amd64";;
    386*) ARCH="386";;
    *) echo "${unameOut} is not supported. Aborted"; exit 1;;
  esac
}

get_os
get_arch

echo "Downloading tfvm ${VERSION}..."
TMP_FILE=$(mktemp)
curl -L --progress-bar -o ${TMP_FILE} https://github.com/cbuschka/tfvm/releases/download/${VERSION}/tfvm-${OS}_${ARCH}${EXT}

TARGET_TFVM=${TARGET_DIR}/tfvm${EXT}
TARGET_TERRAFORM=${TARGET_DIR}/terraform${EXT}

echo "Installing tfvm as ${TARGET_TFVM}..."
mv ${TMP_FILE} ${TARGET_TFVM}
chmod 755 ${TARGET_TFVM}
echo "Creating symlink ${TARGET_TERRAFORM} to ${TARGET_TFVM}..."
ln -s ${TARGET_TFVM} ${TARGET_TERRAFORM}

echo "tfvm successfully installed."
exit 0
