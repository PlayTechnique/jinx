#!/usr/bin/env bash -l

JINX="jinx"
THIS_SCRIPT_DIR=$(cd $(dirname $0) && pwd)

cd ${THIS_SCRIPT_DIR}/../..

# Some containers are running. Bail and ask the user to fix this.
if [[ $(docker ps --format "{{.Names}}") = "jinkies" ]]; then
  echo "jinkies containers are running. Please calm this situation and try again." >&2
  exit 1
fi

echo "Verifying jinx can start a container"

# Verify jinx can start a container
$(./jinx serve start > /dev/null)

OUTPUT=$(docker ps --format "{{.Names}}")
echo ${OUTPUT}

if [[ $? != 0 ]] || [[ "${OUTPUT}" != "jinkies" ]]; then
  echo "'jinx serve start' failed" >&2
  echo "Containers not started." >&2
  echo ${OUTPUT} >&2
  exit 2
fi

echo "Verifying jinx can stop a container"
$(./jinx serve stop > /dev/null)

if [[ $? != 0 ]] || [[ $(docker ps --format "{{.Names}}") = "jinkies" ]]; then
  echo "'jinx serve stop' failed" >&2
  echo "Containers not stopped." >&2
  echo ${OUTPUT} >&2
  exit 2
fi