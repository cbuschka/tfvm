#!/bin/bash
set -e

TIMESTAMP=$(date +%Y-%m-%dT%H:%M:%S)

if [ "x" = "x${TARGET_DIR}" ]; then
  if [ "${UID}" = "0" ]; then
    TARGET_DIR=/usr/local/bin
  elif [ -d "${HOME}/.local/bin" ] && [ ! -f "${HOME}/bin/tfvm" ]; then
    TARGET_DIR=${HOME}/.local/bin
  else
    TARGET_DIR=${HOME}/bin
  fi
fi
mkdir -p ${TARGET_DIR}

if [ ! -d "${TARGET_DIR}" ]; then
  echo "${TARGET_DIR} is no directory. Aborted."
  exit 1
fi

VERSION=$(basename $(curl -Ls -o /dev/null -w %\{url_effective\} https://github.com/cbuschka/tfvm/releases/latest))

CURRENT_TFVM_VERSION=$(tfvm version 2>/dev/null || true)
if [ "x${CURRENT_TFVM_VERSION}" = "x${VERSION}" ]; then
  echo "tfvm ${VERSION} is already installed."
  exit 0
fi

if [ -f "${TARGET_DIR}/terraform" ]; then
  mv ${TARGET_DIR}/terraform ${TARGET_DIR}/terraform.${TIMESTAMP}
  echo "${TARGET_DIR}/terraform already exists. Moved out of the way as ${TARGET_DIR}/terraform.${TIMESTAMP}."
fi

echo "Will install tfvm as ${TARGET_DIR}/tfvm and ${TARGET_DIR}/terraform."

function get_os() {
  local unameOut="$(uname -s)"
  case "${unameOut}" in
  Linux*)
    OS="linux"
    EXT=""
    ;;
  Darwin*)
    OS="darwin"
    EXT=""
    ;;
  CYGWIN* | MINGW*)
    OS="windows"
    EXT=".exe"
    ;;
  *)
    echo "${unameOut} is not supported. Aborted."
    exit 1
    ;;
  esac
}

function get_arch() {
  local unameOut="$(uname -m)"
  case "${unameOut}" in
  x86_64*) ARCH="amd64" ;;
  386*) ARCH="386" ;;
  arm64*) ARCH="arm64" ;;
  *)
    echo "${unameOut} is not supported. Aborted"
    exit 1
    ;;
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
echo "tfvm ${VERSION} successfully installed."

if [ "x" = "x$(echo ${PATH} | grep ${TARGET_DIR})" ]; then
  echo
  echo "Hint: ${TARGET_DIR} is not part of your PATH yet. Do not forget to add it."
fi

exit 0
