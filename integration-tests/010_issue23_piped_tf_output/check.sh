#!/bin/bash

BANNER=$(grep Downloading version.txt)
if [ ! -z "${BANNER}" ]; then
  exit 1
fi
exit 0
