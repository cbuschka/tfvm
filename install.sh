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

VERSION=$(curl --silent "https://api.github.com/repos/cbuschka/tfvm/releases" | jq -r '.[0].tag_name')

function get_os() {
  local unameOut="$(uname -s)"
  local os=""
  case "${unameOut}" in
    Linux*) os="linux";;
    Darwin*) os="darwin";;
    CYGWIN*|MINGW*) os="windows";;
    *) ;;
  esac
  echo -n "${os}"
}

function get_arch() {
  local unameOut="$(uname -m)"
  local arch=""
  case "${unameOut}" in
    x86_64*) arch="amd64";;
    386*) arch="386";;
    *) ;;
  esac
  echo -n "${arch}"
}

OS=$(get_os)
ARCH="$(get_arch)"
if [ -z "${OS}" ] || [ -z "${ARCH}" ]; then
  echo "Sorry, $(uname -s)/$(uname -m) not supported."
  exit 1
fi

echo "Downloading tfvm ${VERSION}..."
TMP_FILE=$(mktemp)
curl -L --progress-bar -o ${TMP_FILE} https://github.com/cbuschka/tfvm/releases/download/${VERSION}/tfvm-${OS}_${ARCH}

echo "Installing tfvm as ${TARGET_DIR}/tfvm..."
mv ${TMP_FILE} ${TARGET_DIR}/tfvm
chmod 755 ${TARGET_DIR}/tfvm
echo "Creating symlink ${TARGET_DIR}/terraform to ${TARGET_DIR}/tfvm..."
ln -s ${TARGET_DIR}/tfvm ${TARGET_DIR}/terraform

echo "tfvm successfully installed."
exit 0
